// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcast

import (
	"encoding/xml"
	"fmt"
	"github.com/xjustloveux/jgo/jruntime"
	"github.com/xjustloveux/jgo/jtime"
	"math"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	errorInterfaceTo = jError("unable to cast %#v of type %T to %s")
	errorParseTime   = jError("unable to parse time: %s")
	errorNegative    = jError("unable to cast negative value")
	errorOutRange    = jError("parsing %q: value out of range")
)

const (
	pkgName = "jcast"
)

var timeFormat = []jtime.Format{
	{jtime.ANSIC, true},
	{jtime.UnixDate, true},
	{jtime.RubyDate, true},
	{jtime.RFC822, true},
	{jtime.RFC822Z, true},
	{jtime.RFC850, true},
	{jtime.RFC1123, true},
	{jtime.RFC1123Z, true},
	{jtime.RFC3339, true},
	{jtime.RFC3339Nano, true},
	{jtime.Kitchen, false},
	{jtime.Stamp, false},
	{jtime.StampMilli, false},
	{jtime.StampMicro, false},
	{jtime.StampNano, false},
	{jtime.ISO8601, true},
	{jtime.DateTime, true},
	{jtime.DateS, false},
	{jtime.Date, true},
	{jtime.Time, false},
}

// VerifyPtr verify pointer
func VerifyPtr(i interface{}) interface{} {
	if i == nil {
		return nil
	}
	if t := reflect.TypeOf(i); t.Kind() != reflect.Ptr {
		return i
	}
	v := reflect.ValueOf(i)
	for v.Kind() == reflect.Ptr && !v.IsNil() {
		v = v.Elem()
	}
	return v.Interface()
}

