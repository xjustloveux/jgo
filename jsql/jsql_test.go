// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	//_ "github.com/denisenkom/go-mssqldb"
	//_ "github.com/go-sql-driver/mysql"
	//_ "github.com/godror/godror"
	//_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"strings"
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
	SetRoot("../files/")
	SetFileName("test-jconf-error.json")
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("test-jconf.json")
	if err := Init(); err != nil {
		t.Error(err)
	}
	//testCreateTable(t)
	//testInsert(t)
	//testInsertTx(t)
	//testUpdate(t)
	//testUpdateTx(t)
	//testDelete(t)
	//testDeleteTx(t)
	//testQuery(t)
	//testQueryTx(t)
	//testQueryPage(t)
	//testQueryPageTx(t)
	//testCount(t)
	//testCountTx(t)
	//testExists(t)
	//testExistsTx(t)
	//testTableSchema(t)
	//testTableSchemaTx(t)
	//testDropTable(t)
}

func decodeFn(str string) (string, error) {
	return str, nil
}

func testCreateTable(t *testing.T) {
	testData := []struct {
		dao     string
		ds      string
		comment []string
	}{
		{"CreateMySqlTable", "testMySql", nil},
		{"CreateMSSqlTable", "testMSSql", nil},
		{"CreateOracleTable", "testOracle",
			[]string{
				"CreateOracleTableComment1",
				"CreateOracleTableComment2",
				"CreateOracleTableComment3",
				"CreateOracleTableComment4",
				"CreateOracleTableComment5",
				"CreateOracleTableComment6",
			},
		},
		{"CreatePostgreSqlTable", "testPostgreSql", nil},
	}
	for _, v := range testData {
		if agent, err := GetAgent(v.ds); err != nil {
			t.Error(err)
		} else {
			if _, err = agent.Other(v.dao); err != nil {
				t.Error(err)
			}
			if v.comment != nil {
				for _, c := range v.comment {
					if _, err = agent.Other(c); err != nil {
						t.Error(err)
					}
				}
			}
		}
	}
}

func testInsert(t *testing.T) {
	pm1 := map[string]interface{}{
		"TABLE": "TEST",
		"COL1":  1,
		"COL2":  1.1,
		"COL3":  time.Now(),
		"COL4":  "one",
		"COL5":  []byte("one"),
		"LIST":  []string{"COL1", "COL2", "COL3", "COL4", "COL5"},
	}
	pm2 := map[string]interface{}{
		"COL1": 2,
		"COL2": 2.2,
		"COL3": time.Now(),
		"COL4": "two",
		"COL5": []byte("two"),
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if _, err = agent.Insert("Insert", pm1); err != nil {
				t.Error(err)
			}
			table := &TableAgent{
				Agent: agent,
				Table: "TEST",
				Col:   pm2,
			}
			if _, err = table.Insert(); err != nil {
				t.Error(err)
			}
		}
	}
}

