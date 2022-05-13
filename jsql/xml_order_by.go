// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"strings"
)

type xmlOrderBy struct {
	Data    string       `xml:",chardata"`
	Xml     string       `xml:",innerxml"`
	LastTag bool         `xml:"lastTag,attr"`
	If      []xmlIf      `xml:"if"`
	For     []xmlFor     `xml:"foreach"`
	OrderBy []xmlOrderBy `xml:"orderBy"`
}

func (xo xmlOrderBy) getSql(params map[string]interface{}, page bool) (query, order string, err error) {
	var obs string
	if query, obs, err = replaceXmlForeach(xo.If, xo.For, xo.OrderBy, xo.Xml, params, page); err != nil {
		return "", "", err
	}
	if obs != "" {
		order = obs
	} else {
		order = ""
	}
	return query, order, nil
}

func (xo xmlOrderBy) replaceXml(xml string, params map[string]interface{}, page bool) (query, order string, err error) {
	st := "<orderBy"
	et := "</orderBy>"
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
		return "", "", errors(errorWrongNumOfOrderBy)
	}
	var s string
	if s, order, err = xo.getSql(params, page); err != nil {
		return "", "", err
	} else {
		if page && xo.LastTag {
			query = fmt.Sprint(xml[0:si], xml[ei+len(et):])
			order = s
		} else {
			query = fmt.Sprint(xml[0:si], " ORDER BY ", s, xml[ei+len(et):])
			order = ""
		}
	}
	return query, order, nil
}
