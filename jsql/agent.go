// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Agent struct {
	db     *sql.DB
	t      Type
	tx     *sql.Tx
	dbName string
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

// DbName returns this Agent db name
func (a *Agent) DbName() string {
	return a.dbName
}

// Ping same as sql.DB.Ping
// if db does not open, will call open before begin
func (a *Agent) Ping() error {
	if a.db == nil {
		return errorStr(errorDBNil)
	}
	return a.db.Ping()
}

// Begin same as sql.DB.Begin
func (a *Agent) Begin() (*sql.Tx, error) {
	if a.db == nil {
		return nil, errorStr(errorDBNil)
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
		return errorStr(errorDBNil)
	}
	if a.tx == nil {
		return errorStr(errorDbNotBegin)
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
		return errorStr(errorDBNil)
	}
	if a.tx == nil {
		return errorStr(errorDbNotBegin)
	}
	if err := a.tx.Rollback(); err != nil {
		return err
	}
	a.tx = nil
	return nil
}

// UseTx start a transaction as a block, return error will roll back, otherwise to commit
func (a *Agent) UseTx(f func() error) error {
	if _, err := a.Begin(); err != nil {
		return err
	} else {
		defer func() {
			if err != nil {
				if e := a.Rollback(); e != nil {
					fmt.Println(e)
				}
			}
		}()
		if err = f(); err != nil {
			return err
		}
		if err = a.Commit(); err != nil {
			return err
		}
	}
	return nil
}

// Query executes a query that returns Result
// the id are for xml select tag id
// the args are for any placeholder parameters in the query, or result struct point
func (a *Agent) Query(id string, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	var v interface{}
	var query string
	if param, v, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	if result, err = a.QueryWithSql(query, args...); err != nil {
		return result, err
	}
	if v != nil {
		m := map[string]interface{}{"Rows": result.Rows()}
		if err = jfile.Convert(m, v); err != nil {
			return result, err
		}
	}
	return result, nil
}

