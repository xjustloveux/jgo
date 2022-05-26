// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseTag(t *testing.T) {
	tests := []struct {
		input  string
		output tag
	}{
		{"dao", tagDao},
		{"text", tagText},
		{"select", tagSelect},
		{"insert", tagInsert},
		{"update", tagUpdate},
		{"delete", tagDelete},
		{"other", tagOther},
		{"if", tagIf},
		{"foreach", tagForeach},
		{"orderby", tagOrderBy},
		{"unknown", tagUnknown},
	}
	for _, v := range tests {
		output := ParseTag(v.input)
		assert.Equal(t, output, v.output, fmt.Sprintf("%v != %v", output, v.output))
	}
}
