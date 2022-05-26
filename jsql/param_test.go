// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParam_AddParam(t *testing.T) {
	p1 := &Param{Params: make([]*Param, 0)}
	p2 := &Param{Params: make([]*Param, 0)}
	p1.AddParam(p2)
	input := p1.Params
	output := []*Param{p2}
	assert.Equal(t, input, output, fmt.Sprintf("%v != %v", input, output))
}
