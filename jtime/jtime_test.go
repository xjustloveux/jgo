// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jtime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
