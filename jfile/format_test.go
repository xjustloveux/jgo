// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat_String(t *testing.T) {
	str := Json.String()
	assert.Equal(t, "Json", str, fmt.Sprintf("%v != %v", str, "Json"))
}