// Value check pointer, map, slice type and returns new interface value
func Value(i interface{}) interface{} {
	switch v := VerifyPtr(i).(type) {
	case map[interface{}]interface{}:
		m := make(map[interface{}]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[string]interface{}:
		m := make(map[string]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[bool]interface{}:
		m := make(map[bool]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[int]interface{}:
		m := make(map[int]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[int8]interface{}:
		m := make(map[int8]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[int16]interface{}:
		m := make(map[int16]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[int32]interface{}:
		m := make(map[int32]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[int64]interface{}:
		m := make(map[int64]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[uint]interface{}:
		m := make(map[uint]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[uint8]interface{}:
		m := make(map[uint8]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[uint16]interface{}:
		m := make(map[uint16]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[uint32]interface{}:
		m := make(map[uint32]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[uint64]interface{}:
		m := make(map[uint64]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[float32]interface{}:
		m := make(map[float32]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case map[float64]interface{}:
		m := make(map[float64]interface{})
		for mk, mv := range v {
			m[mk] = Value(mv)
		}
		return m
	case []interface{}:
		s := make([]interface{}, len(v))
		for idx, sv := range v {
			s[idx] = Value(sv)
		}
		return s
	case []string:
		s := make([]string, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []bool:
		s := make([]bool, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []int:
		s := make([]int, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []int8:
		s := make([]int8, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []int16:
		s := make([]int16, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []int32:
		s := make([]int32, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []int64:
		s := make([]int64, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []uint:
		s := make([]uint, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []uint8:
		s := make([]uint8, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []uint16:
		s := make([]uint16, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []uint32:
		s := make([]uint32, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []uint64:
		s := make([]uint64, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []float32:
		s := make([]float32, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	case []float64:
		s := make([]float64, len(v))
		for idx, sv := range v {
			s[idx] = sv
		}
		return s
	default:
		if v == nil {
			return nil
		}
		if reflect.TypeOf(v).Kind() == reflect.Map {
			if m, err := InterfaceMapInterface(i); err == nil {
				return m
			}
		} else if reflect.TypeOf(v).Kind() == reflect.Slice {
			if s, err := SliceInterface(i); err == nil {
				return s
			}
		}
		return v
	}
}

// Time interface to time.Time
func Time(i interface{}) (time.Time, error) {
	return TimeLoc(i, time.Local)
}

// TimeLoc interface to time.Time
func TimeLoc(i interface{}, loc *time.Location) (time.Time, error) {
	switch v := VerifyPtr(i).(type) {
	case time.Time:
		return v.In(loc), nil
	case string:
		return strToTime(v, loc)
	case int:
		return time.Unix(int64(v), 0), nil
	case int8:
		return time.Unix(int64(v), 0), nil
	case int16:
		return time.Unix(int64(v), 0), nil
	case int32:
		return time.Unix(int64(v), 0), nil
	case int64:
		return time.Unix(v, 0), nil
	case uint:
		return time.Unix(int64(v), 0), nil
	case uint8:
		return time.Unix(int64(v), 0), nil
	case uint16:
		return time.Unix(int64(v), 0), nil
	case uint32:
		return time.Unix(int64(v), 0), nil
	case uint64:
		return time.Unix(int64(v), 0), nil
	case []byte:
		return strToTime(string(v), loc)
	default:
		return time.Time{}, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// TimeString interface to time string
func TimeString(i interface{}) (string, error) {
	return TimeFormatString(i, jtime.DateTime)
}

// TimeFormatString interface to time string
func TimeFormatString(i interface{}, format string) (string, error) {
	switch v := VerifyPtr(i).(type) {
	case time.Time:
		return v.Format(format), nil
	default:
		if t, err := Time(i); err != nil {
			return "", err
		} else {
			return t.Format(format), nil
		}
	}
}

// String interface to string
func String(i interface{}) string {
	switch v := VerifyPtr(i).(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case xml.CharData:
		return string(v)
	case xml.Comment:
		return string(v)
	case nil:
		return ""
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Bool interface to bool
func Bool(i interface{}) (bool, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		return strconv.ParseBool(v)
	case bool:
		return v, nil
	case int:
		if v != 0 {
			return true, nil
		}
	case int8:
		if v != 0 {
			return true, nil
		}
	case int16:
		if v != 0 {
			return true, nil
		}
	case int32:
		if v != 0 {
			return true, nil
		}
	case int64:
		if v != 0 {
			return true, nil
		}
	case uint:
		if v != 0 {
			return true, nil
		}
	case uint8:
		if v != 0 {
			return true, nil
		}
	case uint16:
		if v != 0 {
			return true, nil
		}
	case uint32:
		if v != 0 {
			return true, nil
		}
	case uint64:
		if v != 0 {
			return true, nil
		}
	case float32:
		if v != 0 {
			return true, nil
		}
	case float64:
		if v != 0 {
			return true, nil
		}
	case []byte:
		return strconv.ParseBool(string(v))
	case nil:
	default:
		return false, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
	return false, nil
}

// Int interface to int
func Int(i interface{}) (int, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseInt(v, 0, strconv.IntSize); err == nil {
			return int(pv), nil
		} else {
			return int(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return v, nil
	case int8:
		return int(v), nil
	case int16:
		return int(v), nil
	case int32:
		return int(v), nil
	case int64:
		if strconv.IntSize == 32 {
			if v < math.MinInt32 {
				return math.MinInt32, errorf(errorOutRange, v)
			}
			if v > math.MaxInt32 {
				return math.MaxInt32, errorf(errorOutRange, v)
			}
		}
		return int(v), nil
	case uint:
		return int(v), nil
	case uint8:
		return int(v), nil
	case uint16:
		return int(v), nil
	case uint32:
		if strconv.IntSize == 32 && v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		return int(v), nil
	case uint64:
		if strconv.IntSize == 32 && v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		if strconv.IntSize == 64 && v > math.MaxInt64 {
			return math.MaxInt64, errorf(errorOutRange, v)
		}
		return int(v), nil
	case float32:
		return int(v), nil
	case float64:
		return int(v), nil
	case []byte:
		if pv, err := strconv.ParseInt(string(v), 0, 0); err == nil {
			return int(pv), nil
		} else {
			return int(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Int8 interface to int8
func Int8(i interface{}) (int8, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseInt(v, 0, 8); err == nil {
			return int8(pv), nil
		} else {
			return int8(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < math.MinInt8 {
			return math.MinInt8, errorf(errorOutRange, v)
		}
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case int8:
		return v, nil
	case int16:
		if v < math.MinInt8 {
			return math.MinInt8, errorf(errorOutRange, v)
		}
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case int32:
		if v < math.MinInt8 {
			return math.MinInt8, errorf(errorOutRange, v)
		}
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case int64:
		if v < math.MinInt8 {
			return math.MinInt8, errorf(errorOutRange, v)
		}
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case uint:
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case uint8:
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case uint16:
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case uint32:
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case uint64:
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case float32:
		if v < math.MinInt8 {
			return math.MinInt8, errorf(errorOutRange, v)
		}
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case float64:
		if v < math.MinInt8 {
			return math.MinInt8, errorf(errorOutRange, v)
		}
		if v > math.MaxInt8 {
			return math.MaxInt8, errorf(errorOutRange, v)
		}
		return int8(v), nil
	case []byte:
		if pv, err := strconv.ParseInt(string(v), 0, 8); err == nil {
			return int8(pv), nil
		} else {
			return int8(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Int16 interface to int16
func Int16(i interface{}) (int16, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseInt(v, 0, 16); err == nil {
			return int16(pv), nil
		} else {
			return int16(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < math.MinInt16 {
			return math.MinInt16, errorf(errorOutRange, v)
		}
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case int8:
		return int16(v), nil
	case int16:
		return v, nil
	case int32:
		if v < math.MinInt16 {
			return math.MinInt16, errorf(errorOutRange, v)
		}
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case int64:
		if v < math.MinInt16 {
			return math.MinInt16, errorf(errorOutRange, v)
		}
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case uint:
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case uint8:
		return int16(v), nil
	case uint16:
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case uint32:
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case uint64:
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case float32:
		if v < math.MinInt16 {
			return math.MinInt16, errorf(errorOutRange, v)
		}
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case float64:
		if v < math.MinInt16 {
			return math.MinInt16, errorf(errorOutRange, v)
		}
		if v > math.MaxInt16 {
			return math.MaxInt16, errorf(errorOutRange, v)
		}
		return int16(v), nil
	case []byte:
		if pv, err := strconv.ParseInt(string(v), 0, 16); err == nil {
			return int16(pv), nil
		} else {
			return int16(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Int32 interface to int32
func Int32(i interface{}) (int32, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseInt(v, 0, 32); err == nil {
			return int32(pv), nil
		} else {
			return int32(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < math.MinInt32 {
			return math.MinInt32, errorf(errorOutRange, v)
		}
		if v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		return int32(v), nil
	case int8:
		return int32(v), nil
	case int16:
		return int32(v), nil
	case int32:
		return v, nil
	case int64:
		if v < math.MinInt32 {
			return math.MinInt32, errorf(errorOutRange, v)
		}
		if v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		return int32(v), nil
	case uint:
		if v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		return int32(v), nil
	case uint8:
		return int32(v), nil
	case uint16:
		return int32(v), nil
	case uint32:
		if v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		return int32(v), nil
	case uint64:
		if v > math.MaxInt32 {
			return math.MaxInt32, errorf(errorOutRange, v)
		}
		return int32(v), nil
	case float32:
		return int32(v), nil
	case float64:
		return int32(v), nil
	case []byte:
		if pv, err := strconv.ParseInt(string(v), 0, 32); err == nil {
			return int32(pv), nil
		} else {
			return int32(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Int64 interface to int64
func Int64(i interface{}) (int64, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseInt(v, 0, 64); err == nil {
			return pv, nil
		} else {
			return pv, err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return int64(v), nil
	case int8:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case uint:
		if v > math.MaxInt64 {
			return math.MaxInt64, errorf(errorOutRange, v)
		}
		return int64(v), nil
	case uint8:
		return int64(v), nil
	case uint16:
		return int64(v), nil
	case uint32:
		return int64(v), nil
	case uint64:
		if v > math.MaxInt64 {
			return math.MaxInt64, errorf(errorOutRange, v)
		}
		return int64(v), nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case []byte:
		if pv, err := strconv.ParseInt(string(v), 0, 64); err == nil {
			return pv, nil
		} else {
			return pv, err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Uint interface to uint
func Uint(i interface{}) (uint, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseUint(v, 0, strconv.IntSize); err == nil {
			return uint(pv), nil
		} else {
			return uint(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint(v), nil
	case int8:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint(v), nil
	case int16:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint(v), nil
	case int32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint(v), nil
	case int64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if strconv.IntSize == 32 && v > math.MaxUint32 {
			return math.MaxUint32, errorf(errorOutRange, v)
		}
		return uint(v), nil
	case uint:
		return v, nil
	case uint8:
		return uint(v), nil
	case uint16:
		return uint(v), nil
	case uint32:
		return uint(v), nil
	case uint64:
		if strconv.IntSize == 32 && v > math.MaxUint32 {
			return math.MaxUint32, errorf(errorOutRange, v)
		}
		return uint(v), nil
	case float32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint(v), nil
	case float64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint(v), nil
	case []byte:
		if pv, err := strconv.ParseUint(string(v), 0, strconv.IntSize); err == nil {
			return uint(pv), nil
		} else {
			return uint(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Uint8 interface to uint8
func Uint8(i interface{}) (uint8, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseUint(v, 0, 8); err == nil {
			return uint8(pv), nil
		} else {
			return uint8(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case int8:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint8(v), nil
	case int16:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case int32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case int64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case uint:
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case uint8:
		return v, nil
	case uint16:
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case uint32:
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case uint64:
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case float32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case float64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint8 {
			return math.MaxUint8, errorf(errorOutRange, v)
		}
		return uint8(v), nil
	case []byte:
		if pv, err := strconv.ParseUint(string(v), 0, 8); err == nil {
			return uint8(pv), nil
		} else {
			return uint8(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Uint16 interface to uint16
func Uint16(i interface{}) (uint16, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseUint(v, 0, 16); err == nil {
			return uint16(pv), nil
		} else {
			return uint16(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case int8:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint16(v), nil
	case int16:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint16(v), nil
	case int32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case int64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case uint:
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case uint8:
		return uint16(v), nil
	case uint16:
		return v, nil
	case uint32:
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case uint64:
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case float32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case float64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint16 {
			return math.MaxUint16, errorf(errorOutRange, v)
		}
		return uint16(v), nil
	case []byte:
		if pv, err := strconv.ParseUint(string(v), 0, 16); err == nil {
			return uint16(pv), nil
		} else {
			return uint16(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Uint32 interface to uint32
func Uint32(i interface{}) (uint32, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseUint(v, 0, 32); err == nil {
			return uint32(pv), nil
		} else {
			return uint32(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint32 {
			return math.MaxUint32, errorf(errorOutRange, v)
		}
		return uint32(v), nil
	case int8:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint32(v), nil
	case int16:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint32(v), nil
	case int32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint32(v), nil
	case int64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		if v > math.MaxUint32 {
			return math.MaxUint32, errorf(errorOutRange, v)
		}
		return uint32(v), nil
	case uint:
		if v > math.MaxUint32 {
			return math.MaxUint32, errorf(errorOutRange, v)
		}
		return uint32(v), nil
	case uint8:
		return uint32(v), nil
	case uint16:
		return uint32(v), nil
	case uint32:
		return v, nil
	case uint64:
		if v > math.MaxUint32 {
			return math.MaxUint32, errorf(errorOutRange, v)
		}
		return uint32(v), nil
	case float32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint32(v), nil
	case float64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint32(v), nil
	case []byte:
		if pv, err := strconv.ParseUint(string(v), 0, 32); err == nil {
			return uint32(pv), nil
		} else {
			return uint32(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Uint64 interface to uint64
func Uint64(i interface{}) (uint64, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseUint(v, 0, 64); err == nil {
			return pv, nil
		} else {
			return pv, err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case int8:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case int16:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case int32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case int64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case uint:
		return uint64(v), nil
	case uint8:
		return uint64(v), nil
	case uint16:
		return uint64(v), nil
	case uint32:
		return uint64(v), nil
	case uint64:
		return v, nil
	case float32:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case float64:
		if v < 0 {
			return 0, errors(errorNegative)
		}
		return uint64(v), nil
	case []byte:
		if pv, err := strconv.ParseUint(string(v), 0, 64); err == nil {
			return pv, nil
		} else {
			return pv, err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Float32 interface to float32
func Float32(i interface{}) (float32, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseFloat(v, 32); err == nil {
			return float32(pv), nil
		} else {
			return float32(pv), err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return float32(v), nil
	case int8:
		return float32(v), nil
	case int16:
		return float32(v), nil
	case int32:
		return float32(v), nil
	case int64:
		return float32(v), nil
	case uint:
		return float32(v), nil
	case uint8:
		return float32(v), nil
	case uint16:
		return float32(v), nil
	case uint32:
		return float32(v), nil
	case uint64:
		return float32(v), nil
	case float32:
		return v, nil
	case float64:
		return float32(v), nil
	case []byte:
		if pv, err := strconv.ParseFloat(string(v), 32); err == nil {
			return float32(pv), nil
		} else {
			return float32(pv), err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// Float64 interface to float64
func Float64(i interface{}) (float64, error) {
	switch v := VerifyPtr(i).(type) {
	case string:
		if pv, err := strconv.ParseFloat(v, 64); err == nil {
			return pv, nil
		} else {
			return pv, err
		}
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int:
		return float64(v), nil
	case int8:
		return float64(v), nil
	case int16:
		return float64(v), nil
	case int32:
		return float64(v), nil
	case int64:
		return float64(v), nil
	case uint:
		return float64(v), nil
	case uint8:
		return float64(v), nil
	case uint16:
		return float64(v), nil
	case uint32:
		return float64(v), nil
	case uint64:
		return float64(v), nil
	case float32:
		return strconv.ParseFloat(String(v), 64)
	case float64:
		return v, nil
	case []byte:
		if pv, err := strconv.ParseFloat(string(v), 64); err == nil {
			return pv, nil
		} else {
			return pv, err
		}
	case nil:
		return 0, nil
	default:
		return 0, errorf(errorInterfaceTo, i, i, strings.ToLower(jruntime.GetFuncName()))
	}
}

// InterfaceMapInterface interface to map[interface{}]interface{}
func InterfaceMapInterface(i interface{}) (map[interface{}]interface{}, error) {
	v := VerifyPtr(i)
	if reflect.TypeOf(v).Kind() == reflect.Map {
		mv := reflect.ValueOf(i)
		m := make(map[interface{}]interface{})
		for _, key := range mv.MapKeys() {
			m[Value(key.Interface())] = Value(mv.MapIndex(key).Interface())
		}
		return m, nil
	} else {
		return nil, errorf(errorInterfaceTo, i, i, "map[interface{}]interface{}")
	}
}

// StringMapInterface interface to map[string]interface{}
func StringMapInterface(i interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if imi, err := InterfaceMapInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "map[string]interface{}")
	} else {
		for k, v := range imi {
			m[String(k)] = v
		}
	}
	return m, nil
}

// StringMapString interface to map[string]string
func StringMapString(i interface{}) (map[string]string, error) {
	m := make(map[string]string)
	if imi, err := InterfaceMapInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "map[string]string")
	} else {
		for k, v := range imi {
			m[String(k)] = String(v)
		}
	}
	return m, nil
}

// SliceInterface interface to []interface{}
func SliceInterface(i interface{}) ([]interface{}, error) {
	v := VerifyPtr(i)
	if reflect.TypeOf(v).Kind() == reflect.Slice {
		sv := reflect.ValueOf(i)
		s := make([]interface{}, sv.Len())
		for c := 0; c < sv.Len(); c++ {
			s[c] = Value(sv.Index(c).Interface())
		}
		return s, nil
	} else {
		return nil, errorf(errorInterfaceTo, i, i, "[]interface{}")
	}
}

// SliceString interface to []string
func SliceString(i interface{}) ([]string, error) {
	if si, err := SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]string")
	} else {
		s := make([]string, len(si))
		for idx, v := range si {
			s[idx] = String(v)
		}
		return s, nil
	}
}

// SliceBool interface to []bool
func SliceBool(i interface{}) ([]bool, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]bool")
	}
	s := make([]bool, len(si))
	for idx, v := range si {
		if s[idx], err = Bool(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]bool")
		}
	}
	return s, nil
}

// SliceInt interface to []int
func SliceInt(i interface{}) ([]int, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]int")
	}
	s := make([]int, len(si))
	for idx, v := range si {
		if s[idx], err = Int(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]int")
		}
	}
	return s, nil
}

// SliceInt8 interface to []int8
func SliceInt8(i interface{}) ([]int8, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]int8")
	}
	s := make([]int8, len(si))
	for idx, v := range si {
		if s[idx], err = Int8(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]int8")
		}
	}
	return s, nil
}

// SliceInt16 interface to []int16
func SliceInt16(i interface{}) ([]int16, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]int16")
	}
	s := make([]int16, len(si))
	for idx, v := range si {
		if s[idx], err = Int16(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]int16")
		}
	}
	return s, nil
}

// SliceInt32 interface to []int32
func SliceInt32(i interface{}) ([]int32, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]int32")
	}
	s := make([]int32, len(si))
	for idx, v := range si {
		if s[idx], err = Int32(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]int32")
		}
	}
	return s, nil
}

// SliceInt64 interface to []int64
func SliceInt64(i interface{}) ([]int64, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]int64")
	}
	s := make([]int64, len(si))
	for idx, v := range si {
		if s[idx], err = Int64(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]int64")
		}
	}
	return s, nil
}

// SliceUint interface to []uint
func SliceUint(i interface{}) ([]uint, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]uint")
	}
	s := make([]uint, len(si))
	for idx, v := range si {
		if s[idx], err = Uint(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]uint")
		}
	}
	return s, nil
}

// SliceUint8 interface to []uint8
func SliceUint8(i interface{}) ([]uint8, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]uint8")
	}
	s := make([]uint8, len(si))
	for idx, v := range si {
		if s[idx], err = Uint8(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]uint8")
		}
	}
	return s, nil
}

// SliceUint16 interface to []uint16
func SliceUint16(i interface{}) ([]uint16, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]uint16")
	}
	s := make([]uint16, len(si))
	for idx, v := range si {
		if s[idx], err = Uint16(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]uint16")
		}
	}
	return s, nil
}

// SliceUint32 interface to []uint32
func SliceUint32(i interface{}) ([]uint32, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]uint32")
	}
	s := make([]uint32, len(si))
	for idx, v := range si {
		if s[idx], err = Uint32(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]uint32")
		}
	}
	return s, nil
}

// SliceUint64 interface to []uint64
func SliceUint64(i interface{}) ([]uint64, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]uint64")
	}
	s := make([]uint64, len(si))
	for idx, v := range si {
		if s[idx], err = Uint64(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]uint64")
		}
	}
	return s, nil
}

// SliceFloat32 interface to []float32
func SliceFloat32(i interface{}) ([]float32, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]float32")
	}
	s := make([]float32, len(si))
	for idx, v := range si {
		if s[idx], err = Float32(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]float32")
		}
	}
	return s, nil
}

// SliceFloat64 interface to []float64
func SliceFloat64(i interface{}) ([]float64, error) {
	var err error
	var si []interface{}
	if si, err = SliceInterface(i); err != nil {
		return nil, errorf(errorInterfaceTo, i, i, "[]float64")
	}
	s := make([]float64, len(si))
	for idx, v := range si {
		if s[idx], err = Float64(v); err != nil {
			return nil, errorf(errorInterfaceTo, i, i, "[]float64")
		}
	}
	return s, nil
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}

func strToTime(s string, loc *time.Location) (time.Time, error) {
	two := "{2}"
	three := "{3}"
	four := "{4}"
	six := "{6}"
	o := "[^A-Za-z0-9_]"
	o2 := fmt.Sprint(o, two)
	on := "[^0-9]"
	n := "[0-9]"
	n2 := fmt.Sprint(n, two)
	n3 := fmt.Sprint(n, three)
	n4 := fmt.Sprint(n, four)
	n6 := fmt.Sprint(n, six)
	n2o := fmt.Sprint(n2, o)
	n4o := fmt.Sprint(n4, o)
	en := "[A-Za-z]"
	en2 := fmt.Sprint(en, two)
	en3 := fmt.Sprint(en, three)
	en3o := fmt.Sprint(en3, o)
	en3o2 := fmt.Sprint(en3o, two)
	date := fmt.Sprint(n4o, n2o, n2)
	dateo := fmt.Sprint(date, o)
	tm := fmt.Sprint(n2o, n2o, n2)
	tmo := fmt.Sprint(tm, o)
	for _, f := range timeFormat {
		pattern := ""
		ck := false
		switch f.F {
		case jtime.ANSIC: // Mon Jan _2 15:04:05 2006
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o, en3o, n2o, tmo, n4)
		case jtime.UnixDate: // Mon Jan _2 15:04:05 MST 2006
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o, en3o, n2o, tmo, en3o, n4)
		case jtime.RubyDate: // Mon Jan 02 15:04:05 -0700 2006
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o, en3o, n2o, tmo, o, n4o, n4)
		case jtime.RFC822: // 02 Jan 06 15:04 MST
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", n2o, en3o, n2o, n2o, n2o, en3)
		case jtime.RFC822Z: // 02 Jan 06 15:04 -0700
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", n2o, en3o, n2o, n2o, n2o, o, n4)
		case jtime.RFC850: // Monday, 02-Jan-06 15:04:05 MST
			ck = len(s) >= len(f.F)
			pattern = fmt.Sprint("^", en, "+", o2, n2o, en3o, n2o, tmo, en3)
		case jtime.RFC1123: // Mon, 02 Jan 2006 15:04:05 MST
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o2, n2o, en3o, n4o, tmo, en3)
		case jtime.RFC1123Z: // Mon, 02 Jan 2006 15:04:05 -0700
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o2, n2o, en3o, n4o, tmo, o, n4)
		case jtime.RFC3339: // 2006-01-02T15:04:05Z07:00
			ck = len(s) >= (len(f.F) - 5)
			pattern = fmt.Sprint("^", date, en, tm, on)
		case jtime.RFC3339Nano: // 2006-01-02T15:04:05.999999999Z07:00
			ck = len(s) >= (len(f.F) - 13)
			pattern = fmt.Sprint("^", date, en, tmo, n, "+", on)
		case jtime.Kitchen: // 3:04PM
			ck = len(s) >= (len(f.F) - 1)
			pattern = fmt.Sprint("^", n, "+", o, n2, en2)
		case jtime.Stamp: // Jan _2 15:04:05
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o, n2o, tm)
		case jtime.StampMilli: // Jan _2 15:04:05.000
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o, n2o, tmo, n3)
		case jtime.StampMicro: // Jan _2 15:04:05.000000
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", en3o, n2o, tmo, n6)
		case jtime.StampNano: // Jan _2 15:04:05.000000000
			ck = len(s) >= (len(f.F) - 8)
			pattern = fmt.Sprint("^", en3o, n2o, tmo, n, "+")
		case jtime.ISO8601: // 2006-01-02T15:04:05
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", date, en, tm)
		case jtime.DateTime: // 2006-01-02 15:04:05
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", dateo, tm)
		case jtime.DateS: // 2006-01-02 15:04:05 -0700 MST
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", dateo, tmo, o, n4o, en3)
		case jtime.Date: // 2006-01-02
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", date)
		case jtime.Time:
			ck = len(f.F) == len(s)
			pattern = fmt.Sprint("^", tm)
		}
		if pattern != "" && ck {
			if b, err := regexp.MatchString(pattern, s); err == nil && b {
				if f.Z {
					return time.ParseInLocation(f.F, s, loc)
				} else {
					return time.Parse(f.F, s)
				}
			}
		}
	}
	return time.Time{}, errorf(errorParseTime, s)
}