// QueryTx executes a query that returns Result
// the id are for xml select tag id
// the args are for any placeholder parameters in the query, or result struct point
func (a *Agent) QueryTx(id string, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	var v interface{}
	var query string
	if param, v, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	if result, err = a.QueryTxWithSql(query, args...); err != nil {
		return result, err
	}
	if v != nil {
		m := map[string]interface{}{"Rows": result.Rows()}
		if err = jfile.Convert(m, v); err != nil {
			return result, err
		}
	}
	return result, nil
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
// the args are for any placeholder parameters in the query, or result struct point
func (a *Agent) QueryRow(id string, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	var v interface{}
	var query string
	if param, v, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	if result, err = a.QueryRowWithSql(query, args...); err != nil {
		return result, err
	}
	if v != nil {
		m := result.Row()
		if m == nil {
			m = make(map[string]interface{})
		}
		if err = jfile.Convert(m, v); err != nil {
			return result, err
		}
	}
	return result, nil
}

// QueryRowTx executes a query that is expected to return at most one row
// the id are for xml select tag id
// the args are for any placeholder parameters in the query, or result struct point
func (a *Agent) QueryRowTx(id string, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	var v interface{}
	var query string
	if param, v, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	if query, args, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	if result, err = a.QueryRowTxWithSql(query, args...); err != nil {
		return result, err
	}
	if v != nil {
		m := result.Row()
		if m == nil {
			m = make(map[string]interface{})
		}
		if err = jfile.Convert(m, v); err != nil {
			return result, err
		}
	}
	return result, nil
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
// the args are for any placeholder parameters in the query, or result struct point
func (a *Agent) QueryPage(id string, start, end int64, args ...interface{}) (Result, error) {
	return a.queryPage(true, id, start, end, args...)
}

// QueryPageTx executes a query that returns Result
// the id are for xml select tag id
// the start and end are for query start row and end row
// the args are for any placeholder parameters in the query, or result struct point
func (a *Agent) QueryPageTx(id string, start, end int64, args ...interface{}) (Result, error) {
	return a.queryPage(false, id, start, end, args...)
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

// QuerySqlAndArgs returns query sql and args
func (a *Agent) QuerySqlAndArgs(id string, args ...interface{}) (string, []interface{}, error) {
	return a.getSqlAndArgs(Select, id, args...)
}

// Count return query count
func (a *Agent) Count(id string, args ...interface{}) (count int, err error) {
	var query string
	if query, args, err = a.QuerySqlAndArgs(id, args...); err != nil {
		return 0, err
	}
	if err = a.queryRowScan(query, &count, args...); err != nil {
		return 0, err
	}
	return
}

// CountTx return query count
func (a *Agent) CountTx(id string, args ...interface{}) (count int, err error) {
	var query string
	if query, args, err = a.QuerySqlAndArgs(id, args...); err != nil {
		return 0, err
	}
	if err = a.queryRowScanTx(query, &count, args...); err != nil {
		return 0, err
	}
	return
}

// Exists return query sql exists data
func (a *Agent) Exists(id string, args ...interface{}) (exists bool, err error) {
	var query string
	if query, args, err = a.QuerySqlAndArgs(id, args...); err != nil {
		return false, err
	}
	var e string
	query = getExistsSql(a.t, query)
	if err = a.queryRowScan(query, &e, args...); err != nil {
		return false, err
	}
	exists = e == "Y"
	return
}

// ExistsTx return query sql exists data
func (a *Agent) ExistsTx(id string, args ...interface{}) (exists bool, err error) {
	var query string
	if query, args, err = a.QuerySqlAndArgs(id, args...); err != nil {
		return false, err
	}
	var e string
	query = getExistsSql(a.t, query)
	if err = a.queryRowScanTx(query, &e, args...); err != nil {
		return false, err
	}
	exists = e == "Y"
	return
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
func (a *Agent) Insert(id string, args ...interface{}) (result Result, err error) {
	return a.execOps(Insert, id, args...)
}

// InsertTx executes a query with tx.Exec
func (a *Agent) InsertTx(id string, args ...interface{}) (result Result, err error) {
	return a.execTxOps(Insert, id, args...)
}

// InsertPrepare creates a prepared statement for later queries or executions
func (a *Agent) InsertPrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Insert, id, param, args...)
}

// InsertPrepareTx creates a prepared statement for later queries or executions
func (a *Agent) InsertPrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Insert, id, param, args...)
}

// InsertSqlAndArgs returns insert sql and args
func (a *Agent) InsertSqlAndArgs(id string, args ...interface{}) (string, []interface{}, error) {
	return a.getSqlAndArgs(Insert, id, args...)
}

// InsertWithLastInsertId return last insert id by QueryRow.Scan
func (a *Agent) InsertWithLastInsertId(id string, args ...interface{}) (lastInsertId int, err error) {
	var query string
	if query, args, err = a.InsertSqlAndArgs(id, args...); err != nil {
		return 0, err
	}
	if err = a.queryRowScan(query, &lastInsertId, args...); err != nil {
		return 0, err
	}
	return
}

// InsertTxWithLastInsertId return last insert id by QueryRow.Scan
func (a *Agent) InsertTxWithLastInsertId(id string, args ...interface{}) (lastInsertId int, err error) {
	var query string
	if query, args, err = a.InsertSqlAndArgs(id, args...); err != nil {
		return 0, err
	}
	if err = a.queryRowScanTx(query, &lastInsertId, args...); err != nil {
		return 0, err
	}
	return
}

// Update executes a query with db.Exec
func (a *Agent) Update(id string, args ...interface{}) (result Result, err error) {
	return a.execOps(Update, id, args...)
}

// UpdateTx executes a query with tx.Exec
func (a *Agent) UpdateTx(id string, args ...interface{}) (result Result, err error) {
	return a.execTxOps(Update, id, args...)
}

