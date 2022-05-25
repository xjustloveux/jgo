// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"database/sql"
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jtime"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

type Agent struct {
	db *sql.DB
	t  Type
	tx *sql.Tx
}

// DB returns this Agent *sql.DB
func (a *Agent) DB() *sql.DB {
	return a.db
}

// Tx returns this Agent *sql.Tx
func (a *Agent) Tx() *sql.Tx {
	return a.tx
}

// DBType returns this Agent db Type
func (a *Agent) DBType() Type {
	return a.t
}

// Ping same as sql.DB.Ping
// if db does not open, will call open before begin
func (a *Agent) Ping() error {
	if a.db == nil {
		return errors(errorDBNil)
	}
	return a.db.Ping()
}

// Begin same as sql.DB.Begin
func (a *Agent) Begin() (*sql.Tx, error) {
	if a.db == nil {
		return nil, errors(errorDBNil)
	}
	var err error
	if a.tx, err = a.db.Begin(); err != nil {
		a.tx = nil
		return nil, err
	}
	return a.tx, nil
}

// Commit same as sql.Tx.Commit
func (a *Agent) Commit() error {
	if a.db == nil {
		return errors(errorDBNil)
	}
	if a.tx == nil {
		return errors(errorDbNotBegin)
	}
	if err := a.tx.Commit(); err != nil {
		return err
	}
	a.tx = nil
	return nil
}

// Rollback same as sql.Tx.Rollback
func (a *Agent) Rollback() error {
	if a.db == nil {
		return errors(errorDBNil)
	}
	if a.tx == nil {
		return errors(errorDbNotBegin)
	}
	if err := a.tx.Rollback(); err != nil {
		return err
	}
	a.tx = nil
	return nil
}

// Query executes a query that returns Result
// the id are for xml select tag id
// the param are for any placeholder parameters in the query
func (a *Agent) Query(id string, param ...map[string]interface{}) (result Result, err error) {
	var query string
	var args []interface{}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	return a.QueryWithSql(query, args...)
}

// QueryTx executes a query that returns Result
// the id are for xml select tag id
// the param are for any placeholder parameters in the query
func (a *Agent) QueryTx(id string, param ...map[string]interface{}) (result Result, err error) {
	var query string
	var args []interface{}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	return a.QueryTxWithSql(query, args...)
}

// QueryWithSql executes a query that returns Result
func (a *Agent) QueryWithSql(query string, cond ...interface{}) (Result, error) {
	return a.query(false, query, cond...)
}

// QueryTxWithSql executes a query that returns Result
func (a *Agent) QueryTxWithSql(query string, cond ...interface{}) (Result, error) {
	return a.queryTx(false, query, cond...)
}

// QueryPrepare creates a prepared statement for later queries or executions
func (a *Agent) QueryPrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.queryPrepareMS(false, id, param, args...)
}

// QueryPrepareTx creates a prepared statement for later queries or executions
func (a *Agent) QueryPrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.queryPrepareTxMS(false, id, param, args...)
}

// QueryRow executes a query that is expected to return at most one row
// the id are for xml select tag id
// the param are for any placeholder parameters in the query
func (a *Agent) QueryRow(id string, param ...map[string]interface{}) (result Result, err error) {
	var query string
	var args []interface{}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	return a.QueryRowWithSql(query, args...)
}

// QueryRowTx executes a query that is expected to return at most one row
// the id are for xml select tag id
// the param are for any placeholder parameters in the query
func (a *Agent) QueryRowTx(id string, param ...map[string]interface{}) (result Result, err error) {
	var query string
	var args []interface{}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	return a.QueryRowTxWithSql(query, args...)
}

// QueryRowWithSql executes a query that is expected to return at most one row
func (a *Agent) QueryRowWithSql(query string, cond ...interface{}) (Result, error) {
	return a.query(true, query, cond...)
}

