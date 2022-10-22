// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jconf

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jfile"
	"testing"
)

func TestNew(t *testing.T) {
	testErr := "TEST ERROR:"
	conf := New()
	inFormat := jfile.Json
	conf.SetFormat(inFormat)
	outFormat := conf.Format()
	assert.Equal(t, inFormat, outFormat, fmt.Sprintf("%v != %v", outFormat, inFormat))

	// test load
	if err := conf.Load(); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Load must be return error"))
	}
	conf.SetFileName("error.json")
	if err := conf.Load(); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Load must be return error"))
	}
	conf.SetRoot("../files/")
	conf.SetFileName("error.json")
	if err := conf.Load(); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Load must be return error"))
	}
	inFileName := "test-jconf.json"
	conf.SetFileName(inFileName)
	outFileName := conf.FileName()
	assert.Equal(t, inFileName, outFileName, fmt.Sprintf("%v != %v", outFileName, inFileName))
	conf.SetEnvFileName("error.json")
	if err := conf.Load(); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Load must be return error"))
	}
	inEnvFileName := ""
	conf.SetEnvFileName(inEnvFileName)
	outEnvFileName := conf.EnvFileName()
	assert.Equal(t, inEnvFileName, outEnvFileName, fmt.Sprintf("%v != %v", outEnvFileName, inEnvFileName))
	conf.SetEnvVal("error")
	if err := conf.Load(); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Load must be return error"))
	}
	inEnvVal := ""
	conf.SetEnvVal(inEnvVal)
	outEnvVal := conf.EnvVal()
	assert.Equal(t, inEnvVal, outEnvVal, fmt.Sprintf("%v != %v", outEnvVal, inEnvVal))
	conf.SetEnvKey("error")
	if err := conf.Load(); err != nil {
		t.Error(err)
	}
	conf.SetEnvVal("")
	conf.SetEnvKey("error-json")
	if err := conf.Load(); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Load must be return error"))
	}
	conf.SetEnvVal("")
	inEnvKey := ""
	conf.SetEnvKey(inEnvKey)
	outEnvKey := conf.EnvKey()
	assert.Equal(t, inEnvKey, outEnvKey, fmt.Sprintf("%v != %v", outEnvKey, inEnvKey))
	conf.DisableEnv()
	conf.EnableEnv()
	if err := conf.Load(); err != nil {
		t.Error(err)
		return
	}

	// test string
	if _, err := conf.String(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.String must be return error"))
	}
	if v, err := conf.String("map1", "map1-1", "str"); err != nil {
		t.Error(err)
	} else {
		output := "test-str"
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test bool
	if _, err := conf.Bool(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Bool must be return error"))
	}
	if v, err := conf.Bool("map1", "map1-1", "bool"); err != nil {
		t.Error(err)
	} else {
		msg := fmt.Sprintf("%v != %v", v, true)
		assert.Equal(t, true, v, msg)
	}

	// test int
	if _, err := conf.Int(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Int must be return error"))
	}
	if v, err := conf.Int("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := 7
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test int8
	if _, err := conf.Int8(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Int8 must be return error"))
	}
	if v, err := conf.Int8("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := int8(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test int16
	if _, err := conf.Int16(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Int16 must be return error"))
	}
	if v, err := conf.Int16("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := int16(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test int32
	if _, err := conf.Int32(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Int32 must be return error"))
	}
	if v, err := conf.Int32("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := int32(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test int64
	if _, err := conf.Int64(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Int64 must be return error"))
	}
	if v, err := conf.Int64("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := int64(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test uint
	if _, err := conf.Uint(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Uint must be return error"))
	}
	if v, err := conf.Uint("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := uint(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test uint8
	if _, err := conf.Uint8(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Uint8 must be return error"))
	}
	if v, err := conf.Uint8("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := uint8(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test uint16
	if _, err := conf.Uint16(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Uint16 must be return error"))
	}
	if v, err := conf.Uint16("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := uint16(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test uint32
	if _, err := conf.Uint32(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Uint32 must be return error"))
	}
	if v, err := conf.Uint32("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := uint32(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test int64
	if _, err := conf.Uint64(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Uint64 must be return error"))
	}
	if v, err := conf.Uint64("map1", "map1-1", "int"); err != nil {
		t.Error(err)
	} else {
		output := uint64(7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test float32
	if _, err := conf.Float32(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Float32 must be return error"))
	}
	if v, err := conf.Float32("map1", "map1-1", "float"); err != nil {
		t.Error(err)
	} else {
		output := float32(7.7)
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test float64
	if _, err := conf.Float64(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.Float64 must be return error"))
	}
	if v, err := conf.Float64("map1", "map1-1", "float"); err != nil {
		t.Error(err)
	} else {
		output := 7.7
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test map[interface{}]interface{}
	if _, err := conf.InterfaceMapInterface(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.InterfaceMapInterface must be return error"))
	}
	if v, err := conf.InterfaceMapInterface("map1", "map1-1"); err != nil {
		t.Error(err)
	} else {
		output := make(map[interface{}]interface{})
		output["str"] = "test-str"
		output["bool"] = true
		output["int"] = float64(7)
		output["float"] = 7.7
		output["test-bool"] = false
		output["test-int"] = float64(-7)
		output["test-float"] = -7.7
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test map[string]interface{}
	if _, err := conf.StringMapInterface(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.StringMapInterface must be return error"))
	}
	if v, err := conf.StringMapInterface("map1", "map1-1"); err != nil {
		t.Error(err)
	} else {
		output := make(map[string]interface{})
		output["str"] = "test-str"
		output["bool"] = true
		output["int"] = float64(7)
		output["float"] = 7.7
		output["test-bool"] = false
		output["test-int"] = float64(-7)
		output["test-float"] = -7.7
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []interface{}
	if _, err := conf.SliceInterface(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceInterface must be return error"))
	}
	if v, err := conf.SliceInterface("slice1"); err != nil {
		t.Error(err)
	} else {
		output := []interface{}{"str", true, float64(7), 7.7}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []string
	if _, err := conf.SliceString(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceString must be return error"))
	}
	if v, err := conf.SliceString("sliceString"); err != nil {
		t.Error(err)
	} else {
		output := []string{"str1", "str2"}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []bool
	if _, err := conf.SliceBool(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceBool must be return error"))
	}
	if v, err := conf.SliceBool("sliceBool"); err != nil {
		t.Error(err)
	} else {
		output := []bool{true, false}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []int
	if _, err := conf.SliceInt(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceInt must be return error"))
	}
	if v, err := conf.SliceInt("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []int{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []int8
	if _, err := conf.SliceInt8(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceInt8 must be return error"))
	}
	if v, err := conf.SliceInt8("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []int8{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []int16
	if _, err := conf.SliceInt16(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceInt16 must be return error"))
	}
	if v, err := conf.SliceInt16("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []int16{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []int32
	if _, err := conf.SliceInt32(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceInt32 must be return error"))
	}
	if v, err := conf.SliceInt32("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []int32{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []int64
	if _, err := conf.SliceInt64(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceInt64 must be return error"))
	}
	if v, err := conf.SliceInt64("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []int64{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []uint
	if _, err := conf.SliceUint(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceUint must be return error"))
	}
	if v, err := conf.SliceUint("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []uint{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []uint8
	if _, err := conf.SliceUint8(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceUint8 must be return error"))
	}
	if v, err := conf.SliceUint8("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []uint8{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []uint16
	if _, err := conf.SliceUint16(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceUint16 must be return error"))
	}
	if v, err := conf.SliceUint16("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []uint16{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []uint32
	if _, err := conf.SliceUint32(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceUint32 must be return error"))
	}
	if v, err := conf.SliceUint32("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []uint32{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []uint64
	if _, err := conf.SliceUint64(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceUint64 must be return error"))
	}
	if v, err := conf.SliceUint64("sliceInt"); err != nil {
		t.Error(err)
	} else {
		output := []uint64{7, 77}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []float32
	if _, err := conf.SliceFloat32(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceFloat32 must be return error"))
	}
	if v, err := conf.SliceFloat32("sliceFloat"); err != nil {
		t.Error(err)
	} else {
		output := []float32{7.7, -7.7}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test []float64
	if _, err := conf.SliceFloat64(nil); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.SliceFloat64 must be return error"))
	}
	if v, err := conf.SliceFloat64("sliceFloat"); err != nil {
		t.Error(err)
	} else {
		output := []float64{7.7, -7.7}
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}

	// test convert
	{
		type TestStruct struct {
			Str   string
			Bool  bool
			Int   int
			Float float64
			M     map[string]interface{}
			S     []interface{}
		}
		ts := TestStruct{}
		output := TestStruct{
			Str:   "str",
			Bool:  true,
			Int:   7,
			Float: 7.7,
			M:     make(map[string]interface{}),
			S:     []interface{}{"str", true, float64(7), 7.7},
		}
		output.M["str"] = "str"
		output.M["bool"] = true
		output.M["int"] = float64(7)
		output.M["float"] = 7.7
		if err := conf.Convert(&ts, nil); err == nil {
			t.Error(fmt.Sprint(testErr, " conf.Convert must be return error"))
		}
		var map3 map[string]interface{}
		if m, err := conf.StringMapInterface("map3"); err != nil {
			t.Error(err)
		} else {
			map3 = m
		}
		if map3 != nil {
			if err := conf.Convert(&ts, map3); err != nil {
				t.Error(err)
			} else {
				msg := fmt.Sprintf("%v != %v", ts, output)
				assert.Equal(t, output, ts, msg)
			}
		}
		if err := conf.Convert(&ts, "map3"); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", ts, output)
			assert.Equal(t, output, ts, msg)
		}
	}

	// test other
	if v, err := conf.String("map1", "slice1-1", 5, 0); err != nil {
		t.Error(err)
	} else {
		output := "test-slice-slice"
		msg := fmt.Sprintf("%v != %v", v, output)
		assert.Equal(t, output, v, msg)
	}
	if _, err := conf.String("map1", "slice1-1", "error"); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.String must be return error"))
	}
	if _, err := conf.String("map1", "slice1-1", 99); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.String must be return error"))
	}
	if _, err := conf.String("map1", "slice1-1", 0, "error"); err == nil {
		t.Error(fmt.Sprint(testErr, " conf.String must be return error"))
	}
	if _, err := conf.Get(); err != nil {
		t.Error(err)
	}
}
