// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jruntime"
	"reflect"
)

type TableAgent struct {
	Agent  *Agent
	DSKey  string
	Table  string
	SelStr string
	OrdStr string
	Col    map[string]interface{}
	Params []*Param
}

// AddColumn add insert or update column data
func (ta *TableAgent) AddColumn(args ...interface{}) error {
	if ta.Col == nil {
		ta.Col = make(map[string]interface{})
	}
	var name string
	for i, c := range args {
		if i%2 == 0 {
			if v, ok := c.(string); ok {
				name = v
			} else {
				return errorFmt(errorColTypeNotStringType, reflect.TypeOf(c))
			}
		} else {
			ta.Col[name] = c
		}
	}
	return nil
}

// AddMap add insert or update column data with map[string]interface{}
func (ta *TableAgent) AddMap(m map[string]interface{}) {
	if ta.Col == nil {
		ta.Col = m
	} else {
		for k, v := range m {
			ta.Col[k] = v
		}
	}
}

// SetMap set insert or update column data with map[string]interface{}
func (ta *TableAgent) SetMap(m map[string]interface{}) {
	ta.Col = m
}

// AddParam add select or update or delete where clause param
func (ta *TableAgent) AddParam(param *Param) {
	if ta.Params == nil {
		ta.Params = make([]*Param, 0)
	}
	ta.Params = append(ta.Params, param)
}

// SetParams set select or update or delete where clause param
func (ta *TableAgent) SetParams(p []*Param) {
	ta.Params = p
}

// Equal add And Equal Param
func (ta *TableAgent) Equal(col string, val interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// NotEqual add And NotEqual Param
func (ta *TableAgent) NotEqual(col string, val interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// In add And In Param
func (ta *TableAgent) In(col string, val []interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// NotIn add And NotIn Param
func (ta *TableAgent) NotIn(col string, val []interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// Between add And Between Param
func (ta *TableAgent) Between(col string, val []interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// NotBetween add And NotBetween Param
func (ta *TableAgent) NotBetween(col string, val []interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// IsNull add And IsNull Param
func (ta *TableAgent) IsNull(col string) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: nil, Opr: opr})
}

// IsNotNull add And IsNotNull Param
func (ta *TableAgent) IsNotNull(col string) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: nil, Opr: opr})
}

