// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAgentResult_Row(t *testing.T) {
	row := make(map[string]interface{})
	tests := []struct {
		input  agentResult
		output map[string]interface{}
	}{
		{agentResult{rows: []map[string]interface{}{row}}, row},
		{agentResult{rows: []map[string]interface{}{}}, nil},
	}
	for _, test := range tests {
		r := test.input.Row()
		assert.Equal(t, r, test.output, fmt.Sprintf("%v != %v", r, test.output))
	}
}

func TestAgentResult_Rows(t *testing.T) {
	row := make(map[string]interface{})
	tests := []struct {
		input  agentResult
		output []map[string]interface{}
	}{
		{agentResult{rows: []map[string]interface{}{row}}, []map[string]interface{}{row}},
		{agentResult{rows: []map[string]interface{}{}}, []map[string]interface{}{}},
	}
	for _, test := range tests {
		r := test.input.Rows()
		assert.Equal(t, r, test.output, fmt.Sprintf("%v != %v", r, test.output))
	}
}

func TestAgentResult_RowStart(t *testing.T) {
	tests := []struct {
		input  agentResult
		output int64
	}{
		{agentResult{rowStart: 7}, 7},
		{agentResult{rowStart: 77}, 77},
	}
	for _, test := range tests {
		r := test.input.RowStart()
		assert.Equal(t, r, test.output, fmt.Sprintf("%v != %v", r, test.output))
	}
}

func TestAgentResult_RowEnd(t *testing.T) {
	tests := []struct {
		input  agentResult
		output int64
	}{
		{agentResult{rowEnd: 7}, 7},
		{agentResult{rowEnd: 77}, 77},
	}
	for _, test := range tests {
		r := test.input.RowEnd()
		assert.Equal(t, r, test.output, fmt.Sprintf("%v != %v", r, test.output))
	}
}

func TestAgentResult_TotalRecord(t *testing.T) {
	tests := []struct {
		input  agentResult
		output int64
	}{
		{agentResult{totalRecord: 7}, 7},
		{agentResult{totalRecord: 77}, 77},
	}
	for _, test := range tests {
		r := test.input.TotalRecord()
		assert.Equal(t, r, test.output, fmt.Sprintf("%v != %v", r, test.output))
	}
}

func TestAgentResult_LastInsertId(t *testing.T) {
	e := errorStr("TEST ERROR")
	tests := []struct {
		input  agentResult
		output lastInsertId
	}{
		{agentResult{lastInsertId: lastInsertId{7, nil}}, lastInsertId{7, nil}},
		{agentResult{lastInsertId: lastInsertId{7, e}}, lastInsertId{7, e}},
	}
	for _, test := range tests {
		id, err := test.input.LastInsertId()
		assert.Equal(t, id, test.output.id, fmt.Sprintf("%v != %v", id, test.output.id))
		assert.Equal(t, err, test.output.err, fmt.Sprintf("%v != %v", err, test.output.err))
	}
}

func TestAgentResult_RowsAffected(t *testing.T) {
	e := errorStr("TEST ERROR")
	tests := []struct {
		input  agentResult
		output rowsAffected
	}{
		{agentResult{rowsAffected: rowsAffected{7, nil}}, rowsAffected{7, nil}},
		{agentResult{rowsAffected: rowsAffected{7, e}}, rowsAffected{7, e}},
	}
	for _, test := range tests {
		rows, err := test.input.RowsAffected()
		assert.Equal(t, rows, test.output.rows, fmt.Sprintf("%v != %v", rows, test.output.rows))
		assert.Equal(t, err, test.output.err, fmt.Sprintf("%v != %v", err, test.output.err))
	}
}