// QueryRowTxWithSql executes a query that is expected to return at most one row
func (a *Agent) QueryRowTxWithSql(query string, cond ...interface{}) (Result, error) {
	return a.queryTx(true, query, cond...)
}

// QueryRowPrepare creates a prepared statement for later queries or executions
func (a *Agent) QueryRowPrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.queryPrepareMS(true, id, param, args...)
}

// QueryRowPrepareTx creates a prepared statement for later queries or executions
func (a *Agent) QueryRowPrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.queryPrepareTxMS(true, id, param, args...)
}

// QueryPage executes a query that returns Result
// the id are for xml select tag id
// the start and end are for query start row and end row
// the param are for any placeholder parameters in the query
func (a *Agent) QueryPage(id string, start, end int64, param ...map[string]interface{}) (Result, error) {
	return a.queryPage(true, id, start, end, param...)
}

// QueryPageTx executes a query that returns Result
// the id are for xml select tag id
// the start and end are for query start row and end row
// the param are for any placeholder parameters in the query
func (a *Agent) QueryPageTx(id string, start, end int64, param ...map[string]interface{}) (Result, error) {
	return a.queryPage(false, id, start, end, param...)
}

// QueryPageWithSql executes a query that returns Result
// the start and end are for query start row and end row
func (a *Agent) QueryPageWithSql(query, order string, start, end int64, args ...interface{}) (Result, error) {
	pageQuery, countQuery := getPageSql(a.t, query, order, start, end)
	return a.queryPageWithSql(true, pageQuery, countQuery, start, end, args...)
}

// QueryPageTxWithSql executes a query that returns Result
// the start and end are for query start row and end row
func (a *Agent) QueryPageTxWithSql(query, order string, start, end int64, args ...interface{}) (Result, error) {
	pageQuery, countQuery := getPageSql(a.t, query, order, start, end)
	return a.queryPageWithSql(false, pageQuery, countQuery, start, end, args...)
}

// ExecWithSql executes a query with db.Exec
func (a *Agent) ExecWithSql(query string, cond ...interface{}) (Result, error) {
	return a.exec(query, cond...)
}

// ExecTxWithSql executes a query with tx.Exec
func (a *Agent) ExecTxWithSql(query string, cond ...interface{}) (Result, error) {
	return a.execTx(query, cond...)
}

// Insert executes a query with db.Exec
func (a *Agent) Insert(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execOps(Insert, id, param...)
}

// InsertTx executes a query with tx.Exec
func (a *Agent) InsertTx(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execTxOps(Insert, id, param...)
}

// InsertPrepare creates a prepared statement for later queries or executions
func (a *Agent) InsertPrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Insert, id, param, args...)
}

// InsertPrepareTx creates a prepared statement for later queries or executions
func (a *Agent) InsertPrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Insert, id, param, args...)
}

// Update executes a query with db.Exec
func (a *Agent) Update(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execOps(Update, id, param...)
}

// UpdateTx executes a query with tx.Exec
func (a *Agent) UpdateTx(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execTxOps(Update, id, param...)
}

// UpdatePrepare creates a prepared statement for later queries or executions
func (a *Agent) UpdatePrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Update, id, param, args...)
}

// UpdatePrepareTx creates a prepared statement for later queries or executions
func (a *Agent) UpdatePrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Update, id, param, args...)
}

// Delete executes a query with db.Exec
func (a *Agent) Delete(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execOps(Delete, id, param...)
}

// DeleteTx executes a query with tx.Exec
func (a *Agent) DeleteTx(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execTxOps(Delete, id, param...)
}

// DeletePrepare creates a prepared statement for later queries or executions
func (a *Agent) DeletePrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Delete, id, param, args...)
}

// DeletePrepareTx creates a prepared statement for later queries or executions
func (a *Agent) DeletePrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Delete, id, param, args...)
}

// Other executes a query with db.Exec
func (a *Agent) Other(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execOps(Other, id, param...)
}

