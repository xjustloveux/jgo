// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncoding(t *testing.T) {
	testErr := "TEST ERROR:"
	RegisterCodec(Json.String(), jsonCodec{})
	if _, err := GetCodec(Json.String()); err != nil {
		t.Error(err)
		return
	}
	m := make(map[string]interface{})
	m["str"] = "str"
	m["bool"] = true
	m["int"] = 7
	m["float"] = 7.7
	m1 := make(map[string]interface{})
	m1["str"] = "str"
	m1["bool"] = true
	m1["int"] = 7
	m1["float"] = 7.7
	m["map"] = m1
	s1 := []interface{}{"str", true, 7, 7.7}
	m["slice"] = s1
	var err error
	var b []byte
	if _, err = Encode(Yaml.String(), m); err != nil {
		fmt.Println(testErr)
		fmt.Println(err)
	} else {
		t.Error(fmt.Sprint(testErr, " Encode must be return error"))
		return
	}
	if b, err = Encode(Json.String(), m); err != nil {
		t.Error(err)
		return
	}
	m2 := make(map[string]interface{})
	if err = Decode(Yaml.String(), b, m2); err != nil {
		fmt.Println(testErr)
		fmt.Println(err)
	} else {
		t.Error(fmt.Sprint(testErr, " Decode must be return error"))
		return
	}
	if err = Decode(Json.String(), b, m2); err != nil {
		t.Error(err)
		return
	} else {
		m["int"] = float64(7)
		m["map"].(map[string]interface{})["int"] = float64(7)
		m["slice"].([]interface{})[2] = float64(7)
		msg := fmt.Sprintf("%v != %v", m, m2)
		assert.Equal(t, m, m2, msg)
	}
}