func testInsertTx(t *testing.T) {
	pm1 := map[string]interface{}{
		"TABLE": "TEST",
		"COL1":  3,
		"COL2":  3.3,
		"COL3":  time.Now(),
		"COL4":  "three",
		"COL5":  []byte("three"),
		"LIST":  []string{"COL1", "COL2", "COL3", "COL4", "COL5"},
	}
	pm2 := map[string]interface{}{
		"COL1": 4,
		"COL2": 4.4,
		"COL3": time.Now(),
		"COL4": "four",
		"COL5": []byte("four"),
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				if _, e := agent.InsertTx("Insert", pm1); e != nil {
					return e
				}
				table := &TableAgent{
					Agent: agent,
					Table: "TEST",
					Col:   pm2,
				}
				if _, e := table.InsertTx(); e != nil {
					return e
				}
				if it, e := jcast.Time("2022-11-21 21:25:11"); e != nil {
					return e
				} else {
					i := 5
					for i < 20 {
						if _, e = agent.InsertTx("Insert", map[string]interface{}{
							"TABLE": "TEST",
							"COL1":  i,
							"COL2":  float64(i) / 10.0,
							"COL3":  it.In(time.UTC),
							"COL4":  fmt.Sprint("test ", i),
							"COL5":  []byte(fmt.Sprint("test ", i)),
							"LIST":  []string{"COL1", "COL2", "COL3", "COL4", "COL5"},
						}); e != nil {
							return e
						}
						i++
					}
				}
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testUpdate(t *testing.T) {
	pm := map[string]interface{}{
		"COL1": 1,
		"COL4": "update one",
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if _, err = agent.Update("Update", pm); err != nil {
				t.Error(err)
			}
			table := &TableAgent{
				Agent:  agent,
				Table:  "TEST",
				Col:    map[string]interface{}{"COL4": "update two"},
				Params: []*Param{{Col: "COL1", Val: 2}},
			}
			if _, err = table.Update(); err != nil {
				t.Error(err)
			}
		}
	}
}

func testUpdateTx(t *testing.T) {
	pm := map[string]interface{}{
		"COL1": 3,
		"COL4": "update three",
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				if _, e := agent.UpdateTx("Update", pm); e != nil {
					return e
				}
				table := &TableAgent{
					Agent:  agent,
					Table:  "TEST",
					Col:    map[string]interface{}{"COL4": "update two"},
					Params: []*Param{{Col: "COL1", Val: 2}},
				}
				if _, e := table.UpdateTx(); e != nil {
					return e
				}
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testDelete(t *testing.T) {
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if _, err = agent.Delete("Delete", map[string]interface{}{"COL1": 7}); err != nil {
				t.Error(err)
			}
			table := &TableAgent{
				Agent:  agent,
				Table:  "TEST",
				Params: []*Param{{Col: "COL1", Val: 11}},
			}
			if _, err = table.Delete(); err != nil {
				t.Error(err)
			}
		}
	}
}

func testDeleteTx(t *testing.T) {
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				if _, e := agent.DeleteTx("Delete", map[string]interface{}{"COL1": 13}); e != nil {
					return e
				}
				table := &TableAgent{
					Agent:  agent,
					Table:  "TEST",
					Params: []*Param{{Col: "COL1", Val: 18}},
				}
				if _, e := table.DeleteTx(); e != nil {
					return e
				}
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testQuery(t *testing.T) {
	var it time.Time
	var err error
	if it, err = jcast.Time("2022-11-21 21:25:11"); err != nil {
		return
	}
	type TestData struct {
		COL1 int64
		COL2 float64
		COL3 time.Time
		COL4 string
		COL5 []byte
	}
	type TestRes struct {
		Rows []TestData
	}
	val := TestData{
		COL1: 5,
		COL2: 0.5,
		COL3: it.In(time.UTC),
		COL4: "test 5",
		COL5: []byte("test 5"),
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			var res TestRes
			var resData TestData
			if _, err = agent.Query("Query", map[string]interface{}{"COL1": 5}, &res); err != nil {
				t.Error(err)
			}
			if len(res.Rows) != 1 {
				assert.Equal(t, val.COL1, res.Rows[0].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[0].COL1))
				assert.Equal(t, val.COL2, res.Rows[0].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[0].COL2))
				assert.Equal(t, val.COL3, res.Rows[0].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[0].COL3))
				assert.Equal(t, val.COL4, res.Rows[0].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[0].COL4))
				assert.Equal(t, val.COL5, res.Rows[0].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[0].COL5))
			}
			if _, err = agent.QueryRow("Query", map[string]interface{}{"COL1": 5}, &resData); err != nil {
				t.Error(err)
			}
			assert.Equal(t, val.COL1, resData.COL1, fmt.Sprintf("%v != %v", val.COL1, resData.COL1))
			assert.Equal(t, val.COL2, resData.COL2, fmt.Sprintf("%v != %v", val.COL2, resData.COL2))
			assert.Equal(t, val.COL3, resData.COL3, fmt.Sprintf("%v != %v", val.COL3, resData.COL3))
			assert.Equal(t, val.COL4, resData.COL4, fmt.Sprintf("%v != %v", val.COL4, resData.COL4))
			assert.Equal(t, val.COL5, resData.COL5, fmt.Sprintf("%v != %v", val.COL5, resData.COL5))
			table := &TableAgent{
				Agent:  agent,
				Table:  "TEST",
				Params: []*Param{{Col: "COL1", Val: 5}},
			}
			if _, err = table.Query(&res); err != nil {
				t.Error(err)
			} else if len(res.Rows) != 1 {
				assert.Equal(t, val.COL1, res.Rows[0].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[0].COL1))
				assert.Equal(t, val.COL2, res.Rows[0].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[0].COL2))
				assert.Equal(t, val.COL3, res.Rows[0].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[0].COL3))
				assert.Equal(t, val.COL4, res.Rows[0].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[0].COL4))
				assert.Equal(t, val.COL5, res.Rows[0].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[0].COL5))
			}
			if _, err = table.QueryRow(&resData); err != nil {
				t.Error(err)
			} else {
				assert.Equal(t, val.COL1, resData.COL1, fmt.Sprintf("%v != %v", val.COL1, resData.COL1))
				assert.Equal(t, val.COL2, resData.COL2, fmt.Sprintf("%v != %v", val.COL2, resData.COL2))
				assert.Equal(t, val.COL3, resData.COL3, fmt.Sprintf("%v != %v", val.COL3, resData.COL3))
				assert.Equal(t, val.COL4, resData.COL4, fmt.Sprintf("%v != %v", val.COL4, resData.COL4))
				assert.Equal(t, val.COL5, resData.COL5, fmt.Sprintf("%v != %v", val.COL5, resData.COL5))
			}
		}
	}
}

