// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOperations_String(t *testing.T) {
	tests := []struct {
		in  Operations
		out string
	}{
		{Select, "Select"},
		{Insert, "Insert"},
		{Update, "Update"},
		{Delete, "Delete"},
		{Other, "Other"},
		{Unknown, "Unknown"},
	}
	for _, v := range tests {
		str := v.in.String()
		assert.Equal(t, str, v.out, fmt.Sprintf("%v != %v", str, v.out))
	}
}
