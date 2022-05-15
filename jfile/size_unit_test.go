// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSizeUnit_ToInt64(t *testing.T) {
	fmt.Println(Byte.ToInt64())
}

func TestParseSizeUnit(t *testing.T) {
	errStr := "Error String"
	tests := []struct {
		input  string
		output SizeUnit
	}{
		{"Byte", Byte},
		{"Kb", Kb},
		{"KB", KB},
		{"Mb", Mb},
		{"MB", MB},
		{"Gb", Gb},
		{"GB", GB},
		{"Tb", Tb},
		{"TB", TB},
		{"Pb", Pb},
		{"PB", PB},
		{"Eb", Eb},
		{"EB", EB},
		{errStr, Unknown},
	}
	for _, test := range tests {
		if v, err := ParseSizeUnit(test.input); err != nil {
			if test.input != errStr {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}
