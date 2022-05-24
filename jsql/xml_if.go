// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"github.com/Knetic/govaluate"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type xmlIf struct {
	Data    string       `xml:",chardata"`
	Xml     string       `xml:",innerxml"`
	Test    string       `xml:"test,attr"`
	If      []xmlIf      `xml:"if"`
	For     []xmlFor     `xml:"foreach"`
	OrderBy []xmlOrderBy `xml:"orderBy"`
}

func (xi xmlIf) getSql(params map[string]interface{}, page bool) (query, order string, err error) {
	pattern := regexp.MustCompile(`nil\(\w+\)`)
	test := ""
	testIdx := 0
	for _, v := range pattern.FindAllStringSubmatchIndex(xi.Test, -1) {
		st := "nil("
		et := ")"
		k := xi.Test[v[0]+len(st) : v[1]-len(et)]
		test = fmt.Sprint(test, xi.Test[testIdx:v[0]], strconv.FormatBool(params == nil || params[k] == nil))
		testIdx = v[1]
	}
	test = fmt.Sprint(test, xi.Test[testIdx:])
	test = strings.ReplaceAll(test, " and ", " && ")
	test = strings.ReplaceAll(test, " or ", " || ")
	var expression *govaluate.EvaluableExpression
	if expression, err = govaluate.NewEvaluableExpression(test); err != nil {
		return "", "", err
	}
	var result interface{}
	if result, err = expression.Evaluate(params); err != nil {
		return "", "", err
	}
	if reflect.TypeOf(result).Kind() != reflect.Bool || !(result.(bool)) {
		return "", "", nil
	}
	var obs string
	if query, obs, err = replaceXmlForeach(xi.If, xi.For, xi.OrderBy, xi.Xml, params, page); err != nil {
		return "", "", err
	}
	if obs != "" {
		order = obs
	}
	return trim(query), order, nil
}

func (xi xmlIf) replaceXml(xml string, params map[string]interface{}, page bool) (query, order string, err error) {
	st := "<if"
	et := "</if>"
	si := strings.Index(xml, st)
	ei := strings.Index(xml, et)
	if nsi := strings.Index(xml[si+len(st):], st); nsi >= 0 {
		nsi += si + len(st)
		for nsi >= 0 && ei > nsi {
			if i := strings.Index(xml[ei+len(et):], et); i >= 0 {
				ei += i + len(et)
			}
			if i := strings.Index(xml[nsi+len(st):], st); i >= 0 {
				nsi += i + len(st)
			} else {
				nsi = i
			}
		}
	}
	if si < 0 || ei < 0 {
		return "", "", errors(errorWrongNumOfIf)
	}
	var s string
	if s, order, err = xi.getSql(params, page); err != nil {
		return "", "", err
	} else {
		query = fmt.Sprint(xml[0:si], s, xml[ei+len(et):])
	}
	return query, order, nil
}
