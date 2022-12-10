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
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
		inCh <- channel{
			action: getSchName,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	} else {
		return s.name, nil
	}
}

func (s *schedule) CronExpression() (string, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
		inCh <- channel{
			action: getSchCronExpression,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	} else {
		return s.cronExpression, nil
	}
}

func (s *schedule) Job() (string, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
		inCh <- channel{
			action: getSchJob,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	} else {
		return s.job, nil
	}
}

func (s *schedule) Data() (map[string]interface{}, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
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
	} else {
		return s.data, nil
	}
}

func (s *schedule) Desc() (string, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
		inCh <- channel{
			action: getSchDesc,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return jcast.String(c.data), c.err
	} else {
		return s.desc, nil
	}
}

func (s *schedule) Status() (Status, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
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
			return Unknown, errorFmt(errorDataStatus, c.data)
		}
	} else {
		return s.status, nil
	}
}

func (s *schedule) NextTime() (time.Time, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
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
	} else {
		return s.next, nil
	}
}

func (s *schedule) PrevTime() (time.Time, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
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
	} else {
		return s.prev, nil
	}
}

func (s *schedule) Stop() error {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
		inCh <- channel{
			action: stopSch,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return c.err
	} else {
		var nt time.Time
		s.status = Stop
		s.next = nt
		s.prev = nt
		return nil
	}
}

func (s *schedule) Resume() error {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
		inCh <- channel{
			action: resumeSch,
			data:   s,
			err:    nil,
		}
		c := <-outCh
		return c.err
	} else {
		if totalSch[s.name] == s {
			add := true
			for _, rs := range runSch {
				if rs == s {
					add = false
					break
				}
			}
			s.status = Run
			if add {
				runSch = append(runSch, s)
			}
		} else {
			return errorStr(errorUnknownSch)
		}
		return nil
	}
}

func (s *schedule) Trigger(data map[string]interface{}) error {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if status == Run {
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
	} else {
		if j := jobs[s.job]; j != nil {
			e := &entry{
				j:    j,
				data: data,
			}
			wgTrigger.Add(1)
			go e.run(wgTrigger)
			return nil
		} else {
			return errorStr(errorJobNil)
		}
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
	year := s.getNum(t.Year()-minYear, s.cron.year) + minYear
	if v := year - t.Year(); v > 0 {
		if t, err = s.getTime(t.Year()+v, 1, 1, 0, 0, 0); err != nil {
			subject.Next(err)
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
		if t, err = s.getTime(t.Year(), nm+v, 1, 0, 0, 0); err != nil {
			subject.Next(err)
			s.status = Stop
			return t
		}
	} else if v < 0 {
		if t, err = s.getTime(t.Year()+1, 1, 1, 0, 0, 0); err != nil {
			subject.Next(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	if s.cron.day > 0 {
		day := s.getNum(t.Day(), s.cron.day)
		if v := day - t.Day(); v > 0 {
			m := t.Month()
			t = t.Add(time.Duration(v) * jtime.Day)
			if m != t.Month() {
				if t, err = s.getTime(t.Year(), int(t.Month()), 1, 0, 0, 0); err != nil {
					subject.Next(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			} else if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), 0, 0, 0); err != nil {
				subject.Next(err)
				s.status = Stop
				return t
			}
		} else if v < 0 {
			ty := t.Year()
			tm := int(t.Month())
			tm++
			if tm > 12 {
				ty++
				tm = 1
			}
			if t, err = s.getTime(ty, tm, 1, 0, 0, 0); err != nil {
				subject.Next(err)
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
				if t, err = s.getTime(t.Year(), int(t.Month()), 1, 0, 0, 0); err != nil {
					subject.Next(err)
					s.status = Stop
					return t
				}
				return s.getNext(t)
			} else if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), 0, 0, 0); err != nil {
				subject.Next(err)
				s.status = Stop
				return t
			}
		} else if v < 0 {
			v = 7 - w
			t = t.Add(time.Duration(v) * jtime.Day)
			if m != t.Month() {
				if t, err = s.getTime(t.Year(), int(t.Month()), 1, 0, 0, 0); err != nil {
					subject.Next(err)
					s.status = Stop
					return t
				}
			} else if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), 0, 0, 0); err != nil {
				subject.Next(err)
				s.status = Stop
				return t
			}
			return s.getNext(t)
		}
	}
	hr := s.getNum(t.Hour(), s.cron.hour)
	if v := hr - t.Hour(); v > 0 {
		t = t.Add(time.Duration(v) * jtime.Hour)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), 0, 0); err != nil {
			subject.Next(err)
			s.status = Stop
			return t
		}
	} else if v < 0 {
		t = t.Add(jtime.Day)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), 0, 0, 0); err != nil {
			subject.Next(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	min := s.getNum(t.Minute(), s.cron.minute)
	if v := min - t.Minute(); v > 0 {
		t = t.Add(time.Duration(v) * jtime.Minute)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), 0); err != nil {
			subject.Next(err)
			s.status = Stop
			return t
		}
	} else if v < 0 {
		t = t.Add(jtime.Hour)
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), 0, 0); err != nil {
			subject.Next(err)
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
		if t, err = s.getTime(t.Year(), int(t.Month()), t.Day(), t.Hour(), t.Minute(), 0); err != nil {
			subject.Next(err)
			s.status = Stop
			return t
		}
		return s.getNext(t)
	}
	return t
}

func (s *schedule) getNum(n int, bits uint64) int {
	num := 0
	if bits > 0 {
		var vBits uint64
		v := n
		vBits = 1 << v
		for vBits <= bits {
			if vBits&bits > 0 {
				num = v
				break
			} else {
				v++
				vBits = 1 << v
			}
		}
	}
	return num
}

func (s *schedule) getTime(year, month, day, hr, min, sec int) (time.Time, error) {
	var t time.Time
	if len(pack.Location) > 0 {
		if loc, err := time.LoadLocation(pack.Location); err != nil {
			return t, err
		} else {
			return jcast.TimeLoc(fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hr, min, sec), loc)
		}
	}
	return jcast.Time(fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hr, min, sec))
}
