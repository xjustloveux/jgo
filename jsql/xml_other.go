// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type xmlOther struct {
	Id      string       `xml:"id,attr"`
	Data    string       `xml:",chardata"`
	Xml     string       `xml:",innerxml"`
	If      []xmlIf      `xml:"if"`
	For     []xmlFor     `xml:"foreach"`
	OrderBy []xmlOrderBy `xml:"orderBy"`
}

func (xo *xmlOther) getSql(param map[string]interface{}) (string, string, error) {
	return xmlToSql(xo.Xml, param, xo.If, xo.For, xo.OrderBy, false)
}
