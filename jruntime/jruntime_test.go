// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jruntime

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFuncName(t *testing.T) {
	null := "null"
	tests := []struct {
		input  string
		output interface{}
	}{
		{null, "TestGetFuncName"},
		{"/", "jruntime.TestGetFuncName"},
		{"*", "github.com/xjustloveux/jgo/jruntime.TestGetFuncName"},
	}
	for _, test := range tests {
		var v string
		if test.input == null {
			v = GetFuncName()
		} else {
			v = GetFuncName(test.input)
		}
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestGetCallerName(t *testing.T) {
	null := "null"
	tests := []struct {
		input  string
		output interface{}
	}{
		{null, "tRunner"},
		{"/", "testing.tRunner"},
		{"*", "testing.tRunner"},
	}
	for _, test := range tests {
		var v string
		if test.input == null {
			v = GetCallerName()
		} else {
			v = GetCallerName(test.input)
		}
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestGetCallerProgramName(t *testing.T) {
	null := "null"
	tests := []struct {
		input  string
		output interface{}
	}{
		{null, "testing.go"},
		{".", "go"},
	}
	for _, test := range tests {
		var v string
		if test.input == null {
			v = GetCallerProgramName()
		} else {
			v = GetCallerProgramName(test.input)
		}
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestGetCallerProgramLine(t *testing.T) {
	test := func() int {
		return GetCallerProgramLine()
	}
	input := 83
	output := test()
	assert.Equal(t, input, output, fmt.Sprintf("%v != %v", input, output))
}

func TestGetPkgName(t *testing.T) {
	input := "jruntime"
	output := GetPkgName()
	assert.Equal(t, input, output, fmt.Sprintf("%v != %v", input, output))
}

func TestGetCallerPkgName(t *testing.T) {
	input := "testing"
	output := GetCallerPkgName()
	assert.Equal(t, input, output, fmt.Sprintf("%v != %v", input, output))
}
