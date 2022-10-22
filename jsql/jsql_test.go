// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	/*_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-oci8"*/
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jtime"
	"strconv"
	"testing"
	"time"
)

func TestSql(t *testing.T) {
	testErr := "TEST ERROR:"
	SetFormat(jfile.Json)
	SetEnvFileName("")
	inEnvKey := ""
	SetEnvKey(inEnvKey)
	outEnvKey := EnvKey()
	assert.Equal(t, inEnvKey, outEnvKey, fmt.Sprintf("%v != %v", outEnvKey, inEnvKey))
	inEnvVal := ""
	SetEnvVal(inEnvVal)
	outEnvVal := EnvVal()
	assert.Equal(t, inEnvVal, outEnvVal, fmt.Sprintf("%v != %v", outEnvVal, inEnvVal))
	DisableEnv()
	EnableEnv()
	SubscribeSql(func(i ...interface{}) {})
	SetDecodeFunc(decodeFn)
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("../files/test-jconf-error.json")
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("../files/test-jconf.json")
	if err := Init(); err != nil {
		t.Error(err)
	}
	testSql1(t)
	testSql2(t)
	testSql3(t)
	testSql4(t)
	testSql5(t)
	testSql6(t)
	testSqlInsert1(t)
	testSqlUpdate1(t)
	testSqlDelete1(t)
	testSqlOther1(t)
	testSqlOther2(t)
	testSqlTable1(t)
	testSqlTable2(t)
	testSqlTable3(t)
	testSqlTable4(t)
	testSqlTable5(t)
	testSqlTable6(t)
	testSqlTable7(t)
	testSqlTable8(t)
	testSqlTable9(t)
	testSqlTable10(t)
	testSqlTable11(t)
	testSqlTable12(t)
	testSqlTable13(t)
	testSqlTable14(t)
	testSqlTable15(t)
	testSqlTable16(t)
	testSqlTable17(t)
	testSqlTable18(t)
	testSqlTable19(t)
	testSqlTable20(t)
	/*test1(t)
	test2(t)
	test3(t)
	test4(t)
	test5(t)
	test6(t)
	testInsert1(t)
	testUpdate1(t)
	testDelete1(t)
	testOther1(t)
	testOther2(t)
	testTable1(t)
	testTable2(t)
	testTable3(t)
	testTable4(t)
	testTable5(t)
	testTable6(t)
	testTable7(t)
	testTable8(t)
	testTable9(t)
	testTable10(t)
	testTable11(t)
	testTable12(t)
	testTable13(t)
	testTable14(t)
	testTable15(t)
	testTable16(t)
	testTable17(t)
	testTable18(t)
	testTable19(t)
	testTable20(t)*/
}

func decodeFn(str string) (string, error) {
	return str, nil
}