func testQueryTx(t *testing.T) {
	var it time.Time
	var err error
	if it, err = jcast.Time("2022-11-21 21:25:11"); err != nil {
		return
	}
	type TestData struct {
		COL1 int64
		COL2 float64
		COL3 time.Time
		COL4 string
		COL5 []byte
	}
	type TestRes struct {
		Rows []TestData
	}
	val := TestData{
		COL1: 5,
		COL2: 0.5,
		COL3: it.In(time.UTC),
		COL4: "test 5",
		COL5: []byte("test 5"),
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				var res TestRes
				var resData TestData
				if _, e := agent.QueryTx("Query", map[string]interface{}{"COL1": 5}, &res); e != nil {
					return e
				}
				if len(res.Rows) != 1 {
					assert.Equal(t, val.COL1, res.Rows[0].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[0].COL1))
					assert.Equal(t, val.COL2, res.Rows[0].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[0].COL2))
					assert.Equal(t, val.COL3, res.Rows[0].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[0].COL3))
					assert.Equal(t, val.COL4, res.Rows[0].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[0].COL4))
					assert.Equal(t, val.COL5, res.Rows[0].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[0].COL5))
				}
				if _, e := agent.QueryRowTx("Query", map[string]interface{}{"COL1": 5}, &resData); e != nil {
					return e
				}
				assert.Equal(t, val.COL1, resData.COL1, fmt.Sprintf("%v != %v", val.COL1, resData.COL1))
				assert.Equal(t, val.COL2, resData.COL2, fmt.Sprintf("%v != %v", val.COL2, resData.COL2))
				assert.Equal(t, val.COL3, resData.COL3, fmt.Sprintf("%v != %v", val.COL3, resData.COL3))
				assert.Equal(t, val.COL4, resData.COL4, fmt.Sprintf("%v != %v", val.COL4, resData.COL4))
				assert.Equal(t, val.COL5, resData.COL5, fmt.Sprintf("%v != %v", val.COL5, resData.COL5))
				table := &TableAgent{
					Agent:  agent,
					Table:  "TEST",
					Params: []*Param{{Col: "COL1", Val: 5}},
				}
				if _, e := table.QueryTx(&res); e != nil {
					return e
				} else if len(res.Rows) != 1 {
					assert.Equal(t, val.COL1, res.Rows[0].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[0].COL1))
					assert.Equal(t, val.COL2, res.Rows[0].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[0].COL2))
					assert.Equal(t, val.COL3, res.Rows[0].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[0].COL3))
					assert.Equal(t, val.COL4, res.Rows[0].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[0].COL4))
					assert.Equal(t, val.COL5, res.Rows[0].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[0].COL5))
				}
				if _, e := table.QueryRowTx(&resData); e != nil {
					return e
				} else {
					assert.Equal(t, val.COL1, resData.COL1, fmt.Sprintf("%v != %v", val.COL1, resData.COL1))
					assert.Equal(t, val.COL2, resData.COL2, fmt.Sprintf("%v != %v", val.COL2, resData.COL2))
					assert.Equal(t, val.COL3, resData.COL3, fmt.Sprintf("%v != %v", val.COL3, resData.COL3))
					assert.Equal(t, val.COL4, resData.COL4, fmt.Sprintf("%v != %v", val.COL4, resData.COL4))
					assert.Equal(t, val.COL5, resData.COL5, fmt.Sprintf("%v != %v", val.COL5, resData.COL5))
				}
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testQueryPage(t *testing.T) {
	var it time.Time
	var err error
	if it, err = jcast.Time("2022-11-21 21:25:11"); err != nil {
		return
	}
	type TestData struct {
		COL1 int64
		COL2 float64
		COL3 time.Time
		COL4 string
		COL5 []byte
	}
	type TestRes struct {
		Rows []TestData
	}
	list := []TestData{
		{
			COL1: 5,
			COL2: 0.5,
			COL3: it.In(time.UTC),
			COL4: "test 5",
			COL5: []byte("test 5"),
		},
		{
			COL1: 6,
			COL2: 0.6,
			COL3: it.In(time.UTC),
			COL4: "test 6",
			COL5: []byte("test 6"),
		},
		{
			COL1: 8,
			COL2: 0.8,
			COL3: it.In(time.UTC),
			COL4: "test 8",
			COL5: []byte("test 8"),
		},
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			var res TestRes
			if _, err = agent.QueryPage("QueryPage", 5, 7, &res); err != nil {
				t.Error(err)
			}
			for i, val := range list {
				assert.Equal(t, val.COL1, res.Rows[i].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[i].COL1))
				assert.Equal(t, val.COL2, res.Rows[i].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[i].COL2))
				assert.Equal(t, val.COL3, res.Rows[i].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[i].COL3))
				assert.Equal(t, val.COL4, res.Rows[i].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[i].COL4))
				assert.Equal(t, val.COL5, res.Rows[i].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[i].COL5))
			}
			table := &TableAgent{
				Agent:  agent,
				Table:  "TEST",
				OrdStr: "COL1",
			}
			if _, err = table.QueryPage(5, 7, &res); err != nil {
				t.Error(err)
			} else {
				for i, val := range list {
					assert.Equal(t, val.COL1, res.Rows[i].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[i].COL1))
					assert.Equal(t, val.COL2, res.Rows[i].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[i].COL2))
					assert.Equal(t, val.COL3, res.Rows[i].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[i].COL3))
					assert.Equal(t, val.COL4, res.Rows[i].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[i].COL4))
					assert.Equal(t, val.COL5, res.Rows[i].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[i].COL5))
				}
			}
		}
	}
}

