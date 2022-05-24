// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"encoding/xml"
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jfile"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	errorNoRowsAvailable = jError("no rows available")
	errorRowsNil         = jError("rows is nil")
	errorTableEmpty      = jError("table name is empty")
	errorAgentNil        = jError("agent is nil")

	errorColTypeNotStringType = jError("column name type is %q, not string")
	errorColNil               = jError("column is nil")

	errorNotValidDbType    = jError("not a valid db Type %q")
	errorNotValidOperators = jError("not a valid Operators %q")

	errorUnknownDataSource = jError("unknown data source %q")
	errorUnknownSelectId   = jError("unknown select id %q")
	errorUnknownInsertId   = jError("unknown insert id %q")
	errorUnknownUpdateId   = jError("unknown update id %q")
	errorUnknownDeleteId   = jError("unknown delete id %q")
	errorUnknownOtherId    = jError("unknown other id %q")
	errorUnknownOps        = jError("unknown Operations")
	errorUnknownOpr        = jError("unknown Operators")

	errorDbAlreadyOpen = jError("db has already been open")
	errorDbNotOpen     = jError("db has not been opened")
	errorDbNotBegin    = jError("db has not been begin")
	errorDBNil         = jError("db is nil")

	errorDecodeFuncOut2NotErrorType  = jError("decode function second output type not error type")
	errorDecodeFuncOut1NotStringType = jError("decode function first output type not string type")
	errorDecodeFuncType              = jError("decode function input params must be (string), output params must be (string, error)")

	errorXmlNotSelectType = jError("xml type is %q, not *xmlSelect")
	errorXmlNotInsertType = jError("xml type is %q, not *xmlInsert")
	errorXmlNotUpdateType = jError("xml type is %q, not *xmlUpdate")
	errorXmlNotDeleteType = jError("xml type is %q, not *xmlDelete")
	errorXmlNotOtherType  = jError("xml type is %q, not *xmlOther")

	errorWrongTypeOfForeach = jError("wrong params type of tags <foreach>, type must be []string or map[string]string")
	errorWrongNumOfForeach  = jError("wrong number of tags <foreach>")
	errorWrongNumOfIf       = jError("wrong number of tags <if>")
	errorWrongNumOfOrderBy  = jError("wrong number of tags <orderBy>")
	errorWrongSql           = jError("wrong %q sql statements")

	errorOprValLenZero               = jError("operators %q, the value length is zero")
	errorOprValLenNot2               = jError("operators %q, the value length not 2")
	errorOprValTypeNotInterfaceSlice = jError("operators %q, the value type is %q, not []interface{}")
	errorOprValTypeNotString         = jError("operators %q, the value type is %q, not string")
)

const (
	pkgName       = "jsql"
	fileName      = "sql.json"
	totalRecord   = "TOTALRECORD"
	allowPagingId = "ALLOWPAGINGID"
	orderById     = "ORDERBYID"
	Unknown       = -1
)

var (
	conf       = jconf.New()
	mux        = new(sync.RWMutex)
	cd         *configData
	logFunc    func(...interface{})
	decodeFunc func(string) (string, error)
	dsMap      map[string]*dataSource
	selectMap  map[string]*xmlSelect
	insertMap  map[string]*xmlInsert
	updateMap  map[string]*xmlUpdate
	deleteMap  map[string]*xmlDelete
	otherMap   map[string]*xmlOther
)

func init() {
	SetFileName(fileName)
}

// SetFormat set config format
func SetFormat(f jfile.Format) {
	conf.SetFormat(f)
}

// SetFileName set config file name
func SetFileName(name string) {
	conf.SetFileName(name)
}

// SetEnvFileName set config env file name
func SetEnvFileName(name string) {
	conf.SetEnvFileName(name)
}

// EnvKey returns env key
func EnvKey() string {
	return conf.EnvKey()
}

// SetEnvKey set env key
func SetEnvKey(key string) {
	conf.SetEnvKey(key)
}

// EnvVal returns env value
func EnvVal() string {
	return conf.EnvVal()
}

// SetEnvVal set env value
func SetEnvVal(val string) {
	conf.SetEnvVal(val)
}

// EnableEnv enable env
func EnableEnv() {
	conf.EnableEnv()
}

// DisableEnv disable env
func DisableEnv() {
	conf.DisableEnv()
}

// SetLogFunc set fmt.Println log function
func SetLogFunc(f func(...interface{})) {
	logFunc = f
}

// Init initialize
func Init() error {
	if err := conf.Load(); err != nil {
		return err
	}
	cd = &configData{Debug: false}
	if err := conf.Convert(cd); err != nil {
		return err
	}
	if err := createDataSource(); err != nil {
		return err
	}
	return loadDaoXml()
}

