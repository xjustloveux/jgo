// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/xjustloveux/jgo/jcast"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type element struct {
	id    string
	tag   tag
	attr  map[string]string
	text  string
	nodes []*element
}

func (e *element) getSql(param map[string]interface{}, page bool) (string, string, error) {
	query := ""
	order := ""
	switch e.tag {
	case tagSelect:
		fallthrough
	case tagInsert:
		fallthrough
	case tagUpdate:
		fallthrough
	case tagDelete:
		fallthrough
	case tagOther:
		var err error
		if param != nil {
			if param, err = jcast.StringMapInterface(param); err != nil {
				return "", "", err
			}
			for k, v := range param {
				if v != nil {
					if reflect.TypeOf(v).String() == "time.Time" {
						if param[k], err = jcast.TimeString(v); err != nil {
							return "", "", err
						}
					}
				}
			}
		}
		if query, order, err = nodesToQuery(e.nodes, param, page); err != nil {
			return "", "", err
		}
		if param != nil {
			for k, v := range param {
				if v != nil {
					if reflect.TypeOf(v).Kind() == reflect.String {
						query = strings.ReplaceAll(query, fmt.Sprint("${", k, "}"), v.(string))
						order = strings.ReplaceAll(order, fmt.Sprint("${", k, "}"), v.(string))
					}
				}
			}
		}
		if e.tag != tagOther {
			if strings.Index(strings.ToLower(query), e.tag.String()) != 0 {
				return "", "", errorFmt(errorWrongSql, e.tag.String())
			}
		}
	case tagText:
		query = e.text
	case tagIf:
		sorTest := e.attr["test"]
		test := ""
		pattern := regexp.MustCompile(`nil\(\w+\)`)
		testIdx := 0
		for _, v := range pattern.FindAllStringSubmatchIndex(sorTest, -1) {
			st := "nil("
			et := ")"
			k := sorTest[v[0]+len(st) : v[1]-len(et)]
			test = fmt.Sprint(test, sorTest[testIdx:v[0]], strconv.FormatBool(param == nil || param[k] == nil))
			testIdx = v[1]
		}
		test = fmt.Sprint(test, sorTest[testIdx:])
		test = strings.ReplaceAll(test, " and ", " && ")
		test = strings.ReplaceAll(test, " or ", " || ")
		var err error
		var expression *govaluate.EvaluableExpression
		if expression, err = govaluate.NewEvaluableExpression(test); err != nil {
			return "", "", err
		}
		var res interface{}
		if res, err = expression.Evaluate(param); err != nil {
			return "", "", err
		}
		if reflect.TypeOf(res).Kind() != reflect.Bool || !(res.(bool)) {
			return "", "", nil
		}
		if query, order, err = nodesToQuery(e.nodes, param, page); err != nil {
			return "", "", err
		}
	case tagForeach:
		var err error
		var text string
		if text, order, err = nodesToQuery(e.nodes, param, page); err != nil {
			return "", "", err
		}
		query = e.attr["open"]
		pm := e.attr["params"]
		separator := e.attr["separator"]
		var val interface{}
		if param == nil || param[pm] == nil {
			val = make([]interface{}, 0)
		} else {
			val = param[pm]
		}
		switch reflect.TypeOf(val).Kind() {
		case reflect.Map:
			idx := 0
			var vm map[string]string
			if vm, err = jcast.StringMapString(val); err != nil {
				return "", "", errorFmt(errorWrongTypeOfForeach)
			}
			for k, v := range vm {
				str := text
				str = strings.ReplaceAll(str, "#{key}", k)
				str = strings.ReplaceAll(str, "#{val}", v)
				if idx > 0 {
					query += fmt.Sprint(separator, " ")
				}
				query += trim(str)
				idx++
			}
		case reflect.Slice:
			var vs []string
			if vs, err = jcast.SliceString(val); err != nil {
				return "", "", errorFmt(errorWrongTypeOfForeach)
			}
			for idx, v := range vs {
				str := text
				str = strings.ReplaceAll(str, "#{val}", v)
				if idx > 0 {
					query += fmt.Sprint(separator, " ")
				}
				query += trim(str)
			}
		default:
			return "", "", errorStr(errorWrongTypeOfForeach)
		}
		query += e.attr["close"]
	case tagWhere:
		var err error
		if query, order, err = nodesToQuery(e.nodes, param, page); err != nil {
			return "", "", err
		}
		var sw bool
		if sw, err = regexp.MatchString(fmt.Sprint("^", e.tag.String(), "[^A-Za-z0-9_]"), strings.ToLower(query)); err != nil {
			return "", "", err
		}
		if !sw {
			var sa bool
			if sa, err = regexp.MatchString(fmt.Sprint("^", And.String(), "[^A-Za-z0-9_]"), strings.ToUpper(query)); err != nil {
				return "", "", err
			}
			if sa {
				query = trim(query[len(And.String()):])
			}
			var so bool
			if so, err = regexp.MatchString(fmt.Sprint("^", Or.String(), "[^A-Za-z0-9_]"), strings.ToUpper(query)); err != nil {
				return "", "", err
			}
			if so {
				query = trim(query[len(Or.String()):])
			}
			if len(query) > 0 {
				query = fmt.Sprint("WHERE ", query)
			}
		}
	case tagOrderBy:
		var err error
		if query, order, err = nodesToQuery(e.nodes, param, page); err != nil {
			return "", "", err
		}
		var last bool
		if last, err = jcast.Bool(e.attr["last"]); err != nil {
			return "", "", err
		}
		if page && last {
			order = query
			query = ""
		} else {
			if len(query) > 0 {
				query = fmt.Sprint("ORDER BY ", query)
			}
			order = ""
		}
	}
	return trim(query), trim(order), nil
}

func nodesToQuery(nodes []*element, param map[string]interface{}, page bool) (string, string, error) {
	query := ""
	order := ""
	for i, node := range nodes {
		if s, o, err := node.getSql(param, page); err != nil {
			return "", "", err
		} else {
			if i == 0 {
				query = fmt.Sprint(query, s)
			} else {
				query = fmt.Sprint(query, " ", s)
			}
			if o != "" {
				order = o
			}
			query = trim(query)
		}
	}
	return query, order, nil
}
