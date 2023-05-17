// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcast

import (
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jtime"
	"math"
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
	b := []byte{55}
	now := time.Now()
	im := make(map[string]interface{})
	im["str"] = str
	im["bo"] = bo
	im["i"] = i
	im["f"] = f
	im["b"] = b
	im["now"] = now
	im["NowP"] = &now
	om := make(map[string]interface{})
	om["str"] = str
	om["bo"] = bo
	om["i"] = i
	om["f"] = f
	om["b"] = b
	om["now"] = now
	om["NowP"] = now
	inS := []interface{}{str, bo, i, f, b, now, &now}
	outS := []interface{}{str, bo, i, f, b, now, now}
	mii := make(map[interface{}]interface{})
	mbi := make(map[bool]interface{})
	mi00i := make(map[int]interface{})
	mi08i := make(map[int8]interface{})
	mi16i := make(map[int16]interface{})
	mi32i := make(map[int32]interface{})
	mi64i := make(map[int64]interface{})
	mui00i := make(map[uint]interface{})
	mui08i := make(map[uint8]interface{})
	mui16i := make(map[uint16]interface{})
	mui32i := make(map[uint32]interface{})
	mui64i := make(map[uint64]interface{})
	mf32i := make(map[float32]interface{})
	mf64i := make(map[float64]interface{})
	mnii := make(map[interface{}]interface{})
	mn := make(map[time.Time]interface{})
	mii[str] = str
	mbi[bo] = bo
	mi00i[i] = i
	mi08i[int8(i)] = int8(i)
	mi16i[int16(i)] = int16(i)
	mi32i[int32(i)] = int32(i)
	mi64i[int64(i)] = int64(i)
	mui00i[uint(i)] = uint(i)
	mui08i[uint8(i)] = uint8(i)
	mui16i[uint16(i)] = uint16(i)
	mui32i[uint32(i)] = uint32(i)
	mui64i[uint64(i)] = uint64(i)
	mf32i[float32(i)] = float32(f)
	mf64i[float64(i)] = f
	mn[now] = now
	mnii[now] = now
	ss := []string{str}
	sb := []bool{bo}
	si := []int{i}
	si8 := []int8{int8(i)}
	si16 := []int16{int16(i)}
	si32 := []int32{int32(i)}
	si64 := []int64{int64(i)}
	sui := []uint{uint(i)}
	sui8 := []uint8{uint8(i)}
	sui16 := []uint16{uint16(i)}
	sui32 := []uint32{uint32(i)}
	sui64 := []uint64{uint64(i)}
	sf32 := []float32{float32(f)}
	sf64 := []float64{f}
	sn := []time.Time{now}
	snii := []interface{}{now}
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
		{&now, now},
		{im, om},
		{inS, outS},
		{mii, mii},
		{mbi, mbi},
		{mi00i, mi00i},
		{mi08i, mi08i},
		{mi16i, mi16i},
		{mi32i, mi32i},
		{mi64i, mi64i},
		{mui00i, mui00i},
		{mui08i, mui08i},
		{mui16i, mui16i},
		{mui32i, mui32i},
		{mui64i, mui64i},
		{mf32i, mf32i},
		{mf64i, mf64i},
		{mn, mnii},
		{ss, ss},
		{sb, sb},
		{si, si},
		{si8, si8},
		{si16, si16},
		{si32, si32},
		{si64, si64},
		{sui, sui},
		{sui8, sui8},
		{sui16, sui16},
		{sui32, sui32},
		{sui64, sui64},
		{sf32, sf32},
		{sf64, sf64},
		{sn, snii},
		{nil, nil},
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
	i := 7
	it := time.Unix(7, 0)
	tests := []struct {
		input  interface{}
		output time.Time
	}{
		{now, lt},
		{str1, lt},
		{str2, lt},
		{num, lt},
		{i, it},
		{int8(i), it},
		{int16(i), it},
		{int32(i), it},
		{int64(i), it},
		{uint(i), it},
		{uint8(i), it},
		{uint16(i), it},
		{uint32(i), it},
		{uint64(i), it},
		{[]byte(it.Format(jtime.DateS)), it},
		{"Error String", time.Time{}},
		{7.7, time.Time{}},
	}
	for _, test := range tests {
		if v, err := TimeLoc(test.input, loc); err != nil {
			if !test.output.IsZero() {
				t.Error(err)
			}
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
	inStr12 := now.Format(jtime.Date)
	outStr12, _ := time.Parse(jtime.Date, inStr12)
	inStr13 := now.Format(jtime.Time)
	outStr13, _ := time.Parse(jtime.Time, inStr13)
	num := now.Unix()
	errStr := "Error String"
	errTime := "9999-99-99 99:99:ZZ"
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
		{inStr12, outStr12.Format(jtime.ANSIC)},
		{inStr13, outStr13.Format(jtime.ANSIC)},
		{num, str1},
		{errStr, errStr},
		{errTime, errStr},
	}
	for _, test := range tests {
		if v, err := TimeFormatString(test.input, jtime.ANSIC); err != nil {
			if test.output != errStr {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestTimeLocTransform(t *testing.T) {
	now := time.Now()
	var err error
	var ts string
	if ts, err = TimeString(now); err != nil {
		t.Error(err)
		return
	}
	utc := fmt.Sprint(ts, " +0000 UTC")
	cst := fmt.Sprint(ts, " +0800 CST")
	// pst := fmt.Sprint(ts, " -0800 PST")
	var utcLoc, cstLoc /*, pstLoc*/ *time.Location
	if utcLoc, err = time.LoadLocation("UTC"); err != nil {
		t.Error(err)
		return
	}
	if cstLoc, err = time.LoadLocation("Asia/Taipei"); err != nil {
		t.Error(err)
		return
	}
	/*if pstLoc, err = time.LoadLocation("America/Los_Angeles"); err != nil {
		t.Error(err)
		return
	}*/
	tests := []struct {
		input  interface{}
		output string
	}{
		{"UTC", utc},
		{"Asia/Taipei", cst},
		// {"America/Los_Angeles", pst},
		{utcLoc, utc},
		{cstLoc, cst},
		// {pstLoc, pst},
		{"error", "error"},
		{7, "error"},
		{nil, "error"},
	}
	for _, test := range tests {
		var nt time.Time
		if nt, err = TimeLocTransform(now, test.input); err != nil {
			if test.output != "error" {
				t.Error(err)
			}
		} else {
			var nts string
			if nts, err = TimeFormatString(nt, jtime.DateS); err != nil {
				t.Error(err)
			} else {
				msg := fmt.Sprintf("%v != %v", nts, test.output)
				assert.Equal(t, test.output, nts, msg)
			}
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
		{xml.CharData{116, 101, 115, 116}, "test"},
		{xml.Comment{116, 101, 115, 116}, "test"},
		{now, fmt.Sprintf("%v", now)},
		{nil, ""},
	}
	for _, test := range tests {
		v := String(test.input)
		msg := fmt.Sprintf("%v != %v", v, test.output)
		assert.Equal(t, test.output, v, msg)
	}
}

func TestBool(t *testing.T) {
	i := 7
	f := 7.7
	errStr := "Error String"
	tests := []struct {
		input  interface{}
		output bool
	}{
		{"true", true},
		{true, true},
		{i, true},
		{int8(i), true},
		{int16(i), true},
		{int32(i), true},
		{int64(i), true},
		{uint(i), true},
		{uint8(i), true},
		{uint16(i), true},
		{uint32(i), true},
		{uint64(i), true},
		{float32(f), true},
		{f, true},
		{[]byte{116, 114, 117, 101}, true},
		{nil, false},
		{errStr, false},
		{time.Now(), false},
	}
	for _, test := range tests {
		if v, err := Bool(test.input); err != nil {
			if test.output {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInt(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	f := 7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output int
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, -7, -7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, errStr, 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, uint64(math.MaxInt64 + 1), math.MaxInt64},
	}
	for _, test := range tests {
		if v, err := Int(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Int must be return error"))
			}
		}
	}
}

func TestInt8(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	f := 7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output int8
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, -7, -7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, errStr, 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, -777, math.MinInt8},
		{true, 777, math.MaxInt8},
		{true, int16(-777), math.MinInt8},
		{true, int16(777), math.MaxInt8},
		{true, int32(-777), math.MinInt8},
		{true, int32(777), math.MaxInt8},
		{true, int64(-777), math.MinInt8},
		{true, int64(777), math.MaxInt8},
		{true, uint(777), math.MaxInt8},
		{true, uint8(177), math.MaxInt8},
		{true, uint16(777), math.MaxInt8},
		{true, uint32(777), math.MaxInt8},
		{true, uint64(777), math.MaxInt8},
		{true, float32(-777.7), math.MinInt8},
		{true, float32(777.7), math.MaxInt8},
		{true, -777.7, math.MinInt8},
		{true, 777.7, math.MaxInt8},
	}
	for _, test := range tests {
		if v, err := Int8(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Int8 must be return error"))
			}
		}
	}
}

func TestInt16(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	f := 7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output int16
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, -7, -7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, errStr, 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, -77777, math.MinInt16},
		{true, 77777, math.MaxInt16},
		{true, int32(-77777), math.MinInt16},
		{true, int32(77777), math.MaxInt16},
		{true, int64(-77777), math.MinInt16},
		{true, int64(77777), math.MaxInt16},
		{true, uint(77777), math.MaxInt16},
		{true, uint16(57777), math.MaxInt16},
		{true, uint32(77777), math.MaxInt16},
		{true, uint64(77777), math.MaxInt16},
		{true, float32(-77777.7), math.MinInt16},
		{true, float32(77777.7), math.MaxInt16},
		{true, -77777.7, math.MinInt16},
		{true, 77777.7, math.MaxInt16},
	}
	for _, test := range tests {
		if v, err := Int16(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Int16 must be return error"))
			}
		}
	}
}

func TestInt32(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	f := 7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output int32
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, -7, -7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, errStr, 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, -7777777777, math.MinInt32},
		{true, 7777777777, math.MaxInt32},
		{true, int64(-7777777777), math.MinInt32},
		{true, int64(7777777777), math.MaxInt32},
		{true, uint(7777777777), math.MaxInt32},
		{true, uint32(2777777777), math.MaxInt32},
		{true, uint64(7777777777), math.MaxInt32},
	}
	for _, test := range tests {
		if v, err := Int32(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Int32 must be return error"))
			}
		}
	}
}

