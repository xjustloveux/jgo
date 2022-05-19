// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jfile"
	"sort"
	"sync"
	"time"
)

const (
	errorNotValidStatus         = jError("not a valid status %q")
	errorNotValidCronExpression = jError("not a valid CronExpression %q")
	errorJobNil                 = jError("job is nil")
	errorFuncNil                = jError("function is nil")
	errorDataJob                = jError("data type is %q, not *job")
	errorDataSch                = jError("data type is %q, not *schedule")
	errorDataEntry              = jError("data type is %q, not *entry")
	errorDataString             = jError("data type is %q, not string")
	errorDataStatus             = jError("data type is %q, not Status")
	errorUnknownSch             = jError("unknown schedule")
	errorSchRun                 = jError("schedule has already been run")
)

const (
	pkgName  = "jcron"
	fileName = "cron.json"
	minYear  = 2020
	maxYear  = 2080
)

var (
	conf      = jconf.New()
	cd        = &configData{}
	status    = Stop
	wg        = new(sync.WaitGroup)
	wgTrigger = new(sync.WaitGroup)
	logFunc   func(...interface{})
	jobs      map[string]Job
	totalSch  map[string]*schedule
	runSch    []*schedule
	inCh      chan channel
	outCh     chan channel
	noneCh    chan channel
)

func init() {
	SetFileName(fileName)
}

// SetFormat set config format
func SetFormat(f jfile.Format) {
	conf.SetFormat(f)
}

// SetFileName set config file name
func SetFileName(name string) {
	conf.SetFileName(name)
}

// SetEnvFileName set config env file name
func SetEnvFileName(name string) {
	conf.SetEnvFileName(name)
}

// EnvKey returns env key
func EnvKey() string {
	return conf.EnvKey()
}

// SetEnvKey set env key
func SetEnvKey(key string) {
	conf.SetEnvKey(key)
}

// EnvVal returns env value
func EnvVal() string {
	return conf.EnvVal()
}

// SetEnvVal set env value
func SetEnvVal(val string) {
	conf.SetEnvVal(val)
}

// EnableEnv enable env
func EnableEnv() {
	conf.EnableEnv()
}

// DisableEnv disable env
func DisableEnv() {
	conf.DisableEnv()
}

// SetLogFunc set fmt.Println log function
func SetLogFunc(f func(...interface{})) {
	logFunc = f
}

// Init initialize
func Init() error {
	if err := conf.Load(); err != nil {
		return err
	}
	cd = &configData{Debug: false}
	if err := conf.Convert(cd); err != nil {
		return err
	}
	if err := createSchedule(); err != nil {
		return err
	}
	return nil
}

// AddJob add job
func AddJob(name string, j Job) error {
	if status == Run {
		inCh <- channel{
			action: addJob,
			data: &job{
				name: name,
				j:    j,
			},
			err: nil,
		}
		c := <-outCh
		return c.err
	} else {
		jobs[name] = j
	}
	return nil
}

// AddJobFunc add job with function
func AddJobFunc(name string, f func(map[string]interface{})) error {
	if f == nil {
		return AddJob(name, nil)
	} else {
		return AddJob(name, jobFunc(f))
	}
}