// Like add And Like Param
func (ta *TableAgent) Like(col string, val string) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// SLike add And SLike Param
func (ta *TableAgent) SLike(col string, val string) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// ELike add And ELike Param
func (ta *TableAgent) ELike(col string, val string) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// Greater add And Greater Param
func (ta *TableAgent) Greater(col string, val interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// GreaterThanOrEqual add And GreaterThanOrEqual Param
func (ta *TableAgent) GreaterThanOrEqual(col string, val interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// Less add And Less Param
func (ta *TableAgent) Less(col string, val interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// LessThanOrEqual add And LessThanOrEqual Param
func (ta *TableAgent) LessThanOrEqual(col string, val interface{}) {
	opr, _ := ParseOperators(jruntime.GetFuncName())
	ta.AddParam(&Param{Col: col, Val: val, Opr: opr})
}

// Query executes a query that returns Result
func (ta *TableAgent) Query(v ...interface{}) (Result, error) {
	if query, args, err := ta.getQueryAndArgs(); err != nil {
		return nil, err
	} else {
		var r Result
		if r, err = ta.Agent.QueryWithSql(query, args...); err != nil || len(v) == 0 {
			return r, err
		} else {
			m := map[string]interface{}{"Rows": r.Rows()}
			err = jfile.Convert(m, v[0])
			return r, err
		}
	}
}

// QueryTx executes a query that returns Result
func (ta *TableAgent) QueryTx(v ...interface{}) (Result, error) {
	if ta.Agent == nil {
		return nil, errorStr(errorAgentNil)
	}
	if query, args, err := ta.getQueryAndArgs(); err != nil {
		return nil, err
	} else {
		var r Result
		if r, err = ta.Agent.QueryTxWithSql(query, args...); err != nil || len(v) == 0 {
			return r, err
		} else {
			m := map[string]interface{}{"Rows": r.Rows()}
			err = jfile.Convert(m, v[0])
			return r, err
		}
	}
}

// QueryRow executes a query that is expected to return at most one row
func (ta *TableAgent) QueryRow(v ...interface{}) (Result, error) {
	if query, args, err := ta.getQueryAndArgs(); err != nil {
		return nil, err
	} else {
		var r Result
		if r, err = ta.Agent.QueryRowWithSql(query, args...); err != nil || len(v) == 0 {
			return r, err
		} else {
			m := r.Row()
			if m == nil {
				m = make(map[string]interface{})
			}
			err = jfile.Convert(m, v[0])
			return r, err
		}
	}
}

// QueryRowTx executes a query that is expected to return at most one row
func (ta *TableAgent) QueryRowTx(v ...interface{}) (Result, error) {
	if ta.Agent == nil {
		return nil, errorStr(errorAgentNil)
	}
	if query, args, err := ta.getQueryAndArgs(); err != nil {
		return nil, err
	} else {
		var r Result
		if r, err = ta.Agent.QueryRowTxWithSql(query, args...); err != nil || len(v) == 0 {
			return r, err
		} else {
			m := r.Row()
			if m == nil {
				m = make(map[string]interface{})
			}
			err = jfile.Convert(m, v[0])
			return r, err
		}
	}
}

// QueryPage executes a query that returns Result
// the start and end are for query start row and end row
func (ta *TableAgent) QueryPage(start, end int64, v ...interface{}) (Result, error) {
	if query, args, err := ta.getQuery(); err != nil {
		return nil, err
	} else {
		var r Result
		if r, err = ta.Agent.QueryPageWithSql(query, ta.OrdStr, start, end, args...); err != nil || len(v) == 0 {
			return r, err
		} else {
			m := map[string]interface{}{
				"Rows":        r.Rows(),
				"RowStart":    r.RowStart(),
				"RowEnd":      r.RowEnd(),
				"TotalRecord": r.TotalRecord(),
			}
			err = jfile.Convert(m, v[0])
			return r, err
		}
	}
}

// QueryPageTx executes a query that returns Result
// the start and end are for query start row and end row
func (ta *TableAgent) QueryPageTx(start, end int64, v ...interface{}) (Result, error) {
	if ta.Agent == nil {
		return nil, errorStr(errorAgentNil)
	}
	if query, args, err := ta.getQuery(); err != nil {
		return nil, err
	} else {
		var r Result
		if r, err = ta.Agent.QueryPageTxWithSql(query, ta.OrdStr, start, end, args...); err != nil || len(v) == 0 {
			return r, err
		} else {
			m := map[string]interface{}{
				"Rows":        r.Rows(),
				"RowStart":    r.RowStart(),
				"RowEnd":      r.RowEnd(),
				"TotalRecord": r.TotalRecord(),
			}
			err = jfile.Convert(m, v[0])
			return r, err
		}
	}
}

// Count return query count
func (ta *TableAgent) Count() (int, error) {
	if query, args, err := ta.getCountAndArgs(); err != nil {
		return 0, err
	} else {
		var count int
		if err = ta.Agent.queryRowScan(query, &count, args...); err != nil {
			return 0, err
		}
		return count, nil
	}
}

// CountTx return query count
func (ta *TableAgent) CountTx() (int, error) {
	if query, args, err := ta.getCountAndArgs(); err != nil {
		return 0, err
	} else {
		var count int
		if err = ta.Agent.queryRowScanTx(query, &count, args...); err != nil {
			return 0, err
		}
		return count, nil
	}
}

// Exists return query sql exists data
func (ta *TableAgent) Exists() (bool, error) {
	if query, args, err := ta.getQueryAndArgs(); err != nil {
		return false, err
	} else {
		query = getExistsSql(ta.Agent.DBType(), query)
		var e string
		if err = ta.Agent.queryRowScan(query, &e, args...); err != nil {
			return false, err
		}
		return e == "Y", nil
	}
}

// ExistsTx return query sql exists data
func (ta *TableAgent) ExistsTx() (bool, error) {
	if query, args, err := ta.getQueryAndArgs(); err != nil {
		return false, err
	} else {
		query = getExistsSql(ta.Agent.DBType(), query)
		var e string
		if err = ta.Agent.queryRowScanTx(query, &e, args...); err != nil {
			return false, err
		}
		return e == "Y", nil
	}
}

// Insert executes a query with db.Exec
func (ta *TableAgent) Insert() (Result, error) {
	if query, args, err := ta.getInsert(); err != nil {
		return nil, err
	} else {
		return ta.Agent.exec(query, args...)
	}
}

// InsertTx executes a query with tx.Exec
func (ta *TableAgent) InsertTx() (Result, error) {
	if ta.Agent == nil {
		return nil, errorStr(errorAgentNil)
	}
	if query, args, err := ta.getInsert(); err != nil {
		return nil, err
	} else {
		return ta.Agent.execTx(query, args...)
	}
}

// InsertWithLastInsertId return last insert id by QueryRow.Scan
func (ta *TableAgent) InsertWithLastInsertId() (int, error) {
	if query, args, err := ta.getInsert(); err != nil {
		return 0, err
	} else {
		query = ta.getInsertWithLastInsertId(query)
		var id int
		if err = ta.Agent.queryRowScan(query, &id, args...); err != nil {
			return 0, err
		}
		return id, nil
	}
}

// InsertTxWithLastInsertId return last insert id by QueryRow.Scan
func (ta *TableAgent) InsertTxWithLastInsertId() (int, error) {
	if query, args, err := ta.getInsert(); err != nil {
		return 0, err
	} else {
		query = ta.getInsertWithLastInsertId(query)
		var id int
		if err = ta.Agent.queryRowScanTx(query, &id, args...); err != nil {
			return 0, err
		}
		return id, nil
	}
}

// Update executes a query with db.Exec
func (ta *TableAgent) Update() (Result, error) {
	if query, args, err := ta.getUpdate(); err != nil {
		return nil, err
	} else {
		return ta.Agent.exec(query, args...)
	}
}

// UpdateTx executes a query with tx.Exec
func (ta *TableAgent) UpdateTx() (Result, error) {
	if ta.Agent == nil {
		return nil, errorStr(errorAgentNil)
	}
	if query, args, err := ta.getUpdate(); err != nil {
		return nil, err
	} else {
		return ta.Agent.execTx(query, args...)
	}
}

// Delete executes a query with db.Exec
func (ta *TableAgent) Delete() (Result, error) {
	if query, args, err := ta.getDelete(); err != nil {
		return nil, err
	} else {
		return ta.Agent.exec(query, args...)
	}
}

// DeleteTx executes a query with tx.Exec
func (ta *TableAgent) DeleteTx() (Result, error) {
	if ta.Agent == nil {
		return nil, errorStr(errorAgentNil)
	}
	if query, args, err := ta.getDelete(); err != nil {
		return nil, err
	} else {
		return ta.Agent.execTx(query, args...)
	}
}

// Drop executes a query with db.Exec
func (ta *TableAgent) Drop() (Result, error) {
	if ta.Table == "" {
		return nil, errorStr(errorTableEmpty)
	}
	if ta.Agent == nil {
		var err error
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return nil, err
		}
	}
	return ta.Agent.exec(fmt.Sprint("DROP TABLE ", ta.Table))
}

// DropTx executes a query with tx.Exec
func (ta *TableAgent) DropTx() (Result, error) {
	if ta.Table == "" {
		return nil, errorStr(errorTableEmpty)
	}
	if ta.Agent == nil {
		var err error
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return nil, err
		}
	}
	return ta.Agent.execTx(fmt.Sprint("DROP TABLE ", ta.Table))
}

func (ta *TableAgent) getQueryAndArgs() (query string, args []interface{}, err error) {
	if query, args, err = ta.getQuery(); err != nil {
		return "", nil, err
	}
	if ta.OrdStr != "" {
		query = fmt.Sprint(query, " ORDER BY ", ta.OrdStr)
	}
	return query, args, nil
}

func (ta *TableAgent) getQuery() (query string, args []interface{}, err error) {
	if ta.Table == "" {
		return "", nil, errorStr(errorTableEmpty)
	}
	if ta.Agent == nil {
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return "", nil, err
		}
	}
	if ta.SelStr == "" {
		ta.SelStr = "*"
	}
	query = fmt.Sprint("SELECT ", ta.SelStr, " FROM ", ta.Table, " WHERE 1 = 1")
	args = make([]interface{}, 0)
	if ta.Params != nil {
		for _, param := range ta.Params {
			var clause string
			var pm []interface{}
			if clause, pm, err = param.getClauseAndParams(ta.Agent.DBType(), args); err != nil {
				return "", nil, err
			}
			query = fmt.Sprint(query, clause)
			args = pm
		}
	}
	return query, args, nil
}