// UpdatePrepare creates a prepared statement for later queries or executions
func (a *Agent) UpdatePrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Update, id, param, args...)
}

// UpdatePrepareTx creates a prepared statement for later queries or executions
func (a *Agent) UpdatePrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Update, id, param, args...)
}

// UpdateSqlAndArgs returns update sql and args
func (a *Agent) UpdateSqlAndArgs(id string, args ...interface{}) (string, []interface{}, error) {
	return a.getSqlAndArgs(Update, id, args...)
}

// Delete executes a query with db.Exec
func (a *Agent) Delete(id string, args ...interface{}) (result Result, err error) {
	return a.execOps(Delete, id, args...)
}

// DeleteTx executes a query with tx.Exec
func (a *Agent) DeleteTx(id string, args ...interface{}) (result Result, err error) {
	return a.execTxOps(Delete, id, args...)
}

// DeletePrepare creates a prepared statement for later queries or executions
func (a *Agent) DeletePrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Delete, id, param, args...)
}

// DeletePrepareTx creates a prepared statement for later queries or executions
func (a *Agent) DeletePrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Delete, id, param, args...)
}

// DeleteSqlAndArgs returns delete sql and args
func (a *Agent) DeleteSqlAndArgs(id string, args ...interface{}) (string, []interface{}, error) {
	return a.getSqlAndArgs(Delete, id, args...)
}

// Other executes a query with db.Exec
func (a *Agent) Other(id string, args ...interface{}) (result Result, err error) {
	return a.execOps(Other, id, args...)
}

// OtherTx executes a query with tx.Exec
func (a *Agent) OtherTx(id string, args ...interface{}) (result Result, err error) {
	return a.execTxOps(Other, id, args...)
}

// OtherPrepare creates a prepared statement for later queries or executions
func (a *Agent) OtherPrepare(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareOps(Other, id, param, args...)
}

// OtherPrepareTx creates a prepared statement for later queries or executions
func (a *Agent) OtherPrepareTx(id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	return a.execPrepareTxOps(Other, id, param, args...)
}

// OtherSqlAndArgs returns other sql and args
func (a *Agent) OtherSqlAndArgs(id string, args ...interface{}) (string, []interface{}, error) {
	return a.getSqlAndArgs(Other, id, args...)
}

// Tables returns table name list
// or you can use args input query statement
func (a *Agent) Tables(args ...interface{}) ([]string, error) {
	query := ""
	param := make([]interface{}, 0)
	if len(args) > 0 {
		query = jcast.String(args[0])
		param = args[1:]
	} else {
		switch a.t {
		case MySql:
			query = sqlQueryTablesMySql
			param = append(param, a.DbName())
		case MSSql:
			query = sqlQueryTablesMSSql
		case Oracle:
			query = sqlQueryTablesOracle
		case PostgreSql:
			query = sqlQueryTablesPostgreSql
		default:
			return nil, errorStr(errorUnknownSqlTypeForAgentTables)
		}
	}
	if result, err := a.QueryWithSql(query, param...); err != nil {
		return nil, err
	} else {
		list := make([]string, len(result.Rows()))
		for i, v := range result.Rows() {
			var m map[string]string
			if m, err = jcast.StringMapString(v); err != nil {
				return nil, err
			}
			if m["TABLE_NAME"] != "" {
				list[i] = m["TABLE_NAME"]
			} else {
				list[i] = m["table_name"]
			}
		}
		return list, nil
	}
}