func TestInt64(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	f := 7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output int64
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, -7, -7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, errStr, 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, uint64(math.MaxUint64), math.MaxInt64},
	}
	for _, test := range tests {
		if v, err := Int64(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Int64 must be return error"))
			}
		}
	}
}

func TestUint(t *testing.T) {
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		input  interface{}
		output uint
	}{
		{"7", 7},
		{true, 1},
		{false, 0},
		{i, 7},
		{int8(i), 7},
		{int16(i), 7},
		{int32(i), 7},
		{int64(i), 7},
		{uint(i), 7},
		{uint8(i), 7},
		{uint16(i), 7},
		{uint32(i), 7},
		{uint64(i), 7},
		{float32(f), 7},
		{f, 7},
		{[]byte{55}, 7},
		{nil, 0},
		{"-7", 99},
		{errStr, 99},
		{ni, 99},
		{int8(ni), 99},
		{int16(ni), 99},
		{int32(ni), 99},
		{int64(ni), 99},
		{float32(nf), 99},
		{nf, 99},
		{[]byte("-7"), 99},
		{[]byte(errStr), 99},
		{time.Now(), 99},
	}
	for _, test := range tests {
		if v, err := Uint(test.input); err != nil {
			if test.output != 99 {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestUint8(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output uint8
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, "-7", 0},
		{true, errStr, 0},
		{true, ni, 0},
		{true, int8(ni), 0},
		{true, int16(ni), 0},
		{true, int32(ni), 0},
		{true, int64(ni), 0},
		{true, float32(nf), 0},
		{true, nf, 0},
		{true, []byte("-7"), 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, 277, math.MaxUint8},
		{true, int16(777), math.MaxUint8},
		{true, int32(777), math.MaxUint8},
		{true, int64(777), math.MaxUint8},
		{true, uint(777), math.MaxUint8},
		{true, uint16(777), math.MaxUint8},
		{true, uint32(777), math.MaxUint8},
		{true, uint64(777), math.MaxUint8},
		{true, float32(777.7), math.MaxUint8},
		{true, 777.7, math.MaxUint8},
	}
	for _, test := range tests {
		if v, err := Uint8(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Uint8 must be return error"))
			}
		}
	}
}