// GetDaoPath returns conf.DaoPath
func GetDaoPath() string {
	return cd.DaoPath
}

// GetDefaultDataSource returns json default data source
func GetDefaultDataSource() string {
	return cd.Default
}

// SetDecodeFunc set decode data source name or data source json string function
func SetDecodeFunc(f func(string) (string, error)) {
	decodeFunc = f
}

// GetAgent returns SqlAgent
// if not input data source key then return default data source agent
func GetAgent(dsKey ...string) (*Agent, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	key := ""
	if len(dsKey) > 0 && dsKey[0] != "" {
		key = dsKey[0]
	} else {
		key = GetDefaultDataSource()
	}
	ds := dsMap[key]
	if ds == nil {
		return nil, errorf(errorUnknownDataSource, key)
	}
	t, _ := ParseDBType(ds.Type)
	if ds.db == nil {
		if err := ds.open(); err != nil {
			return nil, err
		}
	}
	return &Agent{db: ds.db, t: t}, nil
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}

func fmtPrintln(args ...interface{}) {
	if cd.Debug {
		fmt.Println(args...)
	}
	if logFunc != nil {
		logFunc(args...)
	}
}

func createDataSource() error {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	var err error
	for _, v := range dsMap {
		if v.db != nil {
			if err = v.close(); err != nil {
				return err
			}
		}
	}
	dsMap = make(map[string]*dataSource)
	for dk, dv := range cd.DataSource {
		var dm map[string]interface{}
		if dm, err = jcast.StringMapInterface(dv); err != nil {
			return err
		}
		dsMap[dk] = dataSource{}.getDefault()
		if err = jfile.Convert(dm, dsMap[dk]); err != nil {
			return err
		}
	}
	return nil
}

func loadDaoXml() error {
	if list, err := loadDaoXmlDir(GetDaoPath()); err != nil {
		return err
	} else {
		selectMap = make(map[string]*xmlSelect)
		insertMap = make(map[string]*xmlInsert)
		updateMap = make(map[string]*xmlUpdate)
		deleteMap = make(map[string]*xmlDelete)
		otherMap = make(map[string]*xmlOther)
		for _, xs := range list {
			for _, xss := range xs.Select {
				selectMap[xss.Id] = xss
			}
			for _, xsi := range xs.Insert {
				insertMap[xsi.Id] = xsi
			}
			for _, xsu := range xs.Update {
				updateMap[xsu.Id] = xsu
			}
			for _, xsd := range xs.Delete {
				deleteMap[xsd.Id] = xsd
			}
			for _, xso := range xs.Other {
				otherMap[xso.Id] = xso
			}
		}
	}
	return nil
}

func loadDaoXmlDir(path string) (xmlList []xmlSql, err error) {
	path = strings.Trim(strings.Trim(path, "\\ "), "/ ")
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return nil, err
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
		}
	}()
	var list []os.FileInfo
	if list, err = file.Readdir(0); err != nil {
		return nil, err
	} else {
		xmlList = make([]xmlSql, 0)
		for _, f := range list {
			if f.IsDir() {
				var xs []xmlSql
				if xs, err = loadDaoXmlDir(fmt.Sprint(path, "/", f.Name())); err != nil {
					return nil, err
				} else {
					xmlList = append(xmlList, xs...)
				}
			} else {
				var isXml bool
				if isXml, err = regexp.MatchString(".xml$", f.Name()); err != nil || !isXml {
					continue
				}
				if isXml {
					var b []byte
					if b, err = jfile.Load(fmt.Sprint(path, "/", f.Name())); err != nil {
						return nil, err
					} else {
						var result xmlSql
						if err = xml.Unmarshal(b, &result); err != nil {
							return nil, err
						}
						xmlList = append(xmlList, result)
					}
				}
			}
		}
		return xmlList, nil
	}
}

func xmlToSql(ops Operations, xml string, param map[string]interface{}, xi []xmlIf, xf []xmlFor, xo []xmlOrderBy, page bool) (query, order string, err error) {
	query = xml
	if param != nil {
		if param, err = jcast.StringMapInterface(param); err != nil {
			return "", "", err
		}
		for k, v := range param {
			if reflect.TypeOf(v).String() == "time.Time" {
				if param[k], err = jcast.TimeString(v); err != nil {
					return "", "", err
				}
			}
		}
	}
	if query, order, err = replaceXmlForeach(xi, xf, xo, xml, param, page); err != nil {
		return "", "", err
	}
	if param != nil {
		for k, v := range param {
			if reflect.TypeOf(v).Kind() == reflect.String {
				query = strings.ReplaceAll(query, fmt.Sprint("${", k, "}"), v.(string))
				order = strings.ReplaceAll(order, fmt.Sprint("${", k, "}"), v.(string))
			}
		}
	}
	query = trim(query)
	cs := ""
	switch ops {
	case Select:
		cs = "select"
	case Insert:
		cs = "insert"
	case Update:
		cs = "update"
	case Delete:
		cs = "delete"
	}
	if strings.Index(strings.ToLower(query), cs) != 0 {
		return "", "", errorf(errorWrongSql, cs)
	}
	return query, order, nil
}