func testQueryPageTx(t *testing.T) {
	var it time.Time
	var err error
	if it, err = jcast.Time("2022-11-21 21:25:11"); err != nil {
		return
	}
	type TestData struct {
		COL1 int64
		COL2 float64
		COL3 time.Time
		COL4 string
		COL5 []byte
	}
	type TestRes struct {
		Rows []TestData
	}
	list := []TestData{
		{
			COL1: 5,
			COL2: 0.5,
			COL3: it.In(time.UTC),
			COL4: "test 5",
			COL5: []byte("test 5"),
		},
		{
			COL1: 6,
			COL2: 0.6,
			COL3: it.In(time.UTC),
			COL4: "test 6",
			COL5: []byte("test 6"),
		},
		{
			COL1: 8,
			COL2: 0.8,
			COL3: it.In(time.UTC),
			COL4: "test 8",
			COL5: []byte("test 8"),
		},
	}
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				var res TestRes
				if _, e := agent.QueryPageTx("QueryPage", 5, 7, &res); e != nil {
					return e
				}
				for i, val := range list {
					assert.Equal(t, val.COL1, res.Rows[i].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[i].COL1))
					assert.Equal(t, val.COL2, res.Rows[i].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[i].COL2))
					assert.Equal(t, val.COL3, res.Rows[i].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[i].COL3))
					assert.Equal(t, val.COL4, res.Rows[i].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[i].COL4))
					assert.Equal(t, val.COL5, res.Rows[i].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[i].COL5))
				}
				table := &TableAgent{
					Agent:  agent,
					Table:  "TEST",
					OrdStr: "COL1",
				}
				if _, e := table.QueryPageTx(5, 7, &res); e != nil {
					return e
				} else {
					for i, val := range list {
						assert.Equal(t, val.COL1, res.Rows[i].COL1, fmt.Sprintf("%v != %v", val.COL1, res.Rows[i].COL1))
						assert.Equal(t, val.COL2, res.Rows[i].COL2, fmt.Sprintf("%v != %v", val.COL2, res.Rows[i].COL2))
						assert.Equal(t, val.COL3, res.Rows[i].COL3, fmt.Sprintf("%v != %v", val.COL3, res.Rows[i].COL3))
						assert.Equal(t, val.COL4, res.Rows[i].COL4, fmt.Sprintf("%v != %v", val.COL4, res.Rows[i].COL4))
						assert.Equal(t, val.COL5, res.Rows[i].COL5, fmt.Sprintf("%v != %v", val.COL5, res.Rows[i].COL5))
					}
				}
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testCount(t *testing.T) {
	var err error
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			var count int
			if count, err = agent.Count("Count"); err != nil {
				t.Error(err)
			}
			assert.Equal(t, 15, count, fmt.Sprintf("%v != %v", 15, count))
			table := &TableAgent{
				Agent: agent,
				Table: "TEST",
			}
			if count, err = table.Count(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, 15, count, fmt.Sprintf("%v != %v", 15, count))
		}
	}
}

