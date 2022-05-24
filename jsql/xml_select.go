// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type xmlSelect struct {
	Id      string       `xml:"id,attr"`
	Data    string       `xml:",chardata"`
	Xml     string       `xml:",innerxml"`
	If      []xmlIf      `xml:"if"`
	For     []xmlFor     `xml:"foreach"`
	OrderBy []xmlOrderBy `xml:"orderBy"`
}

func (xs *xmlSelect) getSql(param map[string]interface{}, page bool) (string, string, error) {
	return xmlToSql(Select, xs.Xml, param, xs.If, xs.For, xs.OrderBy, page)
}