func testSql1(t *testing.T) {
	tests := []struct {
		a          *Agent
		id         string
		pm         map[string]interface{}
		pageQuery  string
		countQuery string
		args       []interface{}
	}{
		{
			&Agent{t: MySql},
			"testMySql1",
			nil,
			"SELECT * FROM (SELECT (@i := @i + 1) AS ALLOWPAGINGID, table1.* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT * FROM USR)  as  tbs1 ) as table1, (select @i := 0) temp ORDER BY ORDERBYID DESC ) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10",
			"SELECT COUNT(1) as TOTALRECORD FROM (SELECT * FROM USR) data",
			make([]interface{}, 0),
		},
		{
			&Agent{t: MSSql},
			"testMSSql1",
			nil,
			"SELECT * FROM (SELECT ROW_NUMBER() OVER(ORDER BY ORDERBYID DESC) AS ALLOWPAGINGID,* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT * FROM USR) as tbs1) as table1) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10",
			"SELECT COUNT(1) as TOTALRECORD FROM (SELECT * FROM USR) data",
			make([]interface{}, 0),
		},
		{
			&Agent{t: Oracle},
			"testOracle1",
			nil,
			"SELECT t3.* FROM (SELECT t2.*, rownum as ALLOWPAGINGID FROM (SELECT t1.*, 1 as ORDERBYID FROM (SELECT * FROM M_USER) t1) t2 ORDER BY ORDERBYID) t3 WHERE ALLOWPAGINGID BETWEEN 6 AND 10",
			"SELECT COUNT(1) as TOTALRECORD FROM (SELECT * FROM M_USER) data",
			make([]interface{}, 0),
		},
	}
	for _, v := range tests {
		var elem *element
		var err error
		if elem, err = getElement(Select, v.id); err != nil {
			t.Error(err)
			continue
		}
		var query string
		var order string
		if query, order, err = elem.getSql(v.pm, true); err != nil {
			t.Error(err)
			continue
		}
		var args []interface{}
		pageQuery, countQuery := getPageSql(v.a.t, query, order, 6, 10)
		pageQuery, args = v.a.getQueryAndArgs(pageQuery, v.pm)
		countQuery, _ = v.a.getQueryAndArgs(countQuery, v.pm)
		assert.Equal(t, pageQuery, v.pageQuery, fmt.Sprintf("%v != %v", pageQuery, v.pageQuery))
		assert.Equal(t, countQuery, v.countQuery, fmt.Sprintf("%v != %v", countQuery, v.countQuery))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSql2(t *testing.T) {
	pm1 := make(map[string]interface{})
	pm1["USR_STS"] = "1"
	pm1["USR_ID"] = "Admin"
	pm2 := make(map[string]interface{})
	pm2["USE_STS"] = "1"
	pm2["USR_ID"] = "2531222221"
	pm3 := make(map[string]interface{})
	pm3["USER_STATUS"] = "Y"
	pm3["USER_ID"] = "Z12345678"
	tests := []struct {
		a     *Agent
		id    string
		pm    map[string]interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"testMySql2",
			pm1,
			`SELECT * FROM USR WHERE USR_STS = ? AND USR_ID = ?`,
			[]interface{}{pm1["USR_STS"], pm1["USR_ID"]},
		},
		{
			&Agent{t: MSSql},
			"testMSSql2",
			pm2,
			`SELECT * FROM USR WHERE USE_STS = @p1 AND USR_ID = @p2`,
			[]interface{}{pm2["USE_STS"], pm2["USR_ID"]},
		},
		{
			&Agent{t: Oracle},
			"testOracle2",
			pm3,
			`SELECT * FROM M_USER WHERE USER_STATUS = :0 AND USER_ID = :1`,
			[]interface{}{pm3["USER_STATUS"], pm3["USER_ID"]},
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Select, v.id, v.pm); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSql3(t *testing.T) {
	type param struct {
		COL1 string
		COL2 string
		SORT string
	}
	pm1 := param{"USR_ID", "USR_STS", "USR_ID"}
	pm2 := param{"USR_ID", "USE_STS", "USR_ID"}
	pm3 := param{"USER_ID", "USER_STATUS", "USER_ID"}
	type u1 struct {
		USR_ID  string
		USR_STS string
	}
	type v1 struct {
		Rows []u1
	}
	var l1 v1
	type u2 struct {
		USR_ID  string
		USE_STS string
	}
	type v2 struct {
		Rows []u2
	}
	var l2 v2
	type u3 struct {
		USER_ID     string
		USER_STATUS string
	}
	type v3 struct {
		Rows []u3
	}
	var l3 v3
	tests := []struct {
		a          *Agent
		id         string
		pm         param
		v          interface{}
		pageQuery  string
		countQuery string
		args       []interface{}
	}{
		{
			&Agent{t: MySql},
			"testMySql3",
			pm1,
			&l1,
			`SELECT * FROM (SELECT (@i := @i + 1) AS ALLOWPAGINGID, table1.* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT USR_ID, USR_STS FROM USR ORDER BY USR_ID DESC)  as  tbs1 ) as table1, (select @i := 0) temp ORDER BY ORDERBYID DESC ) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT USR_ID, USR_STS FROM USR) data`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: MSSql},
			"testMSSql3",
			pm2,
			&l2,
			`SELECT * FROM (SELECT ROW_NUMBER() OVER( ORDER BY USR_ID DESC) AS ALLOWPAGINGID,* FROM (SELECT * FROM (SELECT USR_ID, USE_STS FROM USR) as tbs1) as table1) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT USR_ID, USE_STS FROM USR) data`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: Oracle},
			"testOracle3",
			pm3,
			&l3,
			`SELECT t3.* FROM (SELECT t2.*, rownum as ALLOWPAGINGID FROM (SELECT t1.*, 1 as ORDERBYID FROM (SELECT USER_ID, USER_STATUS FROM M_USER ORDER BY USER_ID DESC) t1) t2 ORDER BY ORDERBYID) t3 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT USER_ID, USER_STATUS FROM M_USER) data`,
			make([]interface{}, 0),
		},
	}
	for _, v := range tests {
		var err error
		var pm map[string]interface{}
		if pm, _, err = v.a.checkArgs(v.pm, v.v); err != nil {
			t.Error(err)
			continue
		}
		var elem *element
		if elem, err = getElement(Select, v.id); err != nil {
			t.Error(err)
			continue
		}
		var query string
		var order string
		if query, order, err = elem.getSql(pm, true); err != nil {
			t.Error(err)
			continue
		}
		var args []interface{}
		pageQuery, countQuery := getPageSql(v.a.t, query, order, 6, 10)
		pageQuery, args = v.a.getQueryAndArgs(pageQuery, pm)
		countQuery, _ = v.a.getQueryAndArgs(countQuery, pm)
		assert.Equal(t, pageQuery, v.pageQuery, fmt.Sprintf("%v != %v", pageQuery, v.pageQuery))
		assert.Equal(t, countQuery, v.countQuery, fmt.Sprintf("%v != %v", countQuery, v.countQuery))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSql4(t *testing.T) {
	pm1 := make(map[string]interface{})
	l1 := []string{"USR_ID", "USR_STS"}
	pm1["list"] = l1
	pm2 := make(map[string]interface{})
	l2 := []string{"USR_ID", "USE_STS"}
	pm2["list"] = l2
	pm3 := make(map[string]interface{})
	l3 := []string{"USER_ID", "USER_STATUS"}
	pm3["list"] = l3
	tests := []struct {
		a          *Agent
		id         string
		pm         map[string]interface{}
		pageQuery  string
		countQuery string
		args       []interface{}
	}{
		{
			&Agent{t: MySql},
			"testMySql4",
			pm1,
			`SELECT * FROM (SELECT (@i := @i + 1) AS ALLOWPAGINGID, table1.* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT USR_ID, USR_STS FROM USR)  as  tbs1 ) as table1, (select @i := 0) temp ORDER BY ORDERBYID DESC ) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT USR_ID, USR_STS FROM USR) data`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: MSSql},
			"testMSSql4",
			pm2,
			`SELECT * FROM (SELECT ROW_NUMBER() OVER(ORDER BY ORDERBYID DESC) AS ALLOWPAGINGID,* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT USR_ID, USE_STS FROM USR) as tbs1) as table1) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT USR_ID, USE_STS FROM USR) data`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: Oracle},
			"testOracle4",
			pm3,
			`SELECT t3.* FROM (SELECT t2.*, rownum as ALLOWPAGINGID FROM (SELECT t1.*, 1 as ORDERBYID FROM (SELECT USER_ID, USER_STATUS FROM M_USER) t1) t2 ORDER BY ORDERBYID) t3 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT USER_ID, USER_STATUS FROM M_USER) data`,
			make([]interface{}, 0),
		},
	}
	for _, v := range tests {
		var elem *element
		var err error
		if elem, err = getElement(Select, v.id); err != nil {
			t.Error(err)
			continue
		}
		var query string
		var order string
		if query, order, err = elem.getSql(v.pm, true); err != nil {
			t.Error(err)
			continue
		}
		var args []interface{}
		pageQuery, countQuery := getPageSql(v.a.t, query, order, 6, 10)
		pageQuery, args = v.a.getQueryAndArgs(pageQuery, v.pm)
		countQuery, _ = v.a.getQueryAndArgs(countQuery, v.pm)
		assert.Equal(t, pageQuery, v.pageQuery, fmt.Sprintf("%v != %v", pageQuery, v.pageQuery))
		assert.Equal(t, countQuery, v.countQuery, fmt.Sprintf("%v != %v", countQuery, v.countQuery))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSql5(t *testing.T) {
	pm1 := make(map[string]interface{})
	pm1["TYPE"] = "MySql"
	pm1["AS_NAME"] = false
	pm1["TABLE"] = "USR"
	l1 := []string{"USR_SEQ", "USR_ID"}
	pm1["MySqlList"] = l1
	pm2 := make(map[string]interface{})
	pm2["TYPE"] = "MSSql"
	pm2["AS_NAME"] = true
	pm2["TABLE"] = "USR"
	m2 := make(map[string]string)
	m2["NEW_USR_SEQ"] = "USR_SEQ"
	m2["NEW_USR_ID"] = "USR_ID"
	pm2["MSSqlList"] = m2
	pm3 := make(map[string]interface{})
	pm3["TYPE"] = "Oracle"
	pm3["AS_NAME"] = true
	pm3["TABLE"] = "M_USER"
	m3 := make(map[string]string)
	m3["NEW_USER_ID"] = "USER_ID"
	m3["NEW_USER_NAME"] = "USER_NAME"
	pm3["OracleList"] = m3
	tests := []struct {
		a    *Agent
		pm   map[string]interface{}
		args []interface{}
	}{
		{
			&Agent{t: MySql},
			pm1,
			[]interface{}{},
		},
		{
			&Agent{t: MSSql},
			pm2,
			[]interface{}{},
		},
		{
			&Agent{t: Oracle},
			pm3,
			[]interface{}{},
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Select, "testSelect5", v.pm); err != nil {
			t.Error(err)
			continue
		}
		fmt.Println(query)
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSql6(t *testing.T) {
	pm1 := make(map[string]interface{})
	pm1["TABLE"] = "USR"
	pm1["COL"] = "USR_SEQ"
	pm1["VAL"] = 1
	pm1["TEST"] = true
	pm1["TEST2"] = false
	pm2 := make(map[string]interface{})
	pm2["TABLE"] = "USR"
	pm2["COL"] = "USR_SEQ"
	pm2["VAL"] = 1
	pm2["TEST"] = true
	pm2["TEST2"] = false
	pm3 := make(map[string]interface{})
	pm3["TABLE"] = "M_USER"
	pm3["COL"] = "USER_ID"
	pm3["VAL"] = "Z00000000"
	pm3["TEST"] = true
	pm3["TEST2"] = false
	tests := []struct {
		a     *Agent
		id    string
		pm    map[string]interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"testSelect6",
			pm1,
			`SELECT * FROM USR WHERE USR_SEQ = ? ORDER BY USR_SEQ DESC`,
			[]interface{}{pm1["VAL"]},
		},
		{
			&Agent{t: MSSql},
			"testSelect6",
			pm2,
			`SELECT * FROM USR WHERE USR_SEQ = @p1 ORDER BY USR_SEQ DESC`,
			[]interface{}{pm2["VAL"]},
		},
		{
			&Agent{t: Oracle},
			"testSelect6",
			pm3,
			`SELECT * FROM M_USER WHERE USER_ID = :0 ORDER BY USER_ID DESC`,
			[]interface{}{pm3["VAL"]},
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Select, "testSelect6", v.pm); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSqlInsert1(t *testing.T) {
	pm1 := make(map[string]interface{})
	l1 := []string{"USR_SEQ", "USR_ID", "LAST_TIME"}
	pm1["list"] = l1
	pm1["TABLE"] = "USR"
	pm1["USR_SEQ"] = 777
	pm1["USR_ID"] = "test insert"
	pm1["LAST_TIME"] = time.Now()
	pm2 := make(map[string]interface{})
	l2 := []string{"USR_SEQ", "USR_ID", "LAST_TIME"}
	pm2["list"] = l2
	pm2["TABLE"] = "USR"
	pm2["USR_SEQ"] = 777
	pm2["USR_ID"] = "test insert"
	pm2["LAST_TIME"] = time.Now()
	pm3 := make(map[string]interface{})
	l3 := []string{"USER_ID", "LAST_LOGIN_DT"}
	pm3["list"] = l3
	pm3["TABLE"] = "M_USER"
	pm3["USER_ID"] = "test insert"
	pm3["LAST_LOGIN_DT"] = time.Now()
	tests := []struct {
		a     *Agent
		pm    map[string]interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			pm1,
			`INSERT INTO USR (USR_SEQ, USR_ID, LAST_TIME) VALUES (?, ?, ?)`,
			[]interface{}{pm1["USR_SEQ"], pm1["USR_ID"], pm1["LAST_TIME"]},
		},
		{
			&Agent{t: MSSql},
			pm2,
			`INSERT INTO USR (USR_SEQ, USR_ID, LAST_TIME) VALUES (@p1, @p2, @p3)`,
			[]interface{}{pm2["USR_SEQ"], pm2["USR_ID"], pm2["LAST_TIME"]},
		},
		{
			&Agent{t: Oracle},
			pm3,
			`INSERT INTO M_USER (USER_ID, LAST_LOGIN_DT) VALUES (:0, :1)`,
			[]interface{}{pm3["USER_ID"], pm3["LAST_LOGIN_DT"]},
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Insert, "testInsert1", v.pm); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSqlUpdate1(t *testing.T) {
	pm1 := make(map[string]interface{})
	pm1["TABLE"] = "USR"
	pm1["COL1"] = "USR_ID"
	pm1["USR_ID"] = "test update"
	pm1["COL2"] = "USR_SEQ"
	pm1["USR_SEQ"] = 777
	pm2 := make(map[string]interface{})
	pm2["TABLE"] = "USR"
	pm2["COL1"] = "USR_ID"
	pm2["USR_ID"] = "test update"
	pm2["COL2"] = "USR_SEQ"
	pm2["USR_SEQ"] = 777
	pm3 := make(map[string]interface{})
	pm3["TABLE"] = "M_USER"
	pm3["COL1"] = "USER_NAME"
	pm3["USER_NAME"] = "test update"
	pm3["COL2"] = "USER_ID"
	pm3["USER_ID"] = "test insert"
	tests := []struct {
		a     *Agent
		pm    map[string]interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			pm1,
			`UPDATE USR SET USR_ID=? WHERE USR_SEQ=?`,
			[]interface{}{pm1["USR_ID"], pm1["USR_SEQ"]},
		},
		{
			&Agent{t: MSSql},
			pm2,
			`UPDATE USR SET USR_ID=@p1 WHERE USR_SEQ=@p2`,
			[]interface{}{pm2["USR_ID"], pm2["USR_SEQ"]},
		},
		{
			&Agent{t: Oracle},
			pm3,
			`UPDATE M_USER SET USER_NAME=:0 WHERE USER_ID=:1`,
			[]interface{}{pm3["USER_NAME"], pm3["USER_ID"]},
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Update, "testUpdate1", v.pm); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSqlDelete1(t *testing.T) {
	pm1 := make(map[string]interface{})
	pm1["TABLE"] = "USR"
	pm1["COL"] = "USR_SEQ"
	pm1["USR_SEQ"] = 777
	pm2 := make(map[string]interface{})
	pm2["TABLE"] = "USR"
	pm2["COL"] = "USR_SEQ"
	pm2["USR_SEQ"] = 777
	pm3 := make(map[string]interface{})
	pm3["TABLE"] = "M_USER"
	pm3["COL"] = "USER_ID"
	pm3["USER_ID"] = "test insert"
	tests := []struct {
		a     *Agent
		pm    map[string]interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			pm1,
			`DELETE FROM USR WHERE USR_SEQ=?`,
			[]interface{}{pm1["USR_SEQ"]},
		},
		{
			&Agent{t: MSSql},
			pm2,
			`DELETE FROM USR WHERE USR_SEQ=@p1`,
			[]interface{}{pm2["USR_SEQ"]},
		},
		{
			&Agent{t: Oracle},
			pm3,
			`DELETE FROM M_USER WHERE USER_ID=:0`,
			[]interface{}{pm3["USER_ID"]},
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Delete, "testDelete1", v.pm); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSqlOther1(t *testing.T) {
	tests := []struct {
		a     *Agent
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			`CREATE TABLE TEST_CREATE (TEST_COL VARCHAR(255))`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: MSSql},
			`CREATE TABLE TEST_CREATE (TEST_COL VARCHAR(255))`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: Oracle},
			`CREATE TABLE TEST_CREATE (TEST_COL VARCHAR(255))`,
			make([]interface{}, 0),
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Other, "testOther1", nil); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSqlOther2(t *testing.T) {
	tests := []struct {
		a     *Agent
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			`DROP TABLE TEST_CREATE`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: MSSql},
			`DROP TABLE TEST_CREATE`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: Oracle},
			`DROP TABLE TEST_CREATE`,
			make([]interface{}, 0),
		},
	}
	for _, v := range tests {
		var query string
		var args []interface{}
		var err error
		if query, args, err = v.a.xmlAndParamsToQueryAndArgs(Other, "testOther2", nil); err != nil {
			t.Error(err)
			continue
		}
		assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
	}
}

func testSqlTable1(t *testing.T) {
	tests := []struct {
		a          *Agent
		table      string
		pageQuery  string
		countQuery string
		args       []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			`SELECT * FROM (SELECT (@i := @i + 1) AS ALLOWPAGINGID, table1.* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT * FROM USR WHERE 1 = 1)  as  tbs1 ) as table1, (select @i := 0) temp ORDER BY ORDERBYID DESC ) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT * FROM USR WHERE 1 = 1) data`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: MSSql},
			"USR",
			`SELECT * FROM (SELECT ROW_NUMBER() OVER(ORDER BY ORDERBYID DESC) AS ALLOWPAGINGID,* FROM (SELECT *, 1 as ORDERBYID FROM (SELECT * FROM USR WHERE 1 = 1) as tbs1) as table1) as table2 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT * FROM USR WHERE 1 = 1) data`,
			make([]interface{}, 0),
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			`SELECT t3.* FROM (SELECT t2.*, rownum as ALLOWPAGINGID FROM (SELECT t1.*, 1 as ORDERBYID FROM (SELECT * FROM M_USER WHERE 1 = 1) t1) t2 ORDER BY ORDERBYID) t3 WHERE ALLOWPAGINGID BETWEEN 6 AND 10`,
			`SELECT COUNT(1) as TOTALRECORD FROM (SELECT * FROM M_USER WHERE 1 = 1) data`,
			make([]interface{}, 0),
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		if query, args, err := ta.getQuery(); err != nil {
			t.Error(err)
		} else {
			pageQuery, countQuery := getPageSql(v.a.t, query, ta.OrdStr, 6, 10)
			assert.Equal(t, pageQuery, v.pageQuery, fmt.Sprintf("%v != %v", pageQuery, v.pageQuery))
			assert.Equal(t, countQuery, v.countQuery, fmt.Sprintf("%v != %v", countQuery, v.countQuery))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable2(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"USR_SEQ",
			1,
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ = ?`,
			[]interface{}{1},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"USR_SEQ",
			2,
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ = @p1`,
			[]interface{}{2},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			"Z00000000",
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID = :0`,
			[]interface{}{"Z00000000"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.Equal(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable3(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"USR_SEQ",
			1,
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ != ?`,
			[]interface{}{1},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"USR_SEQ",
			2,
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ != @p1`,
			[]interface{}{2},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			"Z00000000",
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID != :0`,
			[]interface{}{"Z00000000"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.NotEqual(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable4(t *testing.T) {
	testErr := "TEST ERROR:"
	tests := []struct {
		err   bool
		a     *Agent
		table string
		col   string
		val   []interface{}
		query string
		args  []interface{}
	}{
		{
			true,
			&Agent{t: MySql},
			"USR",
			"USR_SEQ",
			make([]interface{}, 0),
			``,
			[]interface{}{1, 2},
		},
		{
			false,
			&Agent{t: MySql},
			"USR",
			"USR_SEQ",
			[]interface{}{1, 2},
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ IN (?, ?)`,
			[]interface{}{1, 2},
		},
		{
			false,
			&Agent{t: MSSql},
			"USR",
			"USR_SEQ",
			[]interface{}{1, 2},
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ IN (@p1, @p2)`,
			[]interface{}{1, 2},
		},
		{
			false,
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			[]interface{}{"Z00000000", "Z12345678"},
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID IN (:0, :1)`,
			[]interface{}{"Z00000000", "Z12345678"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.In(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.In must be return error"))
			} else {
				assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
				assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
			}
		}
	}
}

func testSqlTable5(t *testing.T) {
	testErr := "TEST ERROR:"
	tests := []struct {
		err   bool
		a     *Agent
		table string
		col   string
		val   []interface{}
		query string
		args  []interface{}
	}{
		{
			true,
			&Agent{t: MySql},
			"USR",
			"USR_SEQ",
			make([]interface{}, 0),
			``,
			[]interface{}{1},
		},
		{
			false,
			&Agent{t: MySql},
			"USR",
			"USR_SEQ",
			[]interface{}{1},
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ NOT IN (?)`,
			[]interface{}{1},
		},
		{
			false,
			&Agent{t: MSSql},
			"USR",
			"USR_SEQ",
			[]interface{}{2},
			`SELECT * FROM USR WHERE 1 = 1 AND USR_SEQ NOT IN (@p1)`,
			[]interface{}{2},
		},
		{
			false,
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			[]interface{}{"Z00000000"},
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID NOT IN (:0)`,
			[]interface{}{"Z00000000"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.NotIn(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.NotIn must be return error"))
			} else {
				assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
				assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
			}
		}
	}
}

func testSqlTable6(t *testing.T) {
	testErr := "TEST ERROR:"
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	et, _ := time.Parse(jtime.DateTime, "2023-01-01 00:00:00")
	tests := []struct {
		err   bool
		a     *Agent
		table string
		col   string
		val   []interface{}
		query string
		args  []interface{}
	}{
		{
			true,
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			make([]interface{}, 0),
			``,
			[]interface{}{},
		},
		{
			false,
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			[]interface{}{st, et},
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME BETWEEN ? AND ?`,
			[]interface{}{st, et},
		},
		{
			false,
			&Agent{t: MSSql},
			"USR",
			"LAST_TIME",
			[]interface{}{st, et},
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME BETWEEN @p1 AND @p2`,
			[]interface{}{st, et},
		},
		{
			false,
			&Agent{t: Oracle},
			"M_USER",
			"LAST_LOGIN_DT",
			[]interface{}{st, et},
			`SELECT * FROM M_USER WHERE 1 = 1 AND LAST_LOGIN_DT BETWEEN :0 AND :1`,
			[]interface{}{st, et},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.Between(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.Between must be return error"))
			} else {
				assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
				assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
			}
		}
	}
}

func testSqlTable7(t *testing.T) {
	testErr := "TEST ERROR:"
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	et, _ := time.Parse(jtime.DateTime, "2023-01-01 00:00:00")
	tests := []struct {
		err   bool
		a     *Agent
		table string
		col   string
		val   []interface{}
		query string
		args  []interface{}
	}{
		{
			true,
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			make([]interface{}, 0),
			``,
			[]interface{}{},
		},
		{
			false,
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			[]interface{}{st, et},
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME NOT BETWEEN ? AND ?`,
			[]interface{}{st, et},
		},
		{
			false,
			&Agent{t: MSSql},
			"USR",
			"LAST_TIME",
			[]interface{}{st, et},
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME NOT BETWEEN @p1 AND @p2`,
			[]interface{}{st, et},
		},
		{
			false,
			&Agent{t: Oracle},
			"M_USER",
			"LAST_LOGIN_DT",
			[]interface{}{st, et},
			`SELECT * FROM M_USER WHERE 1 = 1 AND LAST_LOGIN_DT NOT BETWEEN :0 AND :1`,
			[]interface{}{st, et},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.NotBetween(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.NotBetween must be return error"))
			} else {
				assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
				assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
			}
		}
	}
}

func testSqlTable8(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		query string
	}{
		{
			&Agent{t: MySql},
			"USR",
			"CT_ID",
			`SELECT * FROM USR WHERE 1 = 1 AND CT_ID IS NULL`,
		},
		{
			&Agent{t: MSSql},
			"USR",
			"CT_ID",
			`SELECT * FROM USR WHERE 1 = 1 AND CT_ID IS NULL`,
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"CREATOR",
			`SELECT * FROM M_USER WHERE 1 = 1 AND CREATOR IS NULL`,
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.IsNull(v.col)
		if query, _, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		}
	}
}

func testSqlTable9(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		query string
	}{
		{
			&Agent{t: MySql},
			"USR",
			"CT_ID",
			`SELECT * FROM USR WHERE 1 = 1 AND CT_ID IS NOT NULL`,
		},
		{
			&Agent{t: MSSql},
			"USR",
			"CT_ID",
			`SELECT * FROM USR WHERE 1 = 1 AND CT_ID IS NOT NULL`,
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"CREATOR",
			`SELECT * FROM M_USER WHERE 1 = 1 AND CREATOR IS NOT NULL`,
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.IsNotNull(v.col)
		if query, _, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
		}
	}
}

func testSqlTable10(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   string
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"USR_ID",
			"A",
			`SELECT * FROM USR WHERE 1 = 1 AND USR_ID LIKE ?`,
			[]interface{}{"%A%"},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"USR_ID",
			"A",
			`SELECT * FROM USR WHERE 1 = 1 AND USR_ID LIKE @p1`,
			[]interface{}{"%A%"},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			"A",
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID LIKE :0`,
			[]interface{}{"%A%"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.Like(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable11(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   string
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"USR_ID",
			"A",
			`SELECT * FROM USR WHERE 1 = 1 AND USR_ID LIKE ?`,
			[]interface{}{"A%"},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"USR_ID",
			"A",
			`SELECT * FROM USR WHERE 1 = 1 AND USR_ID LIKE @p1`,
			[]interface{}{"A%"},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			"A",
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID LIKE :0`,
			[]interface{}{"A%"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.SLike(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable12(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   string
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"USR_ID",
			"A",
			`SELECT * FROM USR WHERE 1 = 1 AND USR_ID LIKE ?`,
			[]interface{}{"%A"},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"USR_ID",
			"A",
			`SELECT * FROM USR WHERE 1 = 1 AND USR_ID LIKE @p1`,
			[]interface{}{"%A"},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"USER_ID",
			"A",
			`SELECT * FROM M_USER WHERE 1 = 1 AND USER_ID LIKE :0`,
			[]interface{}{"%A"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.ELike(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable13(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME > ?`,
			[]interface{}{st},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME > @p1`,
			[]interface{}{st},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"LAST_LOGIN_DT",
			st,
			`SELECT * FROM M_USER WHERE 1 = 1 AND LAST_LOGIN_DT > :0`,
			[]interface{}{st},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.Greater(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable14(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME >= ?`,
			[]interface{}{st},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME >= @p1`,
			[]interface{}{st},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"LAST_LOGIN_DT",
			st,
			`SELECT * FROM M_USER WHERE 1 = 1 AND LAST_LOGIN_DT >= :0`,
			[]interface{}{st},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.GreaterThanOrEqual(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable15(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME < ?`,
			[]interface{}{st},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME < @p1`,
			[]interface{}{st},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"LAST_LOGIN_DT",
			st,
			`SELECT * FROM M_USER WHERE 1 = 1 AND LAST_LOGIN_DT < :0`,
			[]interface{}{st},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.Less(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable16(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		a     *Agent
		table string
		col   string
		val   interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME <= ?`,
			[]interface{}{st},
		},
		{
			&Agent{t: MSSql},
			"USR",
			"LAST_TIME",
			st,
			`SELECT * FROM USR WHERE 1 = 1 AND LAST_TIME <= @p1`,
			[]interface{}{st},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			"LAST_LOGIN_DT",
			st,
			`SELECT * FROM M_USER WHERE 1 = 1 AND LAST_LOGIN_DT <= :0`,
			[]interface{}{st},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.LessThanOrEqual(v.col, v.val)
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable17(t *testing.T) {
	now := time.Now()
	m := make(map[string]interface{})
	m["USR_SEQ"] = 777
	m["USR_ID"] = "test insert"
	m["LAST_TIME"] = now
	tests := []struct {
		err   bool
		a     *Agent
		table string
		col   []interface{}
		am    map[string]interface{}
		sm    map[string]interface{}
		args  []interface{}
	}{
		{
			true,
			&Agent{t: MySql},
			"USR",
			[]interface{}{777, 777},
			nil,
			nil,
			[]interface{}{},
		},
		{
			false,
			&Agent{t: MySql},
			"USR",
			nil,
			m,
			nil,
			[]interface{}{777, "test insert", now},
		},
		{
			false,
			&Agent{t: MSSql},
			"USR",
			nil,
			nil,
			m,
			[]interface{}{777, "test insert", now},
		},
		{
			false,
			&Agent{t: Oracle},
			"M_USER",
			[]interface{}{"USER_ID", "test insert", "LAST_LOGIN_DT", now},
			nil,
			nil,
			[]interface{}{"test insert", now},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		if v.col != nil {
			if err := ta.AddColumn(v.col...); err != nil {
				if !v.err {
					t.Error(err)
				}
				continue
			}
		}
		if v.am != nil {
			ta.AddMap(v.am)
			ta.AddMap(v.am)
			ta.SetMap(v.am)
		}
		if v.sm != nil {
			ta.AddMap(v.sm)
			ta.AddMap(v.sm)
			ta.SetMap(v.sm)
		}
		if query, args, err := ta.getInsert(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(query)
			for _, a := range args {
				valid := false
				for _, va := range v.args {
					if a == va {
						valid = true
						break
					}
				}
				if !valid {
					t.Error(fmt.Sprintf("%v != %v", args, v.args))
					break
				}
			}
		}
	}
}

func testSqlTable18(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		col   []interface{}
		pm    []interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			[]interface{}{"USR_ID", "test update"},
			[]interface{}{"USR_SEQ", 777},
			`UPDATE USR SET USR_ID = ? WHERE 1 = 1 AND USR_SEQ = ?`,
			[]interface{}{"test update", 777},
		},
		{
			&Agent{t: MSSql},
			"USR",
			[]interface{}{"USR_ID", "test update"},
			[]interface{}{"USR_SEQ", 777},
			`UPDATE USR SET USR_ID = @p1 WHERE 1 = 1 AND USR_SEQ = @p2`,
			[]interface{}{"test update", 777},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			[]interface{}{"USER_NAME", "test update"},
			[]interface{}{"USER_ID", "test insert"},
			`UPDATE M_USER SET USER_NAME = :0 WHERE 1 = 1 AND USER_ID = :1`,
			[]interface{}{"test update", "test insert"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		if err := ta.AddColumn(v.col...); err != nil {
			t.Error(err)
			continue
		}
		ta.Equal(v.pm[0].(string), v.pm[1])
		if query, args, err := ta.getUpdate(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable19(t *testing.T) {
	tests := []struct {
		a     *Agent
		table string
		pm    []interface{}
		query string
		args  []interface{}
	}{
		{
			&Agent{t: MySql},
			"USR",
			[]interface{}{"USR_SEQ", 777},
			`DELETE FROM USR WHERE 1 = 1 AND USR_SEQ = ?`,
			[]interface{}{777},
		},
		{
			&Agent{t: MSSql},
			"USR",
			[]interface{}{"USR_SEQ", 777},
			`DELETE FROM USR WHERE 1 = 1 AND USR_SEQ = @p1`,
			[]interface{}{777},
		},
		{
			&Agent{t: Oracle},
			"M_USER",
			[]interface{}{"USER_ID", "test insert"},
			`DELETE FROM M_USER WHERE 1 = 1 AND USER_ID = :0`,
			[]interface{}{"test insert"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		ta.Equal(v.pm[0].(string), v.pm[1])
		if query, args, err := ta.getDelete(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func testSqlTable20(t *testing.T) {
	tests := []struct {
		add   bool
		a     *Agent
		table string
		pm    []*Param
		query string
		args  []interface{}
	}{
		{
			false,
			&Agent{t: MySql},
			"USR",
			[]*Param{
				{
					Logic: And,
					Params: []*Param{
						{Logic: And, Col: "USR_SEQ", Opr: Equal, Val: 1},
						{Logic: Or, Col: "USR_SEQ", Opr: Equal, Val: 2},
					},
				},
			},
			`SELECT * FROM USR WHERE 1 = 1 AND (1 = 1 AND USR_SEQ = ? OR USR_SEQ = ?)`,
			[]interface{}{1, 2},
		},
		{
			true,
			&Agent{t: MSSql},
			"USR",
			[]*Param{
				{
					Logic: And,
					Params: []*Param{
						{Logic: And, Col: "USR_SEQ", Opr: Equal, Val: 1},
						{Logic: Or, Col: "USR_SEQ", Opr: Equal, Val: 2},
					},
				},
			},
			`SELECT * FROM USR WHERE 1 = 1 AND (1 = 1 AND USR_SEQ = @p1 OR USR_SEQ = @p2)`,
			[]interface{}{1, 2},
		},
		{
			false,
			&Agent{t: Oracle},
			"M_USER",
			[]*Param{
				{
					Logic: And,
					Params: []*Param{
						{Logic: And, Col: "USER_ID", Opr: Equal, Val: "Z00000000"},
						{Logic: Or, Col: "USER_ID", Opr: Equal, Val: "Z00000001"},
					},
				},
			},
			`SELECT * FROM M_USER WHERE 1 = 1 AND (1 = 1 AND USER_ID = :0 OR USER_ID = :1)`,
			[]interface{}{"Z00000000", "Z00000001"},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{Agent: v.a, Table: v.table}
		if v.add {
			for _, p := range v.pm {
				ta.AddParam(p)
			}
		} else {
			ta.SetParams(v.pm)
		}
		if query, args, err := ta.getQueryAndArgs(); err != nil {
			t.Error(err)
		} else {
			assert.Equal(t, query, v.query, fmt.Sprintf("%v != %v", query, v.query))
			assert.Equal(t, args, v.args, fmt.Sprintf("%v != %v", args, v.args))
		}
	}
}

func test1(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		var res Result
		if res, err = agent.QueryPage("testMySql1", 6, 10); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		var res Result
		if res, err = agent.QueryPage("testMSSql1", 6, 10); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		var res Result
		if res, err = agent.QueryPage("testOracle1", 6, 10); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func test2(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		params := make(map[string]interface{})
		params["USR_STS"] = "1"
		params["USR_ID"] = "Admin"
		var res Result
		if res, err = agent.Query("testMySql2", params); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		params := make(map[string]interface{})
		params["USE_STS"] = "1"
		params["USR_ID"] = "2531222221"
		var res Result
		if res, err = agent.Query("testMSSql2", params); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		params := make(map[string]interface{})
		params["USER_STATUS"] = "Y"
		params["USER_ID"] = "Z12345678"
		var res Result
		if res, err = agent.Query("testOracle2", params); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func test3(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		type param struct {
			COL1 string
			COL2 string
			SORT string
		}
		pm := param{"USR_ID", "USR_STS", "USR_ID"}
		type user struct {
			USR_ID  string
			USR_STS string
		}
		type userList struct {
			Rows []user
		}
		var v userList
		var res Result
		if res, err = agent.QueryPage("testMySql3", 6, 10, pm, &v); err != nil {
			t.Error(err)
		} else {
			fmt.Println(v)
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		type param struct {
			COL1 string
			COL2 string
			SORT string
		}
		pm := param{"USR_ID", "USE_STS", "USR_ID"}
		type user struct {
			USR_ID  string
			USE_STS string
		}
		type userList struct {
			Rows []user
		}
		var v userList
		var res Result
		if res, err = agent.QueryPage("testMSSql3", 6, 10, pm, &v); err != nil {
			t.Error(err)
		} else {
			fmt.Println(v)
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		type param struct {
			COL1 string
			COL2 string
			SORT string
		}
		pm := param{"USER_ID", "USER_STATUS", "USER_ID"}
		type user struct {
			USER_ID     string
			USER_STATUS string
		}
		type userList struct {
			Rows []user
		}
		var v userList
		var res Result
		if res, err = agent.QueryPage("testOracle3", 6, 10, pm, &v); err != nil {
			t.Error(err)
		} else {
			fmt.Println(v)
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func test4(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		params := make(map[string]interface{})
		list := []string{"USR_ID", "USR_STS"}
		params["list"] = list
		var res Result
		if res, err = agent.QueryPage("testMySql4", 6, 10, params); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		params := make(map[string]interface{})
		list := []string{"USR_ID", "USE_STS"}
		params["list"] = list
		var res Result
		if res, err = agent.QueryPage("testMSSql4", 6, 10, params); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		params := make(map[string]interface{})
		list := []string{"USER_ID", "USER_STATUS"}
		params["list"] = list
		var res Result
		if res, err = agent.QueryPage("testOracle4", 6, 10, params); err != nil {
			t.Error(err)
		} else {
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func test5(t *testing.T) {
	params := make(map[string]interface{})
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		params["TYPE"] = "MySql"
		params["AS_NAME"] = false
		params["TABLE"] = "USR"
		list := []string{"USR_SEQ", "USR_ID"}
		params["MySqlList"] = list
		var res Result
		if res, err = agent.QueryRow("testSelect5", params); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		params["TYPE"] = "MSSql"
		params["AS_NAME"] = true
		params["TABLE"] = "USR"
		m := make(map[string]string)
		m["NEW_USR_SEQ"] = "USR_SEQ"
		m["NEW_USR_ID"] = "USR_ID"
		params["MSSqlList"] = m
		var res Result
		if res, err = agent.QueryRow("testSelect5", params); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		params["TYPE"] = "Oracle"
		params["AS_NAME"] = true
		params["TABLE"] = "M_USER"
		m := make(map[string]string)
		m["NEW_USER_ID"] = "USER_ID"
		m["NEW_USER_NAME"] = "USER_NAME"
		params["OracleList"] = m
		var res Result
		if res, err = agent.QueryRow("testSelect5", params); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}

func test6(t *testing.T) {
	params := make(map[string]interface{})
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		params["TABLE"] = "USR"
		params["COL"] = "USR_SEQ"
		params["VAL"] = 1
		params["TEST"] = true
		params["TEST2"] = false
		var res Result
		if res, err = agent.QueryRow("testSelect6", params); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		params["TABLE"] = "USR"
		params["COL"] = "USR_SEQ"
		params["VAL"] = 1
		params["TEST"] = true
		params["TEST2"] = false
		var res Result
		if res, err = agent.QueryRow("testSelect6", params); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		params["TABLE"] = "M_USER"
		params["COL"] = "USER_ID"
		params["VAL"] = "Z00000000"
		params["TEST"] = true
		params["TEST2"] = false
		var res Result
		if res, err = agent.QueryRow("testSelect6", params); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}

func testInsert1(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		list := []string{"USR_SEQ", "USR_ID", "LAST_TIME"}
		pm := make(map[string]interface{})
		pm["list"] = list
		pm["TABLE"] = "USR"
		pm["USR_SEQ"] = 777
		pm["USR_ID"] = "test insert"
		pm["LAST_TIME"] = time.Now()
		var res Result
		if res, err = agent.Insert("testInsert1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
			fmt.Println(res.LastInsertId())
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		list := []string{"USR_SEQ", "USR_ID", "LAST_TIME"}
		pm := make(map[string]interface{})
		pm["list"] = list
		pm["TABLE"] = "USR"
		pm["USR_SEQ"] = 777
		pm["USR_ID"] = "test insert"
		pm["LAST_TIME"] = time.Now()
		var res Result
		if res, err = agent.Insert("testInsert1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
			fmt.Println(res.LastInsertId())
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		list := []string{"USER_ID", "LAST_LOGIN_DT"}
		pm := make(map[string]interface{})
		pm["list"] = list
		pm["TABLE"] = "M_USER"
		pm["USER_ID"] = "test insert"
		pm["LAST_LOGIN_DT"] = time.Now()
		var res Result
		if res, err = agent.Insert("testInsert1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
			fmt.Println(res.LastInsertId())
		}
	}
}

func testUpdate1(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		pm := make(map[string]interface{})
		pm["TABLE"] = "USR"
		pm["COL1"] = "USR_ID"
		pm["USR_ID"] = "test update"
		pm["COL2"] = "USR_SEQ"
		pm["USR_SEQ"] = 777
		var res Result
		if res, err = agent.Update("testUpdate1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		pm := make(map[string]interface{})
		pm["TABLE"] = "USR"
		pm["COL1"] = "USR_ID"
		pm["USR_ID"] = "test update"
		pm["COL2"] = "USR_SEQ"
		pm["USR_SEQ"] = 777
		var res Result
		if res, err = agent.Update("testUpdate1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		pm := make(map[string]interface{})
		pm["TABLE"] = "M_USER"
		pm["COL1"] = "USER_NAME"
		pm["USER_NAME"] = "test update"
		pm["COL2"] = "USER_ID"
		pm["USER_ID"] = "test insert"
		var res Result
		if res, err = agent.Update("testUpdate1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
}

func testDelete1(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		pm := make(map[string]interface{})
		pm["TABLE"] = "USR"
		pm["COL"] = "USR_SEQ"
		pm["USR_SEQ"] = 777
		var res Result
		if res, err = agent.Delete("testDelete1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		pm := make(map[string]interface{})
		pm["TABLE"] = "USR"
		pm["COL"] = "USR_SEQ"
		pm["USR_SEQ"] = 777
		var res Result
		if res, err = agent.Delete("testDelete1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		pm := make(map[string]interface{})
		pm["TABLE"] = "M_USER"
		pm["COL"] = "USER_ID"
		pm["USER_ID"] = "test insert"
		var res Result
		if res, err = agent.Delete("testDelete1", pm); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
}

func testOther1(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		if _, err = agent.Other("testOther1"); err != nil {
			t.Error(err)
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		if _, err = agent.Other("testOther1"); err != nil {
			t.Error(err)
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		if _, err = agent.Other("testOther1"); err != nil {
			t.Error(err)
		}
	}
}

func testOther2(t *testing.T) {
	// MySql
	if agent, err := GetAgent(); err != nil {
		t.Error(err)
	} else {
		if _, err = agent.Other("testOther2"); err != nil {
			t.Error(err)
		}
	}
	// MSSql
	if agent, err := GetAgent("testMSSql"); err != nil {
		t.Error(err)
	} else {
		if _, err = agent.Other("testOther2"); err != nil {
			t.Error(err)
		}
	}
	// Oracle
	if agent, err := GetAgent("testOracle"); err != nil {
		t.Error(err)
	} else {
		if _, err = agent.Other("testOther2"); err != nil {
			t.Error(err)
		}
	}
}

func testTable1(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
	}{
		{"testMySql", "USR"},
		{"testMSSql", "USR"},
		{"testOracle", "M_USER"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		if res, err := ta.QueryPage(6, 10); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.RowStart())
			fmt.Println(res.RowEnd())
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable2(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
		val   interface{}
	}{
		{"testMySql", "USR", "USR_SEQ", 1},
		{"testMSSql", "USR", "USR_SEQ", 2},
		{"testOracle", "M_USER", "USER_ID", "Z00000000"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.Equal(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable3(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
		val   interface{}
	}{
		{"testMySql", "USR", "USR_SEQ", 1},
		{"testMSSql", "USR", "USR_SEQ", 2},
		{"testOracle", "M_USER", "USER_ID", "Z00000000"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.NotEqual(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable4(t *testing.T) {
	testErr := "TEST ERROR:"
	tests := []struct {
		err   bool
		dsKey string
		table string
		col   string
		val   []interface{}
	}{
		{true, "testMySql", "USR", "USR_SEQ", make([]interface{}, 0)},
		{false, "testMySql", "USR", "USR_SEQ", []interface{}{1, 2}},
		{false, "testMSSql", "USR", "USR_SEQ", []interface{}{1, 2}},
		{false, "testOracle", "M_USER", "USER_ID", []interface{}{"Z00000000", "Z12345678"}},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.In(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.In must be return error"))
			} else {
				fmt.Println(res.TotalRecord())
				for _, item := range res.Rows() {
					fmt.Println(item)
				}
			}
		}
	}
}

func testTable5(t *testing.T) {
	testErr := "TEST ERROR:"
	tests := []struct {
		err   bool
		dsKey string
		table string
		col   string
		val   []interface{}
	}{
		{true, "testMySql", "USR", "USR_SEQ", make([]interface{}, 0)},
		{false, "testMySql", "USR", "USR_SEQ", []interface{}{1}},
		{false, "testMSSql", "USR", "USR_SEQ", []interface{}{2}},
		{false, "testOracle", "M_USER", "USER_ID", []interface{}{"Z00000000"}},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.NotIn(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.NotIn must be return error"))
			} else {
				fmt.Println(res.TotalRecord())
				for _, item := range res.Rows() {
					fmt.Println(item)
				}
			}
		}
	}
}

func testTable6(t *testing.T) {
	testErr := "TEST ERROR:"
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	et, _ := time.Parse(jtime.DateTime, "2023-01-01 00:00:00")
	tests := []struct {
		err   bool
		dsKey string
		table string
		col   string
		val   []interface{}
	}{
		{true, "testMySql", "USR", "LAST_TIME", make([]interface{}, 0)},
		{false, "testMySql", "USR", "LAST_TIME", []interface{}{st, et}},
		{false, "testMSSql", "USR", "LAST_TIME", []interface{}{st, et}},
		{false, "testOracle", "M_USER", "LAST_LOGIN_DT", []interface{}{st, et}},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.Between(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.Between must be return error"))
			} else {
				fmt.Println(res.TotalRecord())
				for _, item := range res.Rows() {
					fmt.Println(item)
				}
			}
		}
	}
}

func testTable7(t *testing.T) {
	testErr := "TEST ERROR:"
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	et, _ := time.Parse(jtime.DateTime, "2023-01-01 00:00:00")
	tests := []struct {
		err   bool
		dsKey string
		table string
		col   string
		val   []interface{}
	}{
		{true, "testMySql", "USR", "LAST_TIME", make([]interface{}, 0)},
		{false, "testMySql", "USR", "LAST_TIME", []interface{}{st, et}},
		{false, "testMSSql", "USR", "LAST_TIME", []interface{}{st, et}},
		{false, "testOracle", "M_USER", "LAST_LOGIN_DT", []interface{}{st, et}},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.NotBetween(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			if !v.err {
				t.Error(err)
			}
		} else {
			if v.err {
				t.Error(fmt.Sprint(testErr, " *TableAgent.NotBetween must be return error"))
			} else {
				fmt.Println(res.TotalRecord())
				for _, item := range res.Rows() {
					fmt.Println(item)
				}
			}
		}
	}
}

func testTable8(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
	}{
		{"testMySql", "USR", "CT_ID"},
		{"testMSSql", "USR", "CT_ID"},
		{"testOracle", "M_USER", "CREATOR"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.IsNull(v.col)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable9(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
	}{
		{"testMySql", "USR", "CT_ID"},
		{"testMSSql", "USR", "CT_ID"},
		{"testOracle", "M_USER", "CREATOR"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.IsNotNull(v.col)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable10(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
		val   string
	}{
		{"testMySql", "USR", "USR_ID", "A"},
		{"testMSSql", "USR", "USR_ID", "A"},
		{"testOracle", "M_USER", "USER_ID", "A"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.Like(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable11(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
		val   string
	}{
		{"testMySql", "USR", "USR_ID", "A"},
		{"testMSSql", "USR", "USR_ID", "A"},
		{"testOracle", "M_USER", "USER_ID", "A"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.SLike(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable12(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   string
		val   string
	}{
		{"testMySql", "USR", "USR_ID", "A"},
		{"testMSSql", "USR", "USR_ID", "A"},
		{"testOracle", "M_USER", "USER_ID", "A"},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.ELike(v.col, v.val)
		if res, err := ta.Query(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.TotalRecord())
			for _, item := range res.Rows() {
				fmt.Println(item)
			}
		}
	}
}

func testTable13(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		dsKey string
		table string
		col   string
		val   interface{}
	}{
		{"testMySql", "USR", "LAST_TIME", st},
		{"testMSSql", "USR", "LAST_TIME", st},
		{"testOracle", "M_USER", "LAST_LOGIN_DT", st},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.Greater(v.col, v.val)
		if res, err := ta.QueryRow(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}

func testTable14(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		dsKey string
		table string
		col   string
		val   interface{}
	}{
		{"testMySql", "USR", "LAST_TIME", st},
		{"testMSSql", "USR", "LAST_TIME", st},
		{"testOracle", "M_USER", "LAST_LOGIN_DT", st},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.GreaterThanOrEqual(v.col, v.val)
		if res, err := ta.QueryRow(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}

func testTable15(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		dsKey string
		table string
		col   string
		val   interface{}
	}{
		{"testMySql", "USR", "LAST_TIME", st},
		{"testMSSql", "USR", "LAST_TIME", st},
		{"testOracle", "M_USER", "LAST_LOGIN_DT", st},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.Less(v.col, v.val)
		if res, err := ta.QueryRow(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}

func testTable16(t *testing.T) {
	st, _ := time.Parse(jtime.DateTime, "2022-01-01 00:00:00")
	tests := []struct {
		dsKey string
		table string
		col   string
		val   interface{}
	}{
		{"testMySql", "USR", "LAST_TIME", st},
		{"testMSSql", "USR", "LAST_TIME", st},
		{"testOracle", "M_USER", "LAST_LOGIN_DT", st},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.LessThanOrEqual(v.col, v.val)
		if res, err := ta.QueryRow(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}

func testTable17(t *testing.T) {
	now := time.Now()
	m := make(map[string]interface{})
	m["USR_SEQ"] = 777
	m["USR_ID"] = "test insert"
	m["LAST_TIME"] = now
	tests := []struct {
		err   bool
		dsKey string
		table string
		col   []interface{}
		am    map[string]interface{}
		sm    map[string]interface{}
	}{
		{true, "testMySql", "USR", []interface{}{777, 777}, nil, nil},
		{false, "testMySql", "USR", nil, m, nil},
		{false, "testMSSql", "USR", nil, nil, m},
		{false, "testOracle", "M_USER", []interface{}{"USER_ID", "test insert", "LAST_LOGIN_DT", now}, nil, nil},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		if v.col != nil {
			if err := ta.AddColumn(v.col...); err != nil {
				if !v.err {
					t.Error(err)
				}
				continue
			}
		}
		if v.am != nil {
			ta.AddMap(v.am)
			ta.AddMap(v.am)
			ta.SetMap(v.am)
		}
		if v.sm != nil {
			ta.AddMap(v.sm)
			ta.AddMap(v.sm)
			ta.SetMap(v.sm)
		}
		if res, err := ta.Insert(); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
}

func testTable18(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		col   []interface{}
		pm    []interface{}
	}{
		{"testMySql", "USR", []interface{}{"USR_ID", "test update"}, []interface{}{"USR_SEQ", 777}},
		{"testMSSql", "USR", []interface{}{"USR_ID", "test update"}, []interface{}{"USR_SEQ", 777}},
		{"testOracle", "M_USER", []interface{}{"USER_NAME", "test update"}, []interface{}{"USER_ID", "test insert"}},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		if err := ta.AddColumn(v.col...); err != nil {
			t.Error(err)
			continue
		}
		ta.Equal(v.pm[0].(string), v.pm[1])
		if res, err := ta.Update(); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
}

func testTable19(t *testing.T) {
	tests := []struct {
		dsKey string
		table string
		pm    []interface{}
	}{
		{"testMySql", "USR", []interface{}{"USR_SEQ", 777}},
		{"testMSSql", "USR", []interface{}{"USR_SEQ", 777}},
		{"testOracle", "M_USER", []interface{}{"USER_ID", "test insert"}},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		ta.Equal(v.pm[0].(string), v.pm[1])
		if res, err := ta.Delete(); err != nil {
			t.Error(err)
		} else {
			var af int64
			if af, err = res.RowsAffected(); err != nil {
				t.Error(err)
			} else {
				fmt.Println("Rows Affected: " + strconv.FormatInt(af, 10))
			}
		}
	}
}

func testTable20(t *testing.T) {
	tests := []struct {
		add   bool
		dsKey string
		table string
		pm    []*Param
	}{
		{
			false,
			"testMySql",
			"USR",
			[]*Param{
				{
					Logic: And,
					Params: []*Param{
						{Logic: And, Col: "USR_SEQ", Opr: Equal, Val: 1},
						{Logic: Or, Col: "USR_SEQ", Opr: Equal, Val: 2},
					},
				},
			},
		},
		{
			true,
			"testMSSql",
			"USR",
			[]*Param{
				{
					Logic: And,
					Params: []*Param{
						{Logic: And, Col: "USR_SEQ", Opr: Equal, Val: 1},
						{Logic: Or, Col: "USR_SEQ", Opr: Equal, Val: 2},
					},
				},
			},
		},
		{
			false,
			"testOracle",
			"M_USER",
			[]*Param{
				{
					Logic: And,
					Params: []*Param{
						{Logic: And, Col: "USER_ID", Opr: Equal, Val: "Z00000000"},
						{Logic: Or, Col: "USER_ID", Opr: Equal, Val: "Z00000001"},
					},
				},
			},
		},
	}
	for _, v := range tests {
		ta := &TableAgent{DSKey: v.dsKey, Table: v.table}
		if v.add {
			for _, p := range v.pm {
				ta.AddParam(p)
			}
		} else {
			ta.SetParams(v.pm)
		}
		if res, err := ta.QueryRow(); err != nil {
			t.Error(err)
		} else {
			fmt.Println(res.Row())
		}
	}
}
