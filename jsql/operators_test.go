// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperators_String(t *testing.T) {
	tests := []struct {
		in  Operators
		out string
	}{
		{Equal, "Equal"},
		{NotEqual, "NotEqual"},
		{In, "In"},
		{NotIn, "NotIn"},
		{Between, "Between"},
		{NotBetween, "NotBetween"},
		{IsNull, "IsNull"},
		{IsNotNull, "IsNotNull"},
		{Like, "Like"},
		{SLike, "SLike"},
		{ELike, "ELike"},
		{Greater, "Greater"},
		{GreaterThanOrEqual, "GreaterThanOrEqual"},
		{Less, "Less"},
		{LessThanOrEqual, "LessThanOrEqual"},
		{Unknown, "Unknown"},
	}
	for _, v := range tests {
		str := v.in.String()
		assert.Equal(t, str, v.out, fmt.Sprintf("%v != %v", str, v.out))
	}
}

func TestParseOperators(t *testing.T) {
	testErr := "TEST ERROR:"
	tests := []struct {
		in  string
		out Operators
	}{
		{"Equal", Equal},
		{"NotEqual", NotEqual},
		{"In", In},
		{"NotIn", NotIn},
		{"Between", Between},
		{"NotBetween", NotBetween},
		{"IsNull", IsNull},
		{"IsNotNull", IsNotNull},
		{"Like", Like},
		{"SLike", SLike},
		{"ELike", ELike},
		{"Greater", Greater},
		{"GreaterThanOrEqual", GreaterThanOrEqual},
		{"Less", Less},
		{"LessThanOrEqual", LessThanOrEqual},
		{"Unknown", Unknown},
	}
	for _, v := range tests {
		if v.in == "Unknown" {
			if ck, err := ParseOperators(v.in); err == nil {
				t.Error(fmt.Sprint(testErr, " ParseOperators must be return error"))
			} else {
				assert.Equal(t, ck, v.out, fmt.Sprintf("%v != %v", ck, v.out))
			}
		} else {
			if ck, err := ParseOperators(v.in); err != nil {
				t.Error(err)
			} else {
				assert.Equal(t, ck, v.out, fmt.Sprintf("%v != %v", ck, v.out))
			}
		}
	}
}
