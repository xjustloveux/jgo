// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	testErr := "TEST ERROR:"
	if _, err := Load("error.json"); err == nil {
		t.Error(fmt.Sprint(testErr, " Load must be return error"))
	}
	if _, err := Load("../files/error.json"); err != nil {
		t.Error(err)
	}
}

func TestConvert(t *testing.T) {
	m1 := make(map[string]interface{})
	m2 := make(map[string]interface{})
	m3 := make(map[string]interface{})
	m1["str"] = "str"
	m1["bool"] = true
	m1["int"] = 7
	m1["float"] = 7.7
	m2["str"] = "str"
	m2["bool"] = true
	m2["int"] = 7
	m2["float"] = 7.7
	m3["str"] = "str"
	m3["bool"] = true
	m3["int"] = 7
	m3["float"] = 7.7
	m1["m2"] = m2
	m2["m3"] = m3
	type TestStruct struct {
		Str   string
		Bool  bool
		Int   int
		Float float64
		M2    map[string]interface{}
	}
	ts := TestStruct{}
	output := TestStruct{
		Str:   "str",
		Bool:  true,
		Int:   7,
		Float: 7.7,
		M2:    m2,
	}
	if err := Convert(m1, &ts); err != nil {
		t.Error(err)
	} else {
		output.M2["int"] = float64(7)
		output.M2["m3"].(map[string]interface{})["int"] = float64(7)
		msg := fmt.Sprintf("%v != %v", ts, output)
		assert.Equal(t, output, ts, msg)
	}
}
