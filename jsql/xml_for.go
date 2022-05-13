// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"reflect"
	"strings"
)

type xmlFor struct {
	Data      string       `xml:",chardata"`
	Xml       string       `xml:",innerxml"`
	Params    string       `xml:"params,attr"`
	Open      string       `xml:"open,attr"`
	Separator string       `xml:"separator,attr"`
	Close     string       `xml:"close,attr"`
	If        []xmlIf      `xml:"if"`
	For       []xmlFor     `xml:"foreach"`
	OrderBy   []xmlOrderBy `xml:"orderBy"`
}

func (xf xmlFor) getSql(params map[string]interface{}, page bool) (query, order string, err error) {
	query = xf.Open
	var val interface{}
	if params == nil || params[xf.Params] == nil {
		val = make([]interface{}, 0)
	} else {
		val = params[xf.Params]
	}
	switch reflect.TypeOf(val).Kind() {
	case reflect.Map:
		idx := 0
		if vm, ok := val.(map[string]string); ok {
			for k, v := range vm {
				var xml, obs string
				if xml, obs, err = replaceXmlForeach(xf.If, xf.For, xf.OrderBy, xf.Xml, params, page); err != nil {
					return "", "", err
				}
				if obs != "" {
					order = obs
				}
				xml = strings.ReplaceAll(xml, "#{key}", k)
				xml = strings.ReplaceAll(xml, "#{val}", v)
				if idx > 0 {
					query += xf.Separator
				}
				query += xml
				idx++
			}
		} else {
			return "", "", errorf(errorWrongTypeOfForeach)
		}
	case reflect.Slice:
		if vs, ok := val.([]string); ok {
			for idx, v := range vs {
				var xml, obs string
				if xml, obs, err = replaceXmlForeach(xf.If, xf.For, xf.OrderBy, xf.Xml, params, page); err != nil {
					return "", "", err
				}
				if obs != "" {
					order = obs
				}
				xml = strings.ReplaceAll(xml, "#{val}", v)
				if idx > 0 {
					query += xf.Separator
				}
				query += xml
			}
		} else {
			return "", "", errorf(errorWrongTypeOfForeach)
		}
	default:
		return "", "", errors(errorWrongTypeOfForeach)
	}
	query += xf.Close
	return query, order, nil
}

func (xf xmlFor) replaceXml(xml string, params map[string]interface{}, page bool) (query, order string, err error) {
	st := "<foreach"
	et := "</foreach>"
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
		return "", "", errors(errorWrongNumOfForeach)
	}
	var s string
	if s, order, err = xf.getSql(params, page); err != nil {
		return "", "", err
	} else {
		query = fmt.Sprint(xml[0:si], s, xml[ei+len(et):])
	}
	return query, order, nil
}