// OtherTx executes a query with tx.Exec
func (a *Agent) OtherTx(id string, param ...map[string]interface{}) (result Result, err error) {
	return a.execTxOps(Other, id, param...)
}

// OtherPrepare creates a prepared statement for later queries or executions
func (a *Agent) OtherPrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Other, id, param, args...)
}

// OtherPrepareTx creates a prepared statement for later queries or executions
func (a *Agent) OtherPrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Other, id, param, args...)
}

func (a *Agent) queryPrepareMS(single bool, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	pm := make([]map[string]interface{}, 1)
	pm[0] = param
	if query, _, err = a.xmlAndParamsToQueryAndArgs(Select, id, pm); err != nil {
		return nil, err
	}
	return a.queryPrepare(single, query, args...)
}

func (a *Agent) queryPrepareTxMS(single bool, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	pm := make([]map[string]interface{}, 1)
	pm[0] = param
	if query, _, err = a.xmlAndParamsToQueryAndArgs(Select, id, pm); err != nil {
		return nil, err
	}
	return a.queryPrepareTx(single, query, args...)
}

func (a *Agent) execOps(ops Operations, id string, param ...map[string]interface{}) (result Result, err error) {
	var query string
	var args []interface{}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(ops, id, param); err != nil {
		return nil, err
	}
	return a.exec(query, args...)
}

func (a *Agent) execTxOps(ops Operations, id string, param ...map[string]interface{}) (result Result, err error) {
	var query string
	var args []interface{}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(ops, id, param); err != nil {
		return nil, err
	}
	return a.execTx(query, args...)
}

func (a *Agent) execPrepareOps(ops Operations, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	pm := make([]map[string]interface{}, 1)
	pm[0] = param
	if query, _, err = a.xmlAndParamsToQueryAndArgs(ops, id, pm); err != nil {
		return nil, err
	}
	return a.execPrepare(query, args...)
}

func (a *Agent) execPrepareTxOps(ops Operations, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	pm := make([]map[string]interface{}, 1)
	pm[0] = param
	if query, _, err = a.xmlAndParamsToQueryAndArgs(ops, id, pm); err != nil {
		return nil, err
	}
	return a.execPrepareTx(query, args...)
}

func (a *Agent) getXmlAndParam(ops Operations, id string, param []map[string]interface{}) (xml interface{}, pm map[string]interface{}, err error) {
	if xml, err = getXmlSql(ops, id); err != nil {
		return nil, nil, err
	}
	if len(param) > 0 {
		pm = param[0]
	} else {
		pm = nil
	}
	return xml, pm, nil
}

func (a *Agent) xmlAndParamsToQueryAndArgs(ops Operations, id string, param []map[string]interface{}) (query string, args []interface{}, err error) {
	var xml interface{}
	var xs *xmlSelect
	var xi *xmlInsert
	var xu *xmlUpdate
	var xd *xmlDelete
	var xo *xmlOther
	var ok bool
	var pm map[string]interface{}
	if xml, pm, err = a.getXmlAndParam(ops, id, param); err != nil {
		return "", nil, err
	}
	switch ops {
	case Select:
		if xs, ok = xml.(*xmlSelect); !ok {
			return "", nil, errorf(errorXmlNotSelectType, reflect.TypeOf(xml))
		}
		if query, _, err = xs.getSql(pm, false); err != nil {
			return "", nil, err
		}
	case Insert:
		if xi, ok = xml.(*xmlInsert); !ok {
			return "", nil, errorf(errorXmlNotInsertType, reflect.TypeOf(xml))
		}
		if query, _, err = xi.getSql(pm); err != nil {
			return "", nil, err
		}
	case Update:
		if xu, ok = xml.(*xmlUpdate); !ok {
			return "", nil, errorf(errorXmlNotUpdateType, reflect.TypeOf(xml))
		}
		if query, _, err = xu.getSql(pm); err != nil {
			return "", nil, err
		}
	case Delete:
		if xd, ok = xml.(*xmlDelete); !ok {
			return "", nil, errorf(errorXmlNotDeleteType, reflect.TypeOf(xml))
		}
		if query, _, err = xd.getSql(pm); err != nil {
			return "", nil, err
		}
	case Other:
		if xo, ok = xml.(*xmlOther); !ok {
			return "", nil, errorf(errorXmlNotOtherType, reflect.TypeOf(xml))
		}
		if query, _, err = xo.getSql(pm); err != nil {
			return "", nil, err
		}
	}
	query, args = a.getQueryAndArgs(query, pm)
	return trim(query), args, nil
}