func (ta *TableAgent) getCountAndArgs() (query string, args []interface{}, err error) {
	if ta.Table == "" {
		return "", nil, errorStr(errorTableEmpty)
	}
	if ta.Agent == nil {
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return "", nil, err
		}
	}
	if ta.SelStr == "" {
		ta.SelStr = "COUNT(*) AS NUM"
	}
	query = fmt.Sprint("SELECT ", ta.SelStr, " FROM ", ta.Table, " WHERE 1 = 1")
	args = make([]interface{}, 0)
	if ta.Params != nil {
		for _, param := range ta.Params {
			var clause string
			var pm []interface{}
			if clause, pm, err = param.getClauseAndParams(ta.Agent.DBType(), args); err != nil {
				return "", nil, err
			}
			query = fmt.Sprint(query, clause)
			args = pm
		}
	}
	return query, args, nil
}

func (ta *TableAgent) getInsert() (query string, args []interface{}, err error) {
	if ta.Table == "" {
		return "", nil, errorStr(errorTableEmpty)
	}
	if ta.Col == nil || len(ta.Col) <= 0 {
		return "", nil, errorStr(errorColNil)
	}
	if ta.Agent == nil {
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return "", nil, err
		}
	}
	query = fmt.Sprint("INSERT INTO ", ta.Table)
	col := "("
	val := "("
	args = make([]interface{}, len(ta.Col))
	count := 0
	for k, c := range ta.Col {
		if count == 0 {
			col = fmt.Sprint(col, k)
			val = fmt.Sprint(val, ta.Agent.t.Param(count))
		} else {
			col = fmt.Sprint(col, ", ", k)
			val = fmt.Sprint(val, ", ", ta.Agent.t.Param(count))
		}
		args[count] = c
		count++
	}
	col = fmt.Sprint(col, ")")
	val = fmt.Sprint(val, ")")
	query = fmt.Sprint(query, " ", col, " VALUES ", val)
	return query, args, nil
}

