// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestSchedule(t *testing.T) {
	testData := []struct {
		c  string
		o1 []int
		o2 []int
	}{
		{"7/17 7/17 3/7 7/17 1/3 ? 2037/37", []int{2037, 1, 7, 3, 7, 24}, []int{2074, 10, 24, 17, 58, 58}},
		{"3/23 3/23 3/3 ? 3/3 1/3 2033/23", []int{2033, 3, 3, 3, 3, 26}, []int{2079, 12, 28, 21, 49, 49}},
	}
	for _, data := range testData {
		var s *schedule
		sch := &SchInfo{Cron: data.c}
		if cron, err := ParseCronExpression(sch.Cron); err != nil {
			t.Error(err)
		} else {
			s = &schedule{
				name:           sch.Name,
				cronExpression: sch.Cron,
				cron:           cron,
				job:            sch.JobName,
				data:           sch.JobData,
				desc:           sch.Desc,
				status:         Run,
			}
		}
		if s != nil {
			var startTime time.Time
			var outTime1 time.Time
			var outTime2 time.Time
			var inTime1 time.Time
			var inTime2 time.Time
			var err error
			if startTime, err = s.getTime(minYear, 1, 1, 0, 0, 0); err != nil {
				t.Error(err)
			}
			if outTime1, err = s.getTime(data.o1[0], data.o1[1], data.o1[2], data.o1[3], data.o1[4], data.o1[5]); err != nil {
				t.Error(err)
			}
			if outTime2, err = s.getTime(data.o2[0], data.o2[1], data.o2[2], data.o2[3], data.o2[4], data.o2[5]); err != nil {
				t.Error(err)
			}
			count := 0
			s.toNext(startTime)
			for s.status == Run {
				inTime2 = s.next
				s.toNext(s.next)
				count++
				if count == 1 {
					inTime1 = s.next
				}
			}
			assert.Equal(t, outTime1, inTime1, fmt.Sprintf("%v != %v", inTime1, outTime1))
			assert.Equal(t, outTime2, inTime2, fmt.Sprintf("%v != %v", inTime2, outTime2))
		}
	}
}
