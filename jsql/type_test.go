// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestType_DriverName(t *testing.T) {
	tests := []struct {
		in  Type
		out string
	}{
		{MySql, "mysql"},
		{MSSql, "sqlserver"},
		{Oracle, "godror"},
		{PostgreSql, "postgres"},
		{Unknown, ""},
	}
	for _, v := range tests {
		str := v.in.DriverName()
		assert.Equal(t, str, v.out, fmt.Sprintf("%v != %v", str, v.out))
	}
}

func TestType_Param(t *testing.T) {
	tests := []struct {
		in  Type
		out string
	}{
		{MySql, "?"},
		{MSSql, "@p1"},
		{Oracle, ":0"},
		{PostgreSql, "$1"},
		{Unknown, "?"},
	}
	for _, v := range tests {
		str := v.in.Param(0)
		assert.Equal(t, str, v.out, fmt.Sprintf("%v != %v", str, v.out))
	}
}

func TestParseDBType(t *testing.T) {
	tests := []struct {
		in  string
		out Type
	}{
		{"MySql", MySql},
		{"MSSql", MSSql},
		{"Oracle", Oracle},
		{"PostgreSql", PostgreSql},
		{"Unknown", Unknown},
	}
	for _, v := range tests {
		if vt, err := ParseDBType(v.in); err != nil {
			if v.in == "Unknown" {
				assert.Equal(t, vt, v.out, fmt.Sprintf("%v != %v", vt, v.out))
			} else {
				t.Error(err)
			}
		} else {
			assert.Equal(t, vt, v.out, fmt.Sprintf("%v != %v", vt, v.out))
		}
	}
}
