// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jtime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestParseTimeDuration(t *testing.T) {
	tests := []struct {
		input  string
		output interface{}
	}{
		{"Nanosecond", Nanosecond},
		{"Microsecond", Microsecond},
		{"Millisecond", Millisecond},
		{"Second", Second},
		{"Minute", Minute},
		{"Hour", Hour},
		{"Day", Day},
		{"NanoSecond", Nanosecond},
		{"MicroSecond", Microsecond},
		{"MilliSecond", Millisecond},
		{"SECOND", Second},
		{"Month", Unknown},
	}
	for _, test := range tests {
		if v, err := ParseTimeDuration(test.input); err != nil {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Error(t, err, msg)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestFormatString(t *testing.T) {
	var err error
	var testTime time.Time
	if testTime, err = time.Parse(DateS, "2022-06-20 14:44:55.012345678 -0700 MST"); err != nil {
		assert.Error(t, err)
		return
	}
	l, _ := time.LoadLocation("America/Phoenix")
	testTime = testTime.In(l)
	tests := []struct {
		input  string
		output string
	}{
		{"%d", "20"},
		{"%dd", "20"},
		{"%ddd", testTime.Weekday().String()[:3]},
		{"%dddd", testTime.Weekday().String()},
		{"%D", "171"},
		{"%DD", "171"},
		{"%DDD", "171"},
		{"%f", "0"},
		{"%ff", "01"},
		{"%fff", "012"},
		{"%ffff", "0123"},
		{"%F", ""},
		{"%FF", "01"},
		{"%FFF", "012"},
		{"%FFFF", "0123"},
		{"%g", "A.D."},
		{"%h", "2"},
		{"%hh", "02"},
		{"%H", "14"},
		{"%HH", "14"},
		{"%k", "2"},
		{"%kk", "02"},
		{"%K", "14"},
		{"%KK", "14"},
		{"%l", "America/Phoenix"},
		{"%m", "44"},
		{"%mm", "44"},
		{"%M", "6"},
		{"%MM", "06"},
		{"%MMM", "Jun"},
		{"%MMMM", "June"},
		{"%s", "55"},
		{"%ss", "55"},
		{"%t", "P"},
		{"%tt", "PM"},
		{"%w", "25"},
		{"%W", "1"},
		{"%y", "22"},
		{"%yy", "22"},
		{"%yyy", "2022"},
		{"%yyyy", "2022"},
		{"%z", "-7"},
		{"%zz", "-07"},
		{"%zzz", "-07:00"},
		{"%zzzz", "-25200"},
		{"%Z", "MST"},
	}
	for _, test := range tests {
		v := FormatString(test.input, testTime)
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}
