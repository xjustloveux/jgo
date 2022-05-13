// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jtime"
	"time"
)

type Schedule interface {
	// Name returns schedule name
	Name() (string, error)
	// CronExpression returns schedule cron expression string
	CronExpression() (string, error)
	// Job returns schedule job name
	Job() (string, error)
	// Data returns schedule default map data
	Data() (map[string]interface{}, error)
	// Desc returns schedule description
	Desc() (string, error)
	// Status returns schedule status
	Status() (Status, error)
	// NextTime returns schedule next execution time
	NextTime() (time.Time, error)
	// PrevTime returns schedule prev execution time
	PrevTime() (time.Time, error)
	// Stop interrupt schedule
	Stop() error
	// Resume restart schedule
	Resume() error
	// Trigger trigger schedule
	Trigger(map[string]interface{}) error
}

type schedule struct {
	name           string
	cronExpression string
	cron           *CronExpression
	job            string
	data           map[string]interface{}
	desc           string
	status         Status
	next           time.Time
	prev           time.Time
}

func (s *schedule) Name() (string, error) {
	if status == Stop {
		return s.name, nil
	} else {
		inCh <- channel{
			action: getSchName,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	}
}

func (s *schedule) CronExpression() (string, error) {
	if status == Stop {
		return s.cronExpression, nil
	} else {
		inCh <- channel{
			action: getSchCronExpression,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	}
}

func (s *schedule) Job() (string, error) {
	if status == Stop {
		return s.job, nil
	} else {
		inCh <- channel{
			action: getSchJob,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	}
}

func (s *schedule) Data() (map[string]interface{}, error) {
	if status == Stop {
		return s.data, nil
	} else {
		inCh <- channel{
			action: getSchData,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		if c.err != nil {
			return nil, c.err
		}
		return jcast.StringMapInterface(c.data)
	}
}

func (s *schedule) Desc() (string, error) {
	if status == Stop {
		return s.desc, nil
	} else {
		inCh <- channel{
			action: getSchDesc,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	}
}

func (s *schedule) Status() (Status, error) {
	if status == Stop {
		return s.status, nil
	} else {
		inCh <- channel{
			action: getSchSts,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		if c.err != nil {
			return Unknown, c.err
		}
		if sts, ok := c.data.(Status); ok {
			return sts, nil
		} else {
			return Unknown, errorf(errorDataStatus, c.data)
		}
	}
}

func (s *schedule) NextTime() (time.Time, error) {
	if status == Stop {
		return s.next, nil
	} else {
		inCh <- channel{
			action: getSchNext,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		if c.err != nil {
			return time.Time{}, c.err
		}
		return jcast.Time(c.data)
	}
}

func (s *schedule) PrevTime() (time.Time, error) {
	if status == Stop {
		return s.prev, nil
	} else {
		inCh <- channel{
			action: getSchPrev,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		if c.err != nil {
			return time.Time{}, c.err
		}
		return jcast.Time(c.data)
	}
}

func (s *schedule) Stop() error {
	if status == Stop {
		var nt time.Time
		s.status = Stop
		s.next = nt
		s.prev = nt
		return nil
	} else {
		inCh <- channel{
			action: stopSch,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return c.err
	}
}

func (s *schedule) Resume() error {
	if status == Stop {
		if totalSch[s.name] == s {
			add := true
			for _, rs := range runSch {
				if rs == s {
					add = false
					break
				}
			}
			if add {
				s.status = Run
				runSch = append(runSch, s)
			}
		} else {
			return errors(errorUnknownSch)
		}
		return nil
	} else {
		inCh <- channel{
			action: resumeSch,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return c.err
	}
}

func (s *schedule) Trigger(data map[string]interface{}) error {
	if status == Stop {
		if j := jobs[s.job]; j != nil {
			e := &entry{
				j:    j,
				data: data,
			}
			wgTrigger.Add(1)
			go e.run(wgTrigger)
			return nil
		} else {
			return errors(errorJobNil)
		}
	} else {
		inCh <- channel{
			action: trigger,
			data: &entry{
				sch:  s,
				data: data,
			},
			err: nil,
		}
		c := <-outCh
		return c.err
	}
}

func (s *schedule) toNext(n time.Time) {
	if !s.next.IsZero() {
		s.prev = s.next
	}
	s.next = s.getNext(n.Add(jtime.Second))
	if s.status == Stop {
		var nt time.Time
		s.next = nt
		s.prev = nt
	}
}

func (s *schedule) getNext(t time.Time) time.Time {
	var err error
	year := s.getNum(t.Year(), s.cron.year)
	if v := year - t.Year(); v > 0 {
		if t, err = s.getTime(t.Year()+v, s.cron.month[0], 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
	} else if v < 0 {
		s.status = Stop
		return t
	}
	nm := int(t.Month())
	month := s.getNum(nm, s.cron.month)
	if v := month - nm; v > 0 {
		if t, err = s.getTime(t.Year(), nm+v, 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
	} else if v < 0 {
		if t, err = s.getTime(t.Year()+1, s.cron.month[0], 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	if len(s.cron.day) > 0 {
		day := s.getNum(t.Day(), s.cron.day)
		if v := day - t.Day(); v > 0 {
			m := t.Month()
			t = t.Add(time.Duration(v) * jtime.Day)
			if m != t.Month() {
				if t, err = s.getTime(t.Year(), int(t.Month()), 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
					fmtPrintln(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			} else {
				if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
					fmtPrintln(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			}
		} else if v < 0 {
			ty := t.Year()
			tm := int(t.Month())
			tm++
			if tm > 12 {
				ty++
				tm = 1
			}
			if t, err = s.getTime(ty, tm, 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
				fmtPrintln(err)
				s.status = Stop
				return t
			}
			return s.getNext(t)
		}
	} else {
		m := t.Month()
		w := int(t.Weekday())
		week := s.getNum(w, s.cron.weekday)
		if v := week - w; v > 0 {
			t = t.Add(time.Duration(v) * jtime.Day)
			if m != t.Month() {
				if t, err = s.getTime(t.Year(), int(t.Month()), 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
					fmtPrintln(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			} else {
				if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
					fmtPrintln(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			}
		} else if v < 0 {
			v = 7 - w + s.cron.weekday[0]
			t = t.Add(time.Duration(v) * jtime.Day)
			if m != t.Month() {
				if t, err = s.getTime(t.Year(), int(t.Month()), 1, s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
					fmtPrintln(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			} else {
				if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
					fmtPrintln(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			}
		}
	}
	hr := s.getNum(t.Hour(), s.cron.hour)
	if v := hr - t.Hour(); v > 0 {
		t = t.Add(time.Duration(v) * jtime.Hour)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), s.cron.minute[0], s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	} else if v < 0 {
		t = t.Add(jtime.Day)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), s.cron.hour[0], s.cron.minute[0], s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	min := s.getNum(t.Minute(), s.cron.minute)
	if v := min - t.Minute(); v > 0 {
		t = t.Add(time.Duration(v) * jtime.Minute)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	} else if v < 0 {
		t = t.Add(jtime.Hour)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), s.cron.minute[0], s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	sec := s.getNum(t.Second(), s.cron.second)
	if v := sec - t.Second(); v > 0 {
		t = t.Add(time.Duration(v) * jtime.Second)
	} else if v < 0 {
		t = t.Add(jtime.Minute)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), s.cron.second[0]); err != nil {
			fmtPrintln(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	return t
}

func (s *schedule) getNum(n int, arr []int) int {
	var num int
	if len(arr) > 0 {
		for _, v := range arr {
			num = v
			if v >= n {
				break
			}
		}
	} else {
		num = n
	}
	return num
}

func (s *schedule) getTime(year, month, day, hr, min, sec int) (time.Time, error) {
	return jcast.Time(fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hr, min, sec))
}
