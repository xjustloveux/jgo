// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jslice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFilter(t *testing.T) {
	iv := make(map[string]interface{})
	s1 := []interface{}{"str", true, iv, 7, 77, 7.7, []byte{55}}
	s2 := []interface{}{7, 77}
	tests := []struct {
		input  []interface{}
		output []interface{}
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := Filter(test.input, func(i interface{}) bool {
			if reflect.TypeOf(i).Kind() == reflect.Int {
				return true
			}
			return false
		}); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInsert(t *testing.T) {
	iv := make(map[string]interface{})
	s1 := []interface{}{"str", true, 7, 7.7, []byte{55}}
	s2 := []interface{}{"str", true, iv, 7, 7.7, []byte{55}}
	tests := []struct {
		input  []interface{}
		output []interface{}
	}{
		{s1, s2},
	}
	for _, test := range tests {
		if v, err := Insert(2, test.input, iv); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}

func TestInsertAll(t *testing.T) {
	s1 := []interface{}{"str", true, 7, 7.7, []byte{55}}
	s2 := []interface{}{77, "77"}
	s3 := []interface{}{"str", true, 77, "77", 7, 7.7, []byte{55}}
	tests := []struct {
		input  []interface{}
		output []interface{}
	}{
		{s1, s3},
	}
	for _, test := range tests {
		if v, err := InsertAll(2, test.input, s2); err != nil {
			t.Error(err)
		} else {
			msg := fmt.Sprintf("%v != %v", v, test.output)
			assert.Equal(t, test.output, v, msg)
		}
	}
}
