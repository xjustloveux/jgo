// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"encoding/xml"
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jevent"
	"github.com/xjustloveux/jgo/jfile"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

const (
	errorNoRowsAvailable   = jError("no rows available")
	errorRowsNil           = jError("rows is nil")
	errorTableEmpty        = jError("table name is empty")
	errorAgentNil          = jError("agent is nil")
	errorNotFountDaoFolder = jError("not found dao folder: %q")

	errorColTypeNotStringType = jError("column name type is %q, not string")
	errorColNil               = jError("column is nil")

	errorNotValidDbType    = jError("not a valid db Type %q")
	errorNotValidOperators = jError("not a valid Operators %q")

	errorUnknownDataSource            = jError("unknown data source %q")
	errorUnknownSelectId              = jError("unknown select id %q")
	errorUnknownInsertId              = jError("unknown insert id %q")
	errorUnknownUpdateId              = jError("unknown update id %q")
	errorUnknownDeleteId              = jError("unknown delete id %q")
	errorUnknownOtherId               = jError("unknown other id %q")
	errorUnknownOps                   = jError("unknown Operations")
	errorUnknownOpr                   = jError("unknown Operators")
	errorUnknownSqlTypeForAgentTables = jError("unknown sql type, you can use args input query statement")

	errorDbAlreadyOpen = jError("db has already been open")
	errorDbNotOpen     = jError("db has not been opened")
	errorDbNotBegin    = jError("db has not been begin")
	errorDBNil         = jError("db is nil")

	errorDecodeFuncOut2NotErrorType  = jError("decode function second output type not error type")
	errorDecodeFuncOut1NotStringType = jError("decode function first output type not string type")
	errorDecodeFuncType              = jError("decode function input params must be (string), output params must be (string, error)")

	errorWrongTypeOfForeach = jError("wrong params type of tags <foreach>, type must be []string or map[string]string")
	errorWrongSql           = jError("wrong %q sql statements")

	errorOprValLenZero               = jError("operators %q, the value length is zero")
	errorOprValLenNot2               = jError("operators %q, the value length not 2")
	errorOprValTypeNotInterfaceSlice = jError("operators %q, the value type is %q, not []interface{}")
	errorOprValTypeNotString         = jError("operators %q, the value type is %q, not string")
)

const (
	pkgName       = "jsql"
	fileName      = "config.json"
	totalRecord   = "TOTALRECORD"
	allowPagingId = "ALLOWPAGINGID"
	orderById     = "ORDERBYID"
	Unknown       = -1
)

var (
	conf       = jconf.New()
	subject    = jevent.New()
	mux        = new(sync.RWMutex)
	data       *configData
	pack       *configPack
	decodeFunc func(string) (string, error)
	dsMap      map[string]*dataSource
	selectMap  map[string]*element
	insertMap  map[string]*element
	updateMap  map[string]*element
	deleteMap  map[string]*element
	otherMap   map[string]*element
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

// SetRoot set config root path
func SetRoot(root string) {
	conf.SetRoot(root)
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

// SubscribeSql subscribe sql event
func SubscribeSql(e jevent.Event) jevent.Subscription {
	return subject.Subscribe(e)
}

// Init initialize
func Init() error {
	if err := conf.Load(); err != nil {
		return err
	}
	data = &configData{}
	if err := conf.Convert(data); err != nil {
		return err
	}
	pack = data.Db
	if err := createDataSource(); err != nil {
		return err
	}
	return loadDaoXml()
}

// GetDaoPath returns conf.DaoPath
func GetDaoPath() string {
	if pack == nil {
		return ""
	}
	return pack.DaoPath
}

// GetDefaultDataSource returns json default data source
func GetDefaultDataSource() string {
	if pack == nil {
		return ""
	}
	return pack.Default
}

// GetDataSource returns json data source
func GetDataSource() map[string]interface{} {
	if pack == nil {
		return nil
	}
	return pack.DataSource
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
		return nil, errorFmt(errorUnknownDataSource, key)
	}
	t, _ := ParseDBType(ds.Type)
	if ds.db == nil {
		if err := ds.open(); err != nil {
			return nil, err
		}
	}
	return &Agent{db: ds.db, t: t, dbName: ds.DbName}, nil
}

func errorFmt(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errorStr(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
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
	for dk, dv := range pack.DataSource {
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
		selectMap = make(map[string]*element)
		insertMap = make(map[string]*element)
		updateMap = make(map[string]*element)
		deleteMap = make(map[string]*element)
		otherMap = make(map[string]*element)
		for _, dao := range list {
			if dao != nil {
				for _, elem := range dao.nodes {
					switch elem.tag {
					case tagSelect:
						selectMap[elem.id] = elem
					case tagInsert:
						insertMap[elem.id] = elem
					case tagUpdate:
						updateMap[elem.id] = elem
					case tagDelete:
						deleteMap[elem.id] = elem
					case tagOther:
						otherMap[elem.id] = elem
					}
				}
			}
		}
	}
	return nil
}

func loadDaoXmlDir(path string) (xmlList []*element, err error) {
	path = strings.Trim(strings.Trim(path, "\\ "), "/ ")
	exist := false
	if exist, err = jfile.Exist(path); err != nil {
		return nil, err
	} else if !exist {
		return nil, errorFmt(errorNotFountDaoFolder, path)
	}
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
		xmlList = make([]*element, 0)
		for _, f := range list {
			if f.IsDir() {
				var xs []*element
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
					var dao *element
					if dao, err = toElement(fmt.Sprint(path, "/", f.Name())); err != nil {
						return nil, err
					} else {
						xmlList = append(xmlList, dao)
					}
				}
			}
		}
		return xmlList, nil
	}
}