// TablesTx returns table name list
// or you can use args input query statement
func (a *Agent) TablesTx(args ...interface{}) ([]string, error) {
	query := ""
	param := make([]interface{}, 0)
	if len(args) > 0 {
		query = jcast.String(args[0])
		param = args[1:]
	} else {
		switch a.t {
		case MySql:
			query = sqlQueryTablesMySql
			param = append(param, a.DbName())
		case MSSql:
			query = sqlQueryTablesMSSql
		case Oracle:
			query = sqlQueryTablesOracle
		case PostgreSql:
			query = sqlQueryTablesPostgreSql
		default:
			return nil, errorStr(errorUnknownSqlTypeForAgentTables)
		}
	}
	if result, err := a.QueryTxWithSql(query, param...); err != nil {
		return nil, err
	} else {
		list := make([]string, len(result.Rows()))
		for i, v := range result.Rows() {
			var m map[string]string
			if m, err = jcast.StringMapString(v); err != nil {
				return nil, err
			}
			if m["TABLE_NAME"] != "" {
				list[i] = m["TABLE_NAME"]
			} else {
				list[i] = m["table_name"]
			}
		}
		return list, nil
	}
}

// TableSchema return table schema
// or you can use args input query statement
func (a *Agent) TableSchema(table string, args ...interface{}) ([]TableSchema, error) {
	query := ""
	param := make([]interface{}, 0)
	if len(args) > 0 {
		query = jcast.String(args[0])
		param = args[1:]
	} else {
		switch a.t {
		case MySql:
			query = sqlQueryTableSchemaMySql
			param = append(param, a.DbName(), table)
		case MSSql:
			query = sqlQueryTableSchemaMSSql
			param = append(param, table)
		case Oracle:
			query = sqlQueryTableSchemaOracle
			param = append(param, a.DbName(), table)
		case PostgreSql:
			query = sqlQueryTableSchemaPostgreSql
			param = append(param, table)
		default:
			return nil, errorStr(errorUnknownSqlTypeForAgentTables)
		}
	}
	if result, err := a.QueryWithSql(query, param...); err != nil {
		return nil, err
	} else {
		list := make([]TableSchema, len(result.Rows()))
		for i, v := range result.Rows() {
			if err = jfile.Convert(v, &list[i]); err != nil {
				return nil, err
			}
		}
		return list, nil
	}
}

// TableSchemaTx return table schema
// or you can use args input query statement
func (a *Agent) TableSchemaTx(table string, args ...interface{}) ([]TableSchema, error) {
	query := ""
	param := make([]interface{}, 0)
	if len(args) > 0 {
		query = jcast.String(args[0])
		param = args[1:]
	} else {
		switch a.t {
		case MySql:
			query = sqlQueryTableSchemaMySql
			param = append(param, a.DbName(), table)
		case MSSql:
			query = sqlQueryTableSchemaMSSql
			param = append(param, table)
		case Oracle:
			query = sqlQueryTableSchemaOracle
			param = append(param, a.DbName(), table)
		case PostgreSql:
			query = sqlQueryTableSchemaPostgreSql
			param = append(param, table)
		default:
			return nil, errorStr(errorUnknownSqlTypeForAgentTables)
		}
	}
	if result, err := a.QueryTxWithSql(query, param...); err != nil {
		return nil, err
	} else {
		list := make([]TableSchema, len(result.Rows()))
		for i, v := range result.Rows() {
			if err = jfile.Convert(v, &list[i]); err != nil {
				return nil, err
			}
		}
		return list, nil
	}
}

func (a *Agent) checkArgs(args ...interface{}) (map[string]interface{}, interface{}, error) {
	var pm map[string]interface{}
	var v interface{}
	for _, arg := range args {
		switch reflect.TypeOf(arg).Kind() {
		case reflect.Map:
			if m, err := jcast.StringMapInterface(arg); err != nil {
				return nil, nil, err
			} else {
				pm = m
			}
		case reflect.Struct:
			var err error
			var b []byte
			if b, err = json.Marshal(&arg); err != nil {
				return nil, nil, err
			}
			pm = make(map[string]interface{})
			if err = jfile.Decode(jfile.Json.String(), b, pm); err != nil {
				return nil, nil, err
			}
		case reflect.Ptr:
			v = arg
		}
	}
	return pm, v, nil
}

func (a *Agent) queryPrepareMS(single bool, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	pm := make([]map[string]interface{}, 1)
	pm[0] = param
	if query, _, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	return a.queryPrepare(single, query, args...)
}