func TestUint16(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output uint16
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, "-7", 0},
		{true, errStr, 0},
		{true, ni, 0},
		{true, int8(ni), 0},
		{true, int16(ni), 0},
		{true, int32(ni), 0},
		{true, int64(ni), 0},
		{true, float32(nf), 0},
		{true, nf, 0},
		{true, []byte("-7"), 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, 77777, math.MaxUint16},
		{true, int32(77777), math.MaxUint16},
		{true, int64(77777), math.MaxUint16},
		{true, uint(77777), math.MaxUint16},
		{true, uint32(77777), math.MaxUint16},
		{true, uint64(77777), math.MaxUint16},
		{true, float32(77777.7), math.MaxUint16},
		{true, 77777.7, math.MaxUint16},
	}
	for _, test := range tests {
		if v, err := Uint16(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Uint16 must be return error"))
			}
		}
	}
}

func TestUint32(t *testing.T) {
	testErr := "TEST ERROR:"
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		err    bool
		input  interface{}
		output uint32
	}{
		{false, "7", 7},
		{false, true, 1},
		{false, false, 0},
		{false, i, 7},
		{false, int8(i), 7},
		{false, int16(i), 7},
		{false, int32(i), 7},
		{false, int64(i), 7},
		{false, uint(i), 7},
		{false, uint8(i), 7},
		{false, uint16(i), 7},
		{false, uint32(i), 7},
		{false, uint64(i), 7},
		{false, float32(f), 7},
		{false, f, 7},
		{false, []byte{55}, 7},
		{false, nil, 0},
		{true, "-7", 0},
		{true, errStr, 0},
		{true, ni, 0},
		{true, int8(ni), 0},
		{true, int16(ni), 0},
		{true, int32(ni), 0},
		{true, int64(ni), 0},
		{true, float32(nf), 0},
		{true, nf, 0},
		{true, []byte("-7"), 0},
		{true, []byte(errStr), 0},
		{true, time.Now(), 0},
		{true, 7777777777, math.MaxUint32},
		{true, int64(7777777777), math.MaxUint32},
		{true, uint64(7777777777), math.MaxUint32},
	}
	for _, test := range tests {
		if v, err := Uint32(test.input); err != nil {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if !test.err {
				t.Error(err)
			}
		} else {
			assert.Equal(t, test.output, v, fmt.Sprintf("%v != %v", v, test.output))
			if test.err {
				t.Error(fmt.Sprint(testErr, " Uint32 must be return error"))
			}
		}
	}
}