func getElement(ops Operations, id string) (*element, error) {
	mux.RLock()
	defer func() {
		mux.RUnlock()
	}()
	switch ops {
	case Select:
		if xs := selectMap[id]; xs == nil {
			return nil, errorFmt(errorUnknownSelectId, id)
		} else {
			return xs, nil
		}
	case Insert:
		if xi := insertMap[id]; xi == nil {
			return nil, errorFmt(errorUnknownInsertId, id)
		} else {
			return xi, nil
		}
	case Update:
		if xu := updateMap[id]; xu == nil {
			return nil, errorFmt(errorUnknownUpdateId, id)
		} else {
			return xu, nil
		}
	case Delete:
		if xd := deleteMap[id]; xd == nil {
			return nil, errorFmt(errorUnknownDeleteId, id)
		} else {
			return xd, nil
		}
	case Other:
		if xo := otherMap[id]; xo == nil {
			return nil, errorFmt(errorUnknownOtherId, id)
		} else {
			return xo, nil
		}
	}
	return nil, errorStr(errorUnknownOps)
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

func removeComment(str string) string {
	for {
		st := "<!--"
		et := "-->"
		si := strings.Index(str, st)
		ei := strings.Index(str, et)
		if si >= ei || si < 0 || ei < 0 {
			break
		}
		str = fmt.Sprint(str[:si], str[ei+len(et):])
	}
	for {
		st := "--"
		et := "\r\n"
		si := strings.Index(str, st)
		if si >= 0 {
			ei := strings.Index(str[si:], et)
			if ei < 0 {
				et = "\n"
				ei = strings.Index(str[si:], et)
			}
			if ei < 0 {
				et = "\r"
				ei = strings.Index(str[si:], et)
			}
			if ei < 0 {
				str = str[:si]
			} else {
				str = fmt.Sprint(str[:si], str[si+ei+len(et):])
			}
		} else {
			break
		}
	}
	return str
}

func toElement(path string) (dao *element, err error) {
	var file *os.File
	if file, err = os.Open(path); err != nil {
		return nil, err
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
		}
	}()
	parser := xml.NewDecoder(file)
	idx := make([]int, 0)
	for {
		var token xml.Token
		if token, err = parser.Token(); err != nil {
			err = nil
			break
		}
		switch t := token.(type) {
		case xml.StartElement:
			name := t.Name.Local
			switch tn := parseTag(name); tn {
			case tagDao:
				if dao == nil {
					dao = &element{id: "", tag: tn, nodes: make([]*element, 0)}
				}
			case tagSelect:
				fallthrough
			case tagInsert:
				fallthrough
			case tagUpdate:
				fallthrough
			case tagDelete:
				fallthrough
			case tagOther:
				if dao != nil && len(idx) <= 0 {
					attr := make(map[string]string)
					for _, a := range t.Attr {
						attr[a.Name.Local] = a.Value
					}
					dao.nodes = append(dao.nodes, &element{id: attr["id"], tag: tn, attr: attr, text: "", nodes: make([]*element, 0)})
					idx = append(idx, len(dao.nodes)-1)
				}
			case tagIf:
				fallthrough
			case tagForeach:
				fallthrough
			case tagWhere:
				fallthrough
			case tagOrderBy:
				if dao != nil && len(idx) > 0 {
					e := dao
					for _, i := range idx {
						e = e.nodes[i]
					}
					attr := make(map[string]string)
					for _, a := range t.Attr {
						attr[strings.ToLower(a.Name.Local)] = a.Value
					}
					e.nodes = append(e.nodes, &element{id: "", tag: tn, attr: attr, text: "", nodes: make([]*element, 0)})
					idx = append(idx, len(e.nodes)-1)
				}
			}
		case xml.EndElement:
			if l := len(idx); l > 0 {
				idx = idx[:l-1]
			}
		case xml.CharData:
			if dao != nil && len(idx) > 0 {
				e := dao
				for _, i := range idx {
					e = e.nodes[i]
				}
				text := trim(removeComment(jcast.String(t)))
				e.nodes = append(e.nodes, &element{id: "", tag: tagText, attr: nil, text: text, nodes: nil})
			}
		case xml.Comment:
		case xml.ProcInst:
		case xml.Directive:
		default:
		}
	}
	return dao, nil
}