func (a *Agent) queryPrepareTxMS(single bool, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	if query, _, err = a.xmlAndParamsToQueryAndArgs(Select, id, param); err != nil {
		return nil, err
	}
	return a.queryPrepareTx(single, query, args...)
}

func (a *Agent) execOps(ops Operations, id string, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	if param, _, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	var query string
	if query, args, err = a.xmlAndParamsToQueryAndArgs(ops, id, param); err != nil {
		return nil, err
	}
	return a.exec(query, args...)
}

func (a *Agent) execTxOps(ops Operations, id string, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	if param, _, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	var query string
	if query, args, err = a.xmlAndParamsToQueryAndArgs(ops, id, param); err != nil {
		return nil, err
	}
	return a.execTx(query, args...)
}

func (a *Agent) execPrepareOps(ops Operations, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	if query, _, err = a.xmlAndParamsToQueryAndArgs(ops, id, param); err != nil {
		return nil, err
	}
	return a.execPrepare(query, args...)
}

func (a *Agent) execPrepareTxOps(ops Operations, id string, param map[string]interface{}, args ...[]interface{}) (result []Result, err error) {
	var query string
	if query, _, err = a.xmlAndParamsToQueryAndArgs(ops, id, param); err != nil {
		return nil, err
	}
	return a.execPrepareTx(query, args...)
}

func (a *Agent) getXmlAndParam(ops Operations, id string, param []map[string]interface{}) (elem *element, pm map[string]interface{}, err error) {
	if elem, err = getElement(ops, id); err != nil {
		return nil, nil, err
	}
	if len(param) > 0 {
		pm = param[0]
	} else {
		pm = nil
	}
	return elem, pm, nil
}

func (a *Agent) xmlAndParamsToQueryAndArgs(ops Operations, id string, param map[string]interface{}) (query string, args []interface{}, err error) {
	var elem *element
	if elem, err = getElement(ops, id); err != nil {
		return "", nil, err
	}
	if query, _, err = elem.getSql(param, false); err != nil {
		return "", nil, err
	}
	query, args = a.getQueryAndArgs(query, param)
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
		return nil, errorStr(errorDBNil)
	}
	var rows *sql.Rows
	subject.Next(query)
	if rows, err = a.db.Query(query, args...); err != nil {
		return nil, err
	}
	return a.getResult(rows, single)
}

func (a *Agent) queryTx(single bool, query string, args ...interface{}) (result Result, err error) {
	if a.tx == nil {
		return nil, errorStr(errorDbNotBegin)
	}
	var rows *sql.Rows
	subject.Next(query)
	if rows, err = a.tx.Query(query, args...); err != nil {
		return nil, err
	}
	return a.getResult(rows, single)
}