func testCountTx(t *testing.T) {
	var err error
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				var count int
				var e error
				if count, e = agent.CountTx("Count"); e != nil {
					return e
				}
				assert.Equal(t, 15, count, fmt.Sprintf("%v != %v", 15, count))
				table := &TableAgent{
					Agent: agent,
					Table: "TEST",
				}
				if count, e = table.CountTx(); e != nil {
					return e
				}
				assert.Equal(t, 15, count, fmt.Sprintf("%v != %v", 15, count))
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testExists(t *testing.T) {
	var err error
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			var exists bool
			if exists, err = agent.Exists("Query", map[string]interface{}{"COL1": 1}); err != nil {
				t.Error(err)
			}
			assert.Equal(t, true, exists, fmt.Sprintf("%v != %v", true, exists))
			if exists, err = agent.Exists("Query", map[string]interface{}{"COL1": 7}); err != nil {
				t.Error(err)
			}
			assert.Equal(t, false, exists, fmt.Sprintf("%v != %v", false, exists))
			table := &TableAgent{
				Agent:  agent,
				Table:  "TEST",
				Params: []*Param{{Col: "COL1", Val: 1}},
			}
			if exists, err = table.Exists(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, true, exists, fmt.Sprintf("%v != %v", true, exists))
			table = &TableAgent{
				Agent:  agent,
				Table:  "TEST",
				Params: []*Param{{Col: "COL1", Val: 7}},
			}
			if exists, err = table.Exists(); err != nil {
				t.Error(err)
			}
			assert.Equal(t, false, exists, fmt.Sprintf("%v != %v", false, exists))
		}
	}
}