func replaceXmlForeach(xi []xmlIf, xf []xmlFor, xo []xmlOrderBy, xml string, params map[string]interface{}, page bool) (sql, orderByStr string, err error) {
	sql = xml
	var obs string
	for _, i := range xi {
		if sql, obs, err = i.replaceXml(sql, params, page); err != nil {
			return "", "", err
		}
		if obs != "" {
			orderByStr = obs
		}
	}
	for _, f := range xf {
		if sql, obs, err = f.replaceXml(sql, params, page); err != nil {
			return "", "", err
		}
		if obs != "" {
			orderByStr = obs
		}
	}
	for _, o := range xo {
		if sql, obs, err = o.replaceXml(sql, params, page); err != nil {
			return "", "", err
		}
		if obs != "" {
			orderByStr = obs
		}
	}
	return sql, orderByStr, nil
}

func getXmlSql(ops Operations, id string) (interface{}, error) {
	mux.RLock()
	defer func() {
		mux.RUnlock()
	}()
	switch ops {
	case Select:
		if xs := selectMap[id]; xs == nil {
			return nil, errorf(errorUnknownSelectId, id)
		} else {
			return xs, nil
		}
	case Insert:
		if xi := insertMap[id]; xi == nil {
			return nil, errorf(errorUnknownInsertId, id)
		} else {
			return xi, nil
		}
	case Update:
		if xu := updateMap[id]; xu == nil {
			return nil, errorf(errorUnknownUpdateId, id)
		} else {
			return xu, nil
		}
	case Delete:
		if xd := deleteMap[id]; xd == nil {
			return nil, errorf(errorUnknownDeleteId, id)
		} else {
			return xd, nil
		}
	case Other:
		if xo := otherMap[id]; xo == nil {
			return nil, errorf(errorUnknownOtherId, id)
		} else {
			return xo, nil
		}
	}
	return nil, errors(errorUnknownOps)
}

func getPageSql(t Type, sql, obs string, start, end int64) (pageSql, countSql string) {
	if obs != "" {
		obs = fmt.Sprint(" ORDER BY ", obs)
	}
	switch t {
	case MySql:
		pageSql = fmt.Sprint(
			"SELECT * FROM (SELECT (@i := @i + 1) AS ", allowPagingId, ", table1.* ",
			"FROM (SELECT *, 1 as ", orderById, " FROM (", sql, obs,
			")  as  tbs1 ) as table1, (select @i := 0) temp ORDER BY ", orderById, " DESC ) as table2 ",
			"WHERE ", allowPagingId, " BETWEEN ", strconv.FormatInt(start, 10), " AND ", strconv.FormatInt(end, 10))
	case MSSql:
		obis := ""
		if obs == "" {
			obs = fmt.Sprint("ORDER BY ", orderById, " DESC")
			obis = fmt.Sprint(", 1 as ", orderById)
		}
		pageSql = fmt.Sprint(
			"SELECT * FROM (SELECT ROW_NUMBER() OVER(", obs, ") AS ", allowPagingId, ",* ",
			"FROM (SELECT *", obis, " FROM (", sql, ") as tbs1) as table1) as table2 ",
			"WHERE ", allowPagingId, " BETWEEN ", strconv.FormatInt(start, 10), " AND ", strconv.FormatInt(end, 10))
	case Oracle:
		pageSql = fmt.Sprint(
			"SELECT t3.* FROM (SELECT t2.*, rownum as ", allowPagingId, " ",
			"FROM (SELECT t1.*, 1 as ", orderById, " FROM (", sql, obs, ") t1) t2 ORDER BY ", orderById, ") t3 ",
			"WHERE ", allowPagingId, " BETWEEN ", strconv.FormatInt(start, 10), " AND ", strconv.FormatInt(end, 10))
	}
	countSql = fmt.Sprint("SELECT COUNT(1) as ", totalRecord, " FROM (", sql, ") data")
	return pageSql, countSql
}

func trim(str string) string {
	ts := []string{" ", "　", "\r\n", "\r", "\n"}
	for i := 0; i < len(ts); i++ {
		for _, v := range ts {
			str = strings.Trim(str, v)
		}
	}
	return str
}