// AddSchedule add schedule
func AddSchedule(sch *SchInfo) (err error) {
	var cron *CronExpression
	if cron, err = ParseCronExpression(sch.Cron); err != nil {
		return err
	}
	s := &schedule{
		name:           sch.Name,
		cronExpression: sch.Cron,
		cron:           cron,
		job:            sch.JobName,
		data:           sch.JobData,
		desc:           sch.Desc,
		status:         Run,
	}
	if status == Run {
		inCh <- channel{
			action: addSch,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return c.err
	} else {
		totalSch[s.name] = s
		runSch = append(runSch, s)
	}
	return nil
}

// GetSchedule returns schedule
func GetSchedule(name string) (Schedule, error) {
	if status == Run {
		inCh <- channel{
			action: getSch,
			data:   name,
			err:    nil,
		}
		c := <-outCh
		if c.err != nil {
			return nil, c.err
		}
		if v, ok := c.data.(*schedule); ok {
			if v == nil {
				return nil, errors(errorUnknownSch)
			}
			return v, c.err
		} else {
			return nil, errorf(errorDataSch, c.data)
		}
	} else {
		if sch := totalSch[name]; sch == nil {
			return nil, errors(errorUnknownSch)
		} else {
			return sch, nil
		}
	}
}

// GetStatus returns cron status
func GetStatus() Status {
	return status
}

// Start cron start
func Start() {
	if status != Stop {
		return
	}
	inCh = make(chan channel)
	outCh = make(chan channel)
	noneCh = make(chan channel, 10)
	wg.Add(1)
	go run()
	status = Run
}

// Interrupt cron stop
func Interrupt() {
	if status != Run {
		return
	}
	defer func() {
		close(inCh)
		close(outCh)
		close(noneCh)
	}()
	status = Stop
	inCh <- channel{
		action: stop,
		data:   nil,
		err:    nil,
	}
	<-outCh
}

// Wait sync.WaitGroup.Wait()
func Wait() {
	if status != Run {
		return
	}
	defer func() {
		close(inCh)
		close(outCh)
		close(noneCh)
	}()
	status = SyncWait
	inCh <- channel{
		action: stop,
		data:   nil,
		err:    nil,
	}
	<-outCh
	wg.Wait()
	status = Stop
}

// WaitTrigger sync.WaitGroup.Wait()
func WaitTrigger() {
	wgTrigger.Wait()
}

// Trigger trigger job
func Trigger(j Job, data map[string]interface{}) error {
	if j != nil {
		if status == Run {
			inCh <- channel{
				action: trigger,
				data: &entry{
					j:    j,
					data: data,
				},
				err: nil,
			}
			c := <-outCh
			return c.err
		} else {
			e := &entry{
				j:    j,
				data: data,
			}
			wgTrigger.Add(1)
			go e.run(wgTrigger)
			return nil
		}
	} else {
		return errors(errorJobNil)
	}
}

// TriggerFunc trigger job with function
func TriggerFunc(f func(map[string]interface{}), data map[string]interface{}) error {
	if f != nil {
		if status == Run {
			inCh <- channel{
				action: trigger,
				data: &entry{
					j:    jobFunc(f),
					data: data,
				},
				err: nil,
			}
			c := <-outCh
			return c.err
		} else {
			e := &entry{
				j:    jobFunc(f),
				data: data,
			}
			wgTrigger.Add(1)
			go e.run(wgTrigger)
			return nil
		}
	} else {
		return errors(errorFuncNil)
	}
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}

func fmtPrintln(args ...interface{}) {
	if cd.Debug {
		fmt.Println(args...)
	}
	if logFunc != nil {
		logFunc(args...)
	}
}

func createSchedule() error {
	Wait()
	jobs = make(map[string]Job)
	totalSch = make(map[string]*schedule)
	runSch = make([]*schedule, 0)
	for _, info := range cd.Schedule {
		if err := AddSchedule(info); err != nil {
			return err
		}
	}
	return nil
}

func run() {
	defer func() {
		wg.Done()
	}()
	var nextTime time.Time
	now := time.Now().Local()
	for _, s := range runSch {
		if jobs[s.job] == nil {
			var nt time.Time
			s.status = Stop
			s.next = nt
			s.prev = nt
		} else if s.status == Run {
			s.toNext(now)
		}
	}
	bk := false
	for {
		runSch = removeSch()
		sort.Sort(sortSch(runSch))
		if len(runSch) > 0 {
			nextTime = runSch[0].next
		} else {
			nextTime = now.AddDate(15, 0, 0)
		}
		select {
		case <-time.After(nextTime.Sub(now)):
			now = time.Now().Local()
			for _, s := range runSch {
				if now.After(s.next) {
					if j := jobs[s.job]; j != nil {
						s.toNext(s.next)
						e := &entry{
							j:    j,
							data: s.data,
						}
						wg.Add(1)
						go e.run(wg)
					} else {
						var nt time.Time
						s.status = Stop
						s.next = nt
						s.prev = nt
					}
				} else {
					break
				}
			}
		case c := <-inCh:
			c = runCh(c, now)
			outCh <- c
			if c.action == stop {
				bk = true
			}
		case c := <-noneCh:
			c = runCh(c, now)
			if c.action == stop {
				bk = true
			}
		}
		if bk {
			break
		}
	}
}

func removeSch() []*schedule {
	ns := make([]*schedule, 0)
	for _, s := range runSch {
		if s.status == Run {
			ns = append(ns, s)
		}
	}
	return ns
}

func runCh(c channel, now time.Time) channel {
	switch c.action {
	case addJob:
		if j, ok := c.data.(*job); ok {
			jobs[j.name] = j.j
		} else {
			c.err = errorf(errorDataJob, c.data)
		}
	case addSch:
		if s, ok := c.data.(*schedule); ok {
			now = time.Now().Local()
			if ts := totalSch[s.name]; ts != nil && ts != s {
				noneCh <- channel{
					action: stopSch,
					data:   ts,
					err:    nil,
				}
			}
			s.toNext(now)
			totalSch[s.name] = s
			runSch = append(runSch, s)
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case stopSch:
		if s, ok := c.data.(*schedule); ok {
			if s.status == Run {
				var nt time.Time
				s.status = Stop
				s.next = nt
				s.prev = nt
			}
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case resumeSch:
		if s, ok := c.data.(*schedule); ok {
			if totalSch[s.name] == s {
				if s.status == Stop {
					add := true
					for _, rs := range runSch {
						if rs == s {
							add = false
							break
						}
					}
					s.status = Run
					if add {
						noneCh <- channel{
							action: addSch,
							data:   s,
							err:    nil,
						}
					} else {
						c.err = errors(errorSchRun)
					}
				} else {
					c.err = errors(errorSchRun)
				}
			} else {
				c.err = errors(errorUnknownSch)
			}
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case trigger:
		if e, ok := c.data.(*entry); ok {
			if e.sch != nil {
				e.j = jobs[e.sch.job]
			}
			if e.j == nil {
				c.err = errors(errorJobNil)
			} else {
				wgTrigger.Add(1)
				go e.run(wgTrigger)
			}
		} else {
			c.err = errorf(errorDataEntry, c.data)
		}
	case getSch:
		if v, ok := c.data.(string); ok {
			c.data = totalSch[v]
		} else {
			c.err = errorf(errorDataString, c.data)
		}
	case getSchName:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.name
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchCronExpression:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.cronExpression
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchJob:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.job
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchData:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.data
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchDesc:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.desc
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchSts:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.status
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchNext:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.next
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	case getSchPrev:
		if s, ok := c.data.(*schedule); ok {
			c.data = s.prev
		} else {
			c.err = errorf(errorDataSch, c.data)
		}
	}
	return c
}
