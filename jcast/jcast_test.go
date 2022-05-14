// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcast

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jtime"
	"testing"
	"time"
)

func TestVerifyPtr(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{1, 2, 3}
	now := time.Now()
	m := make(map[string]interface{})
	s := make([]interface{}, 0)
	m1 := make(map[string]interface{})
	m1["str"] = str
	m1["bo"] = bo
	m1["i"] = i
	m1["f"] = f
	m1["b"] = b
	m1["now"] = now
	m1["m"] = m
	m1["s"] = s
	m1["addrNow"] = &now
	s1 := []interface{}{str, bo, i, f, b, now, m, s, &now}
	tests := []struct {
		input  interface{}
		output interface{}
	}{
		{str, str},
		{bo, bo},
		{i, i},
		{f, f},
		{b, b},
		{now, now},
		{m1, m1},
		{s1, s1},
		{&now, now},
	}
	for _, test := range tests {
		v := VerifyPtr(test.input)
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestValue(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{1, 2, 3}
	now := time.Now()
	m := make(map[string]interface{})
	s := make([]interface{}, 0)
	m1 := make(map[string]interface{})
	m1["str"] = str
	m1["bo"] = bo
	m1["i"] = i
	m1["f"] = f
	m1["b"] = b
	m1["now"] = now
	m1["m"] = m
	m1["s"] = s
	m1["addrNow"] = &now
	m2 := make(map[string]interface{})
	m2["str"] = str
	m2["bo"] = bo
	m2["i"] = i
	m2["f"] = f
	m2["b"] = b
	m2["now"] = now
	m2["m"] = m
	m2["s"] = s
	m2["addrNow"] = now
	s1 := []interface{}{str, bo, i, f, b, now, m, s, &now}
	s2 := []interface{}{str, bo, i, f, b, now, m, s, now}
	tests := []struct {
		input  interface{}
		output interface{}
	}{
		{str, str},
		{bo, bo},
		{i, i},
		{f, f},
		{b, b},
		{now, now},
		{m1, m2},
		{s1, s2},
		{&now, now},
	}
	for _, test := range tests {
		v := Value(test.input)
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestTime(t *testing.T) {
	now := time.Now()
	str := now.Format(jtime.DateTime)
	num := now.Unix()
	tests := []struct {
		input  interface{}
		output time.Time
	}{
		{now, now},
		{str, now},
		{num, now},
		{&now, now},
	}
	for _, test := range tests {
		if v, err := Time(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output.Unix(), v.Unix(), msg)
		}
	}
}

func TestTimeLoc(t *testing.T) {
	now := time.Now()
	loc, _ := time.LoadLocation("America/Nipigon")
	lt := now.In(loc)
	str1 := now.Format(jtime.DateS)
	str2 := now.In(loc).Format(jtime.DateTime)
	num := now.Unix()
	tests := []struct {
		input  interface{}
		output time.Time
	}{
		{now, lt},
		{str1, lt},
		{str2, lt},
		{num, lt},
	}
	for _, test := range tests {
		if v, err := TimeLoc(test.input, loc); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output.Unix(), v.Unix(), msg)
		}
	}
}

func TestTimeString(t *testing.T) {
	now := time.Now()
	str1 := now.Format(jtime.DateS)
	str2 := now.Format(jtime.DateTime)
	num := now.Unix()
	tests := []struct {
		input  interface{}
		output string
	}{
		{now, str2},
		{str1, str2},
		{str2, str2},
		{num, str2},
	}
	for _, test := range tests {
		if v, err := TimeString(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestTimeFormatString(t *testing.T) {
	now := time.Now()
	str1 := now.Format(jtime.ANSIC)
	str2 := now.Format(jtime.UnixDate)
	str3 := now.Format(jtime.RubyDate)
	str4 := now.Format(jtime.RFC850)
	str5 := now.Format(jtime.RFC1123)
	str6 := now.Format(jtime.RFC1123Z)
	str7 := now.Format(jtime.RFC3339)
	str8 := now.Format(jtime.RFC3339Nano)
	str9 := now.Format(jtime.ISO8601)
	str10 := now.Format(jtime.DateTime)
	str11 := now.Format(jtime.DateS)
	num := now.Unix()
	tests := []struct {
		input  interface{}
		output string
	}{
		{now, str1},
		{str1, str1},
		{str2, str1},
		{str3, str1},
		{str4, str1},
		{str5, str1},
		{str6, str1},
		{str7, str1},
		{str8, str1},
		{str9, str1},
		{str10, str1},
		{str11, str1},
		{num, str1},
	}
	for _, test := range tests {
		if v, err := TimeFormatString(test.input, jtime.ANSIC); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestString(t *testing.T) {
	now := time.Now()
	tests := []struct {
		input  interface{}
		output string
	}{
		{"test", "test"},
		{true, "true"},
		{7, "7"},
		{7.7, "7.7"},
		{[]byte{116, 101, 115, 116}, "test"},
		{now, fmt.Sprintf("%v", now)},
	}
	for _, test := range tests {
		v := String(test.input)
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestBool(t *testing.T) {
	tests := []struct {
		input  interface{}
		output bool
	}{
		{"true", true},
		{true, true},
		{7, true},
		{7.7, true},
		{[]byte{116, 114, 117, 101}, true},
		{nil, false},
	}
	for _, test := range tests {
		if v, err := Bool(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInt(t *testing.T) {
	tests := []struct {
		input  interface{}
		output int
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Int(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInt8(t *testing.T) {
	tests := []struct {
		input  interface{}
		output int8
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Int8(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInt16(t *testing.T) {
	tests := []struct {
		input  interface{}
		output int16
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Int16(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInt32(t *testing.T) {
	tests := []struct {
		input  interface{}
		output int32
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Int32(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInt64(t *testing.T) {
	tests := []struct {
		input  interface{}
		output int64
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Int64(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestUint(t *testing.T) {
	tests := []struct {
		input  interface{}
		output uint
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Uint(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestUint8(t *testing.T) {
	tests := []struct {
		input  interface{}
		output uint8
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Uint8(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestUint16(t *testing.T) {
	tests := []struct {
		input  interface{}
		output uint16
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Uint16(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestUint32(t *testing.T) {
	tests := []struct {
		input  interface{}
		output uint32
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Uint32(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestUint64(t *testing.T) {
	tests := []struct {
		input  interface{}
		output uint64
	}{
		{"7", 7},
		{true, 1},
		{7, 7},
		{7.7, 7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Uint64(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestFloat32(t *testing.T) {
	tests := []struct {
		input  interface{}
		output float32
	}{
		{"7", 7},
		{"-7", -7},
		{"7.7", 7.7},
		{"-7.7", -7.7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7.7},
		{-7.7, -7.7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Float32(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestFloat64(t *testing.T) {
	tests := []struct {
		input  interface{}
		output float64
	}{
		{"7", 7},
		{"-7", -7},
		{"7.7", 7.7},
		{"-7.7", -7.7},
		{true, 1},
		{7, 7},
		{-7, -7},
		{7.7, 7.7},
		{-7.7, -7.7},
		{[]byte{55}, 7},
		{nil, 0},
	}
	for _, test := range tests {
		if v, err := Float64(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInterfaceMapInterface(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{1, 2, 3}
	now := time.Now()
	m := make(map[string]interface{})
	s1 := []interface{}{str, bo, i, f, b, now, m, &now}
	s2 := []interface{}{str, bo, i, f, b, now, m, now}
	m1 := make(map[string]interface{})
	m1["str"] = str
	m1["bo"] = bo
	m1["i"] = i
	m1["f"] = f
	m1["b"] = b
	m1["now"] = now
	m1["m"] = m
	m1["s"] = s1
	m1["addrNow"] = &now
	m2 := make(map[interface{}]interface{})
	m2["str"] = str
	m2["bo"] = bo
	m2["i"] = i
	m2["f"] = f
	m2["b"] = b
	m2["now"] = now
	m2["m"] = m
	m2["s"] = s2
	m2["addrNow"] = now
	tests := []struct {
		input  interface{}
		output map[interface{}]interface{}
	}{
		{m1, m2},
	}
	for _, test := range tests {
		if v, err := InterfaceMapInterface(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestStringMapInterface(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{1, 2, 3}
	now := time.Now()
	m := make(map[string]interface{})
	s1 := []interface{}{str, bo, i, f, b, now, m, &now}
	s2 := []interface{}{str, bo, i, f, b, now, m, now}
	m1 := make(map[string]interface{})
	m1["str"] = str
	m1["bo"] = bo
	m1["i"] = i
	m1["f"] = f
	m1["b"] = b
	m1["now"] = now
	m1["m"] = m
	m1["s"] = s1
	m1["addrNow"] = &now
	m2 := make(map[string]interface{})
	m2["str"] = str
	m2["bo"] = bo
	m2["i"] = i
	m2["f"] = f
	m2["b"] = b
	m2["now"] = now
	m2["m"] = m
	m2["s"] = s2
	m2["addrNow"] = now
	tests := []struct {
		input  interface{}
		output map[string]interface{}
	}{
		{m1, m2},
	}
	for _, test := range tests {
		if v, err := StringMapInterface(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceInterface(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{1, 2, 3}
	now := time.Now()
	m1 := make(map[string]interface{})
	m1["str"] = str
	m1["bo"] = bo
	m1["i"] = i
	m1["f"] = f
	m1["b"] = b
	m1["now"] = now
	m1["addrNow"] = &now
	m2 := make(map[string]interface{})
	m2["str"] = str
	m2["bo"] = bo
	m2["i"] = i
	m2["f"] = f
	m2["b"] = b
	m2["now"] = now
	m2["addrNow"] = now
	s1 := []interface{}{str, bo, i, f, b, now, m1, &now}
	s2 := []interface{}{str, bo, i, f, b, now, m2, now}
	tests := []struct {
		input  interface{}
		output []interface{}
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceInterface(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceString(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	now := time.Now()
	nowS := fmt.Sprintf("%v", now)
	m1 := make(map[string]interface{})
	m1["str"] = str
	m1["bo"] = bo
	m1["i"] = i
	m1["f"] = f
	m1["b"] = b
	m1["now"] = now
	m1["addrNow"] = &now
	m2 := make(map[string]interface{})
	m2["str"] = str
	m2["bo"] = bo
	m2["i"] = i
	m2["f"] = f
	m2["b"] = b
	m2["now"] = now
	m2["addrNow"] = now
	m2S := fmt.Sprintf("%v", m2)
	s1 := []interface{}{str, bo, i, f, b, now, m1, &now}
	s2 := []string{str, "true", "7", "7.7", "7", nowS, m2S, nowS}
	tests := []struct {
		input  interface{}
		output []string
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceString(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceBool(t *testing.T) {
	str := "true"
	bo := true
	i := 7
	f := 7.7
	b := []byte{116, 114, 117, 101}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []bool{true, true, true, true, true}
	tests := []struct {
		input  interface{}
		output []bool
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceBool(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceInt(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []int
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceInt(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceInt8(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int8{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []int8
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceInt8(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceInt16(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int16{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []int16
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceInt16(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceInt32(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int32{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []int32
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceInt32(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceInt64(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int64{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []int64
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceInt64(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceUint(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []uint
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceUint(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceUint8(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint8{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []uint8
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceUint8(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceUint16(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint16{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []uint16
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceUint16(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceUint32(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint32{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []uint32
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceUint32(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceUint64(t *testing.T) {
	str := "7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint64{7, 1, 7, 7, 7}
	tests := []struct {
		input  interface{}
		output []uint64
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceUint64(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceFloat32(t *testing.T) {
	str := "7.7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []float32{7.7, 1, 7, 7.7, 7}
	tests := []struct {
		input  interface{}
		output []float32
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceFloat32(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestSliceFloat64(t *testing.T) {
	str := "7.7"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []float64{7.7, 1, 7, 7.7, 7}
	tests := []struct {
		input  interface{}
		output []float64
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := SliceFloat64(test.input); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}