func (a *Agent) queryPage(ct bool, id string, start, end int64, args ...interface{}) (result Result, err error) {
	var param map[string]interface{}
	var v interface{}
	if param, v, err = a.checkArgs(args...); err != nil {
		return nil, err
	}
	var elem *element
	if elem, err = getElement(Select, id); err != nil {
		return nil, err
	}
	var query string
	var order string
	if query, order, err = elem.getSql(param, true); err != nil {
		return nil, err
	}
	pageQuery, countQuery := getPageSql(a.t, query, order, start, end)
	pageQuery, args = a.getQueryAndArgs(pageQuery, param)
	countQuery, _ = a.getQueryAndArgs(countQuery, param)
	if result, err = a.queryPageWithSql(ct, pageQuery, countQuery, start, end, args...); err != nil {
		return result, err
	}
	if v != nil {
		m := map[string]interface{}{
			"Rows":        result.Rows(),
			"RowStart":    result.RowStart(),
			"RowEnd":      result.RowEnd(),
			"TotalRecord": result.TotalRecord(),
		}
		if err = jfile.Convert(m, v); err != nil {
			return result, err
		}
	}
	return result, nil
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
		return nil, err
	}
	if res, ok := resPage.(agentResult); ok {
		api := allowPagingId
		obi := orderById
		if a.t == PostgreSql {
			api = strings.ToLower(api)
			obi = strings.ToLower(obi)
		}
		for i := range res.rows {
			delete(res.rows[i], api)
			delete(res.rows[i], obi)
		}
	}
	if resCount, err = a.queryTx(true, countQuery, args...); err != nil {
		return nil, err
	}
	if ct {
		if err = a.Commit(); err != nil {
			return nil, err
		}
	}
	if len(resCount.Rows()) <= 0 {
		return nil, errorStr(errorNoRowsAvailable)
	}
	var m map[string]interface{}
	m = resCount.Rows()[0]
	var total int64
	tr := totalRecord
	if a.t == PostgreSql {
		tr = strings.ToLower(tr)
	}
	if total, err = jcast.Int64(m[tr]); err != nil {
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

func (a *Agent) queryRowScan(query string, data interface{}, args ...interface{}) error {
	subject.Next(query)
	return a.db.QueryRow(query, args...).Scan(data)
}

func (a *Agent) queryRowScanTx(query string, data interface{}, args ...interface{}) error {
	subject.Next(query)
	return a.tx.QueryRow(query, args...).Scan(data)
}

func (a *Agent) exec(query string, args ...interface{}) (result Result, err error) {
	if a.db == nil {
		return nil, errorStr(errorDBNil)
	}
	var res sql.Result
	subject.Next(query)
	if res, err = a.db.Exec(query, args...); err != nil {
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
		return nil, errorStr(errorDbNotBegin)
	}
	var res sql.Result
	subject.Next(query)
	if res, err = a.tx.Exec(query, args...); err != nil {
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
		return nil, errorStr(errorDBNil)
	}
	var stmt *sql.Stmt
	subject.Next(query)
	if stmt, err = a.db.Prepare(query); err != nil {
		return nil, err
	}
	return a.stmtQuery(single, stmt, args...)
}

func (a *Agent) queryPrepareTx(single bool, query string, args ...[]interface{}) (result []Result, err error) {
	if a.tx == nil {
		return nil, errorStr(errorDbNotBegin)
	}
	var stmt *sql.Stmt
	subject.Next(query)
	if stmt, err = a.tx.Prepare(query); err != nil {
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
		return nil, errorStr(errorDBNil)
	}
	var stmt *sql.Stmt
	subject.Next(query)
	if stmt, err = a.db.Prepare(query); err != nil {
		return nil, err
	}
	return a.stmtExec(stmt, args...)
}

func (a *Agent) execPrepareTx(query string, args ...[]interface{}) (result []Result, err error) {
	if a.tx == nil {
		return nil, errorStr(errorDbNotBegin)
	}
	var stmt *sql.Stmt
	subject.Next(query)
	if stmt, err = a.tx.Prepare(query); err != nil {
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

func (a *Agent) getSqlAndArgs(ops Operations, id string, args ...interface{}) (string, []interface{}, error) {
	if param, _, err := a.checkArgs(args...); err != nil {
		return "", nil, err
	} else {
		return a.xmlAndParamsToQueryAndArgs(ops, id, param)
	}
}

func (a *Agent) getResult(rows *sql.Rows, single bool) (result Result, err error) {
	if rows == nil {
		err = errorStr(errorRowsNil)
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
			if colType.ScanType() == nil {
				rowValue[i] = nil
			} else {
				rowValue[i] = reflect.New(colType.ScanType())
			}
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
			if colType.ScanType() == nil {
				record[colType.Name()] = rowValue[i]
			} else {
				dbType := colType.DatabaseTypeName()
				switch colType.ScanType().String() {
				case "time.Time":
					fallthrough
				case "sql.NullTime":
					var ok bool
					if record[colType.Name()], ok = rowValue[i].(time.Time); !ok {
						record[colType.Name()] = rowValue[i]
					}
				case "int32":
					fallthrough
				case "sql.NullInt32":
					switch rowValue[i].(type) {
					case []byte:
						var err error
						if record[colType.Name()], err = strconv.ParseInt(string(rowValue[i].([]byte)), 10, 32); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					default:
						record[colType.Name()] = rowValue[i]
					}
				case "int64":
					fallthrough
				case "sql.NullInt64":
					switch rowValue[i].(type) {
					case []byte:
						var err error
						if record[colType.Name()], err = strconv.ParseInt(string(rowValue[i].([]byte)), 10, 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					default:
						record[colType.Name()] = rowValue[i]
					}
				case "sql.RawBytes":
					switch dbType {
					case "DOUBLE":
						var err error
						if record[colType.Name()], err = strconv.ParseFloat(string(rowValue[i].([]byte)), 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					case "DECIMAL":
						var err error
						if record[colType.Name()], err = strconv.ParseFloat(string(rowValue[i].([]byte)), 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					case "NUMERIC":
						var err error
						if record[colType.Name()], err = strconv.ParseFloat(string(rowValue[i].([]byte)), 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					case "BLOB":
						var ok bool
						if record[colType.Name()], ok = rowValue[i].([]byte); !ok {
							record[colType.Name()] = rowValue[i]
						}
					case "JSON":
						if b, ok := rowValue[i].([]byte); ok {
							var j interface{}
							if err := json.Unmarshal(b, &j); err != nil {
								record[colType.Name()] = string(b)
							} else {
								switch j.(type) {
								case []interface{}:
									if record[colType.Name()], ok = j.([]interface{}); !ok {
										record[colType.Name()] = string(b)
									}
								case map[string]interface{}:
									if record[colType.Name()], ok = j.(map[string]interface{}); !ok {
										record[colType.Name()] = string(b)
									}
								default:
									record[colType.Name()] = string(b)
								}
							}
						} else {
							record[colType.Name()] = rowValue[i]
						}
					default:
						if b, ok := rowValue[i].([]byte); !ok {
							record[colType.Name()] = rowValue[i]
						} else {
							record[colType.Name()] = string(b)
						}
					}
				case "sql.NullFloat64":
					fallthrough
				case "interface {}":
					fallthrough
				case "[]uint8":
					switch dbType {
					case "DOUBLE":
						var err error
						if record[colType.Name()], err = strconv.ParseFloat(string(rowValue[i].([]uint8)), 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					case "DECIMAL":
						var err error
						if record[colType.Name()], err = strconv.ParseFloat(string(rowValue[i].([]uint8)), 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					case "NUMERIC":
						var err error
						if record[colType.Name()], err = strconv.ParseFloat(string(rowValue[i].([]uint8)), 64); err != nil {
							record[colType.Name()] = rowValue[i]
						}
					case "JSON":
						if b, ok := rowValue[i].([]byte); ok {
							var j interface{}
							if err := json.Unmarshal(b, &j); err != nil {
								record[colType.Name()] = string(b)
							} else {
								switch j.(type) {
								case []interface{}:
									if record[colType.Name()], ok = j.([]interface{}); !ok {
										record[colType.Name()] = string(b)
									}
								case map[string]interface{}:
									if record[colType.Name()], ok = j.(map[string]interface{}); !ok {
										record[colType.Name()] = string(b)
									}
								default:
									record[colType.Name()] = string(b)
								}
							}
						} else {
							record[colType.Name()] = rowValue[i]
						}
					default:
						record[colType.Name()] = rowValue[i]
					}
				case "godror.Number":
					var err error
					if record[colType.Name()], err = strconv.ParseFloat(jcast.String(rowValue[i]), 64); err != nil {
						record[colType.Name()] = rowValue[i]
					}
				default:
					record[colType.Name()] = rowValue[i]
				}
			}
		} else {
			record[colType.Name()] = nil
		}
	}
	return record
}