func (ta *TableAgent) getInsertWithLastInsertId(query string) string {
	switch ta.Agent.DBType() {
	case MSSql:
		return fmt.Sprint(query, "; SELECT SCOPE_IDENTITY()")
	case PostgreSql:
		return fmt.Sprint(query, " RETURNING id")
	}
	return query
}

func (ta *TableAgent) getUpdate() (query string, args []interface{}, err error) {
	if ta.Table == "" {
		return "", nil, errorStr(errorTableEmpty)
	}
	if ta.Col == nil || len(ta.Col) <= 0 {
		return "", nil, errorStr(errorColNil)
	}
	if ta.Agent == nil {
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return "", nil, err
		}
	}
	query = fmt.Sprint("UPDATE ", ta.Table, " SET")
	col := ""
	args = make([]interface{}, len(ta.Col))
	count := 0
	for k, c := range ta.Col {
		if count == 0 {
			col = fmt.Sprint(col, k, " = ", ta.Agent.t.Param(count))
		} else {
			col = fmt.Sprint(col, ", ", k, " = ", ta.Agent.t.Param(count))
		}
		args[count] = c
		count++
	}
	query = fmt.Sprint(query, " ", col, " WHERE 1 = 1")
	if ta.Params != nil {
		for _, param := range ta.Params {
			var clause string
			var pm []interface{}
			if clause, pm, err = param.getClauseAndParams(ta.Agent.DBType(), args); err != nil {
				return "", nil, err
			}
			query = fmt.Sprint(query, clause)
			args = pm
		}
	}
	return query, args, nil
}

func (ta *TableAgent) getDelete() (query string, args []interface{}, err error) {
	if ta.Table == "" {
		return "", nil, errorStr(errorTableEmpty)
	}
	if ta.Agent == nil {
		if ta.Agent, err = GetAgent(ta.DSKey); err != nil {
			return "", nil, err
		}
	}
	query = fmt.Sprint("DELETE FROM ", ta.Table, " WHERE 1 = 1")
	args = make([]interface{}, 0)
	if ta.Params != nil {
		for _, param := range ta.Params {
			var clause string
			var pm []interface{}
			if clause, pm, err = param.getClauseAndParams(ta.Agent.DBType(), args); err != nil {
				return "", nil, err
			}
			query = fmt.Sprint(query, clause)
			args = pm
		}
	}
	return query, args, nil
}