func (a *Agent) getQueryAndArgs(sorQuery string, params map[string]interface{}) (query string, args []interface{}) {
	query = ""
	args = make([]interface{}, 0)
	pattern := regexp.MustCompile(`@\{\w+\}`)
	queryIdx := 0
	for i, v := range pattern.FindAllStringSubmatchIndex(sorQuery, -1) {
		st := "@{"
		et := "}"
		k := sorQuery[v[0]+len(st) : v[1]-len(et)]
		query = fmt.Sprint(query, sorQuery[queryIdx:v[0]], a.t.Param(i))
		queryIdx = v[1]
		if params != nil {
			args = append(args, params[k])
		}
	}
	query = fmt.Sprint(query, sorQuery[queryIdx:])
	return query, args
}

func (a *Agent) query(single bool, query string, args ...interface{}) (result Result, err error) {
	if a.db == nil {
		return nil, errors(errorDBNil)
	}
	var rows *sql.Rows
	if rows, err = a.db.Query(query, args...); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	return a.getResult(rows, single)
}

func (a *Agent) queryTx(single bool, query string, args ...interface{}) (result Result, err error) {
	if a.tx == nil {
		return nil, errors(errorDbNotBegin)
	}
	var rows *sql.Rows
	if rows, err = a.tx.Query(query, args...); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	return a.getResult(rows, single)
}

func (a *Agent) queryPage(ct bool, id string, start, end int64, param ...map[string]interface{}) (result Result, err error) {
	var xml interface{}
	var pm map[string]interface{}
	if xml, pm, err = a.getXmlAndParam(Select, id, param); err != nil {
		return nil, err
	}
	var xs *xmlSelect
	var ok bool
	if xs, ok = xml.(*xmlSelect); !ok {
		return nil, errorf(errorXmlNotSelectType, reflect.TypeOf(xml))
	}
	var query string
	var order string
	if query, order, err = xs.getSql(pm, true); err != nil {
		return nil, err
	}
	var args []interface{}
	pageQuery, countQuery := getPageSql(a.t, query, order, start, end)
	pageQuery, args = a.getQueryAndArgs(pageQuery, pm)
	countQuery, _ = a.getQueryAndArgs(countQuery, pm)
	return a.queryPageWithSql(ct, pageQuery, countQuery, start, end, args...)
}

func (a *Agent) queryPageWithSql(ct bool, pageQuery, countQuery string, start, end int64, args ...interface{}) (result Result, err error) {
	if ct {
		if _, err = a.Begin(); err != nil {
			return nil, err
		}
		defer func() {
			if err != nil {
				if e := a.Rollback(); e != nil {
					err = e
				}
			}
		}()
	}
	var resPage Result
	var resCount Result
	if resPage, err = a.queryTx(false, pageQuery, args...); err != nil {
		fmtPrintln(pageQuery)
		return nil, err
	}
	if res, ok := resPage.(agentResult); ok {
		for i := range res.rows {
			delete(res.rows[i], allowPagingId)
			delete(res.rows[i], orderById)
		}
	}
	if resCount, err = a.queryTx(true, countQuery, args...); err != nil {
		fmtPrintln(countQuery)
		return nil, err
	}
	if ct {
		if err = a.Commit(); err != nil {
			return nil, err
		}
	}
	if len(resCount.Rows()) <= 0 {
		return nil, errors(errorNoRowsAvailable)
	}
	var m map[string]interface{}
	m = resCount.Rows()[0]
	var total int64
	if total, err = jcast.Int64(m[totalRecord]); err != nil {
		return nil, err
	}
	return agentResult{
		rows:         resPage.Rows(),
		rowStart:     start,
		rowEnd:       end,
		totalRecord:  total,
		lastInsertId: lastInsertId{id: -1, err: nil},
		rowsAffected: rowsAffected{rows: 0, err: nil}}, nil
}

