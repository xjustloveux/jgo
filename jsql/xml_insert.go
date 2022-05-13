// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type xmlInsert struct {
	Id      string       `xml:"id,attr"`
	Data    string       `xml:",chardata"`
	Xml     string       `xml:",innerxml"`
	If      []xmlIf      `xml:"if"`
	For     []xmlFor     `xml:"foreach"`
	OrderBy []xmlOrderBy `xml:"orderBy"`
}

func (xi *xmlInsert) getSql(param map[string]interface{}) (string, string, error) {
	return xmlToSql(xi.Xml, param, xi.If, xi.For, xi.OrderBy, false)
}
