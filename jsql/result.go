// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type Result interface {
	// Row returns query row data
	// if multiple row then return first row data
	Row() map[string]interface{}
	// Rows returns query rows data
	Rows() []map[string]interface{}
	// RowStart returns the integer by query page start row number
	// if not query page, the value default zero
	RowStart() int64
	// RowEnd returns the integer by query page end row number
	// if not query page, the value default zero
	RowEnd() int64
	// TotalRecord returns query page total record
	// the other query then returns rows length
	TotalRecord() int64
	// LastInsertId returns the integer generated by the database
	// in response to a command. Typically this will be from an
	// "auto increment" column when inserting a new row. Not all
	// databases support this feature, and the syntax of such
	// statements varies.
	LastInsertId() (int64, error)
	// RowsAffected returns the number of rows affected by an
	// update, insert, or delete. Not every database or database
	// driver may support this.
	RowsAffected() (int64, error)
}

type agentResult struct {
	rows         []map[string]interface{}
	rowStart     int64
	rowEnd       int64
	totalRecord  int64
	lastInsertId lastInsertId
	rowsAffected rowsAffected
}

func (result agentResult) Row() map[string]interface{} {
	if len(result.rows) > 0 {
		return result.rows[0]
	}
	return nil
}

func (result agentResult) Rows() []map[string]interface{} {
	return result.rows
}

func (result agentResult) RowStart() int64 {
	return result.rowStart
}

func (result agentResult) RowEnd() int64 {
	return result.rowEnd
}

func (result agentResult) TotalRecord() int64 {
	return result.totalRecord
}

func (result agentResult) LastInsertId() (int64, error) {
	return result.lastInsertId.id, result.lastInsertId.err
}

func (result agentResult) RowsAffected() (int64, error) {
	return result.rowsAffected.rows, result.rowsAffected.err
}