func (a *Agent) exec(query string, args ...interface{}) (result Result, err error) {
	if a.db == nil {
		return nil, errors(errorDBNil)
	}
	var res sql.Result
	if res, err = a.db.Exec(query, args...); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	id := lastInsertId{id: -1, err: nil}
	ra := rowsAffected{rows: 0, err: nil}
	id.id, id.err = res.LastInsertId()
	ra.rows, ra.err = res.RowsAffected()
	return agentResult{
		rows:         nil,
		rowStart:     0,
		rowEnd:       0,
		totalRecord:  0,
		lastInsertId: id,
		rowsAffected: ra}, nil
}

func (a *Agent) execTx(query string, args ...interface{}) (result Result, err error) {
	if a.tx == nil {
		return nil, errors(errorDbNotBegin)
	}
	var res sql.Result
	if res, err = a.tx.Exec(query, args...); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	id := lastInsertId{id: -1, err: nil}
	ra := rowsAffected{rows: 0, err: nil}
	id.id, id.err = res.LastInsertId()
	ra.rows, ra.err = res.RowsAffected()
	return agentResult{
		rows:         nil,
		rowStart:     0,
		rowEnd:       0,
		totalRecord:  0,
		lastInsertId: id,
		rowsAffected: ra}, nil
}

func (a *Agent) queryPrepare(single bool, query string, args ...[]interface{}) (result []Result, err error) {
	if a.db == nil {
		return nil, errors(errorDBNil)
	}
	var stmt *sql.Stmt
	if stmt, err = a.db.Prepare(query); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	return a.stmtQuery(single, stmt, args...)
}

func (a *Agent) queryPrepareTx(single bool, query string, args ...[]interface{}) (result []Result, err error) {
	if a.tx == nil {
		return nil, errors(errorDbNotBegin)
	}
	var stmt *sql.Stmt
	if stmt, err = a.tx.Prepare(query); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	return a.stmtQuery(single, stmt, args...)
}

func (a *Agent) stmtQuery(single bool, stmt *sql.Stmt, args ...[]interface{}) (result []Result, err error) {
	defer func() {
		if e := stmt.Close(); e != nil {
			err = e
		}
	}()
	result = make([]Result, len(args))
	for i, arg := range args {
		var rows *sql.Rows
		if rows, err = stmt.Query(arg...); err != nil {
			return nil, err
		}
		var res Result
		if res, err = a.getResult(rows, single); err != nil {
			return nil, err
		}
		result[i] = res
	}
	return result, nil
}

func (a *Agent) execPrepare(query string, args ...[]interface{}) (result []Result, err error) {
	if a.db == nil {
		return nil, errors(errorDBNil)
	}
	var stmt *sql.Stmt
	if stmt, err = a.db.Prepare(query); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	return a.stmtExec(stmt, args...)
}

func (a *Agent) execPrepareTx(query string, args ...[]interface{}) (result []Result, err error) {
	if a.tx == nil {
		return nil, errors(errorDbNotBegin)
	}
	var stmt *sql.Stmt
	if stmt, err = a.tx.Prepare(query); err != nil {
		fmtPrintln(query)
		return nil, err
	}
	return a.stmtExec(stmt, args...)
}