func testExistsTx(t *testing.T) {
	var err error
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		var agent *Agent
		if agent, err = GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				var exists bool
				var e error
				if exists, e = agent.ExistsTx("Query", map[string]interface{}{"COL1": 1}); e != nil {
					return e
				}
				assert.Equal(t, true, exists, fmt.Sprintf("%v != %v", true, exists))
				if exists, e = agent.ExistsTx("Query", map[string]interface{}{"COL1": 7}); e != nil {
					return e
				}
				assert.Equal(t, false, exists, fmt.Sprintf("%v != %v", false, exists))
				table := &TableAgent{
					Agent:  agent,
					Table:  "TEST",
					Params: []*Param{{Col: "COL1", Val: 1}},
				}
				if exists, e = table.ExistsTx(); e != nil {
					return e
				}
				assert.Equal(t, true, exists, fmt.Sprintf("%v != %v", true, exists))
				table = &TableAgent{
					Agent:  agent,
					Table:  "TEST",
					Params: []*Param{{Col: "COL1", Val: 7}},
				}
				if exists, e = table.ExistsTx(); e != nil {
					return e
				}
				assert.Equal(t, false, exists, fmt.Sprintf("%v != %v", false, exists))
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testTableSchema(t *testing.T) {
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	tables := []string{"TEST"}
	schemas := map[string][]TableSchema{
		"testMySql": {
			{ColumnName: "COL1", DataType: "int", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "int type", TableComment: "test table"},
			{ColumnName: "COL2", DataType: "decimal", IsNullable: "NO", DataDefault: "7.70", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "decimal(10, 2) type", TableComment: "test table"},
			{ColumnName: "COL3", DataType: "datetime", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "datetime type", TableComment: "test table"},
			{ColumnName: "COL4", DataType: "varchar", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar(10) type", TableComment: "test table"},
			{ColumnName: "COL5", DataType: "blob", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "blob type", TableComment: "test table"},
		},
		"testMSSql": {
			{ColumnName: "COL1", DataType: "int", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "int type", TableComment: "test table"},
			{ColumnName: "COL2", DataType: "decimal", IsNullable: "NO", DataDefault: "((7.70))", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "decimal(10, 2) type", TableComment: "test table"},
			{ColumnName: "COL3", DataType: "datetime", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "datetime type", TableComment: "test table"},
			{ColumnName: "COL4", DataType: "varchar", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar(10) type", TableComment: "test table"},
			{ColumnName: "COL5", DataType: "varbinary", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varbinary(max) type", TableComment: "test table"},
		},
		"testOracle": {
			{ColumnName: "COL1", DataType: "NUMBER", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "NUMBER type", TableComment: "test table"},
			{ColumnName: "COL2", DataType: "NUMBER", IsNullable: "NO", DataDefault: "7.70 ", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "NUMBER(10, 2) type", TableComment: "test table"},
			{ColumnName: "COL3", DataType: "DATE", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "date type", TableComment: "test table"},
			{ColumnName: "COL4", DataType: "VARCHAR2", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar2(10) type", TableComment: "test table"},
			{ColumnName: "COL5", DataType: "BLOB", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "blob type", TableComment: "test table"},
		},
		"testPostgreSql": {
			{ColumnName: "col1", DataType: "integer", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "integer type", TableComment: "test table"},
			{ColumnName: "col2", DataType: "numeric", IsNullable: "NO", DataDefault: "7.70", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "numeric(10, 2) type", TableComment: "test table"},
			{ColumnName: "col3", DataType: "timestamp without time zone", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "timestamp without time zone type", TableComment: "test table"},
			{ColumnName: "col4", DataType: "character varying", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar(10) type", TableComment: "test table"},
			{ColumnName: "col5", DataType: "bytea", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "bytea type", TableComment: "test table"},
		},
	}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			var list []string
			if list, err = agent.Tables(); err != nil {
				t.Error(err)
			} else if len(list) != 1 {
				t.Error(fmt.Errorf("query table error"))
			} else {
				for i1, tb := range tables {
					assert.Equal(t, strings.ToUpper(tb), strings.ToUpper(list[i1]), fmt.Sprintf("%v != %v", strings.ToUpper(tb), strings.ToUpper(list[i1])))
					var sa []TableSchema
					if sa, err = agent.TableSchema(list[i1]); err != nil {
						t.Error(err)
					} else if len(sa) != 5 {
						t.Error(fmt.Errorf("query table schema error"))
					} else {
						for i2, schema := range schemas[v] {
							assert.Equal(t, schema.ColumnName, sa[i2].ColumnName, fmt.Sprintf("%v != %v", schema.ColumnName, sa[i2].ColumnName))
							assert.Equal(t, schema.DataType, sa[i2].DataType, fmt.Sprintf("%v != %v", schema.DataType, sa[i2].DataType))
							assert.Equal(t, schema.IsNullable, sa[i2].IsNullable, fmt.Sprintf("%v != %v", schema.IsNullable, sa[i2].IsNullable))
							assert.Equal(t, schema.DataDefault, sa[i2].DataDefault, fmt.Sprintf("%v != %v", schema.DataDefault, sa[i2].DataDefault))
							assert.Equal(t, schema.PrimaryKey, sa[i2].PrimaryKey, fmt.Sprintf("%v != %v", schema.PrimaryKey, sa[i2].PrimaryKey))
							assert.Equal(t, schema.IsIdentity, sa[i2].IsIdentity, fmt.Sprintf("%v != %v", schema.IsIdentity, sa[i2].IsIdentity))
							assert.Equal(t, schema.ColumnComment, sa[i2].ColumnComment, fmt.Sprintf("%v != %v", schema.ColumnComment, sa[i2].ColumnComment))
							assert.Equal(t, schema.TableComment, sa[i2].TableComment, fmt.Sprintf("%v != %v", schema.TableComment, sa[i2].TableComment))
						}
					}
				}
			}
		}
	}
}

func testTableSchemaTx(t *testing.T) {
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	tables := []string{"TEST"}
	schemas := map[string][]TableSchema{
		"testMySql": {
			{ColumnName: "COL1", DataType: "int", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "int type", TableComment: "test table"},
			{ColumnName: "COL2", DataType: "decimal", IsNullable: "NO", DataDefault: "7.70", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "decimal(10, 2) type", TableComment: "test table"},
			{ColumnName: "COL3", DataType: "datetime", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "datetime type", TableComment: "test table"},
			{ColumnName: "COL4", DataType: "varchar", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar(10) type", TableComment: "test table"},
			{ColumnName: "COL5", DataType: "blob", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "blob type", TableComment: "test table"},
		},
		"testMSSql": {
			{ColumnName: "COL1", DataType: "int", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "int type", TableComment: "test table"},
			{ColumnName: "COL2", DataType: "decimal", IsNullable: "NO", DataDefault: "((7.70))", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "decimal(10, 2) type", TableComment: "test table"},
			{ColumnName: "COL3", DataType: "datetime", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "datetime type", TableComment: "test table"},
			{ColumnName: "COL4", DataType: "varchar", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar(10) type", TableComment: "test table"},
			{ColumnName: "COL5", DataType: "varbinary", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varbinary(max) type", TableComment: "test table"},
		},
		"testOracle": {
			{ColumnName: "COL1", DataType: "NUMBER", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "NUMBER type", TableComment: "test table"},
			{ColumnName: "COL2", DataType: "NUMBER", IsNullable: "NO", DataDefault: "7.70 ", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "NUMBER(10, 2) type", TableComment: "test table"},
			{ColumnName: "COL3", DataType: "DATE", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "date type", TableComment: "test table"},
			{ColumnName: "COL4", DataType: "VARCHAR2", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar2(10) type", TableComment: "test table"},
			{ColumnName: "COL5", DataType: "BLOB", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "blob type", TableComment: "test table"},
		},
		"testPostgreSql": {
			{ColumnName: "col1", DataType: "integer", IsNullable: "NO", DataDefault: "", PrimaryKey: float64(1), IsIdentity: "NO", ColumnComment: "integer type", TableComment: "test table"},
			{ColumnName: "col2", DataType: "numeric", IsNullable: "NO", DataDefault: "7.70", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "numeric(10, 2) type", TableComment: "test table"},
			{ColumnName: "col3", DataType: "timestamp without time zone", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "timestamp without time zone type", TableComment: "test table"},
			{ColumnName: "col4", DataType: "character varying", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "varchar(10) type", TableComment: "test table"},
			{ColumnName: "col5", DataType: "bytea", IsNullable: "YES", DataDefault: "", PrimaryKey: nil, IsIdentity: "NO", ColumnComment: "bytea type", TableComment: "test table"},
		},
	}
	for _, v := range testData {
		if agent, err := GetAgent(v); err != nil {
			t.Error(err)
		} else {
			if err = agent.UseTx(func() error {
				if list, e := agent.TablesTx(); e != nil {
					return e
				} else if len(list) != 1 {
					return fmt.Errorf("query table error")
				} else {
					for i1, tb := range tables {
						assert.Equal(t, strings.ToUpper(tb), strings.ToUpper(list[i1]), fmt.Sprintf("%v != %v", strings.ToUpper(tb), strings.ToUpper(list[i1])))
						var sa []TableSchema
						if sa, e = agent.TableSchemaTx(list[i1]); e != nil {
							return e
						} else if len(sa) != 5 {
							return fmt.Errorf("query table schema error")
						} else {
							for i2, schema := range schemas[v] {
								assert.Equal(t, schema.ColumnName, sa[i2].ColumnName, fmt.Sprintf("%v != %v", schema.ColumnName, sa[i2].ColumnName))
								assert.Equal(t, schema.DataType, sa[i2].DataType, fmt.Sprintf("%v != %v", schema.DataType, sa[i2].DataType))
								assert.Equal(t, schema.IsNullable, sa[i2].IsNullable, fmt.Sprintf("%v != %v", schema.IsNullable, sa[i2].IsNullable))
								assert.Equal(t, schema.DataDefault, sa[i2].DataDefault, fmt.Sprintf("%v != %v", schema.DataDefault, sa[i2].DataDefault))
								assert.Equal(t, schema.PrimaryKey, sa[i2].PrimaryKey, fmt.Sprintf("%v != %v", schema.PrimaryKey, sa[i2].PrimaryKey))
								assert.Equal(t, schema.IsIdentity, sa[i2].IsIdentity, fmt.Sprintf("%v != %v", schema.IsIdentity, sa[i2].IsIdentity))
								assert.Equal(t, schema.ColumnComment, sa[i2].ColumnComment, fmt.Sprintf("%v != %v", schema.ColumnComment, sa[i2].ColumnComment))
								assert.Equal(t, schema.TableComment, sa[i2].TableComment, fmt.Sprintf("%v != %v", schema.TableComment, sa[i2].TableComment))
							}
						}
					}
				}
				return nil
			}); err != nil {
				t.Error(err)
			}
		}
	}
}

func testDropTable(t *testing.T) {
	testData := []string{"testMySql", "testMSSql", "testOracle", "testPostgreSql"}
	for _, v := range testData {
		agent := &TableAgent{DSKey: v, Table: "TEST"}
		if _, err := agent.Drop(); err != nil {
			t.Error(err)
		}
	}
}