func TestUint64(t *testing.T) {
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		input  interface{}
		output uint64
	}{
		{"7", 7},
		{true, 1},
		{false, 0},
		{i, 7},
		{int8(i), 7},
		{int16(i), 7},
		{int32(i), 7},
		{int64(i), 7},
		{uint(i), 7},
		{uint8(i), 7},
		{uint16(i), 7},
		{uint32(i), 7},
		{uint64(i), 7},
		{float32(f), 7},
		{f, 7},
		{[]byte{55}, 7},
		{nil, 0},
		{"-7", 99},
		{errStr, 99},
		{ni, 99},
		{int8(ni), 99},
		{int16(ni), 99},
		{int32(ni), 99},
		{int64(ni), 99},
		{float32(nf), 99},
		{nf, 99},
		{[]byte("-7"), 99},
		{[]byte(errStr), 99},
		{time.Now(), 99},
	}
	for _, test := range tests {
		if v, err := Uint64(test.input); err != nil {
			if test.output != 99 {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestFloat32(t *testing.T) {
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		input  interface{}
		output float32
	}{
		{"7", 7},
		{"-7", -7},
		{"7.7", 7.7},
		{"-7.7", -7.7},
		{true, 1},
		{false, 0},
		{i, 7},
		{int8(i), 7},
		{int16(i), 7},
		{int32(i), 7},
		{int64(i), 7},
		{uint(i), 7},
		{uint8(i), 7},
		{uint16(i), 7},
		{uint32(i), 7},
		{uint64(i), 7},
		{ni, -7},
		{float32(f), 7.7},
		{f, 7.7},
		{float32(nf), -7.7},
		{nf, -7.7},
		{[]byte{55}, 7},
		{nil, 0},
		{errStr, -1},
		{[]byte(errStr), -1},
		{time.Now(), -1},
	}
	for _, test := range tests {
		if v, err := Float32(test.input); err != nil {
			if test.output != -1 {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestFloat64(t *testing.T) {
	i := 7
	ni := -7
	f := 7.7
	nf := -7.7
	errStr := "Error String"
	tests := []struct {
		input  interface{}
		output float64
	}{
		{"7", 7},
		{"-7", -7},
		{"7.7", 7.7},
		{"-7.7", -7.7},
		{true, 1},
		{false, 0},
		{i, 7},
		{int8(i), 7},
		{int16(i), 7},
		{int32(i), 7},
		{int64(i), 7},
		{uint(i), 7},
		{uint8(i), 7},
		{uint16(i), 7},
		{uint32(i), 7},
		{uint64(i), 7},
		{ni, -7},
		{float32(f), 7.7},
		{f, 7.7},
		{float32(nf), -7.7},
		{nf, -7.7},
		{[]byte{55}, 7},
		{nil, 0},
		{errStr, -1},
		{[]byte(errStr), -1},
		{time.Now(), -1},
	}
	for _, test := range tests {
		if v, err := Float64(test.input); err != nil {
			if test.output != -1 {
				t.Error(err)
			}
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
		{now, nil},
	}
	for _, test := range tests {
		if v, err := InterfaceMapInterface(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
		{now, nil},
	}
	for _, test := range tests {
		if v, err := StringMapInterface(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestStringMapString(t *testing.T) {
	str := "Str"
	bo := true
	i := 7
	f := 7.7
	b := []byte{55}
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
	m2 := make(map[string]string)
	m2["str"] = str
	m2["bo"] = "true"
	m2["i"] = "7"
	m2["f"] = "7.7"
	m2["b"] = "7"
	m2["now"] = fmt.Sprintf("%v", now)
	m2["m"] = fmt.Sprintf("%v", m)
	m2["s"] = fmt.Sprintf("%v", s2)
	m2["addrNow"] = fmt.Sprintf("%v", now)
	tests := []struct {
		input  interface{}
		output map[string]string
	}{
		{m1, m2},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := StringMapString(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceInterface(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceString(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []bool{true, true, true, true, true}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []bool
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceBool(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []int
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceInt(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int8{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []int8
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceInt8(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int16{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []int16
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceInt16(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int32{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []int32
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceInt32(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []int64{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []int64
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceInt64(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []uint
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceUint(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint8{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []uint8
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceUint8(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint16{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []uint16
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceUint16(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint32{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []uint32
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceUint32(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []uint64{7, 1, 7, 7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []uint64
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceUint64(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []float32{7.7, 1, 7, 7.7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []float32
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceFloat32(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
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
	now := time.Now()
	s1 := []interface{}{str, bo, i, f, b}
	s2 := []float64{7.7, 1, 7, 7.7, 7}
	s3 := []interface{}{"Error String"}
	tests := []struct {
		input  interface{}
		output []float64
	}{
		{s1, s2},
		{s3, nil},
		{now, nil},
	}
	for _, test := range tests {
		if v, err := SliceFloat64(test.input); err != nil {
			if test.output != nil {
				t.Error(err)
			}
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}