func (a *Agent) stmtExec(stmt *sql.Stmt, args ...[]interface{}) (result []Result, err error) {
	defer func() {
		if e := stmt.Close(); e != nil {
			err = e
		}
	}()
	result = make([]Result, len(args))
	for i, arg := range args {
		var res sql.Result
		if res, err = stmt.Exec(arg...); err != nil {
			return nil, err
		}
		id := lastInsertId{id: -1, err: nil}
		ra := rowsAffected{rows: 0, err: nil}
		id.id, id.err = res.LastInsertId()
		ra.rows, ra.err = res.RowsAffected()
		result[i] = agentResult{
			rows:         nil,
			rowStart:     0,
			rowEnd:       0,
			totalRecord:  0,
			lastInsertId: id,
			rowsAffected: ra}
	}
	return result, nil
}

func (a *Agent) getResult(rows *sql.Rows, single bool) (result Result, err error) {
	if rows == nil {
		err = errors(errorRowsNil)
		return nil, err
	}
	defer func() {
		if e := rows.Close(); e != nil {
			err = e
		}
	}()
	var colTypes []*sql.ColumnType
	if colTypes, err = rows.ColumnTypes(); err != nil {
		return nil, err
	}
	r := make([]map[string]interface{}, 0)
	for rows.Next() {
		rowValue := make([]interface{}, len(colTypes))
		rowParam := make([]interface{}, len(colTypes))
		for i, colType := range colTypes {
			rowValue[i] = reflect.New(colType.ScanType())
			rowParam[i] = reflect.ValueOf(&rowValue[i]).Interface()
		}
		if err = rows.Scan(rowParam...); err != nil {
			return nil, err
		}
		r = append(r, a.getRecord(colTypes, rowValue))
		if single {
			break
		}
	}
	return agentResult{
		rows:         r,
		rowStart:     0,
		rowEnd:       0,
		totalRecord:  int64(len(r)),
		lastInsertId: lastInsertId{id: -1, err: nil},
		rowsAffected: rowsAffected{rows: 0, err: nil}}, nil
}

func (a *Agent) getRecord(colTypes []*sql.ColumnType, rowValue []interface{}) map[string]interface{} {
	record := make(map[string]interface{})
	for i, colType := range colTypes {
		if rowValue[i] != nil {
			dbType := colType.DatabaseTypeName()
			switch colType.ScanType().String() {
			case "time.Time":
				switch dbType {
				case "DATETIME":
					record[colType.Name()] = rowValue[i].(time.Time).Format(jtime.DateTime)
				case "DATE":
					record[colType.Name()] = rowValue[i].(time.Time).Format(jtime.Date)
				default:
					record[colType.Name()] = rowValue[i].(time.Time).Format(jtime.DateTime)
				}
			case "sql.NullTime":
				switch dbType {
				case "DATETIME":
					record[colType.Name()] = rowValue[i].(time.Time).Format(jtime.DateTime)
				case "DATE":
					record[colType.Name()] = rowValue[i].(time.Time).Format(jtime.Date)
				default:
					record[colType.Name()] = rowValue[i].(time.Time).Format(jtime.DateTime)
				}
			case "int64":
				fallthrough
			case "sql.NullInt64":
				switch rowValue[i].(type) {
				case []byte:
					record[colType.Name()], _ = strconv.ParseInt(string(rowValue[i].([]byte)), 10, 64)
				default:
					record[colType.Name()] = rowValue[i]
				}
			case "sql.RawBytes":
				switch dbType {
				case "DOUBLE":
					record[colType.Name()], _ = strconv.ParseFloat(string(rowValue[i].([]byte)), 64)
				case "DECIMAL":
					record[colType.Name()], _ = strconv.ParseFloat(string(rowValue[i].([]byte)), 64)
				case "NUMERIC":
					record[colType.Name()], _ = strconv.ParseFloat(string(rowValue[i].([]byte)), 64)
				default:
					record[colType.Name()] = string(rowValue[i].([]byte))
				}
			default:
				record[colType.Name()] = rowValue[i]
			}
		} else {
			record[colType.Name()] = nil
		}
	}
	return record
}
