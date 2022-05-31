// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseStatus(t *testing.T) {
	tests := []struct {
		input  string
		output Status
	}{
		{"stop", Stop},
		{"run", Run},
		{"syncwait", SyncWait},
		{"unknown", Unknown},
	}
	for _, v := range tests {
		output := parseStatus(v.input)
		assert.Equal(t, output, v.output, fmt.Sprintf("%v != %v", output, v.output))
	}
}
