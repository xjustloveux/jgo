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
	if testTime, err = time.Parse(DateS, "2022-06-20 14:44:55 -0700 MST"); err != nil {
		assert.Error(t, err)
		return
	}
	tests := []struct {
		input  string
		output string
	}{
		{"%d", "20"},
		{"%dd", "20"},
		{"%ddd", testTime.Weekday().String()[:3]},
		{"%dddd", testTime.Weekday().String()},
	}
	for _, test := range tests {
		v := FormatString(test.input, testTime)
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}
