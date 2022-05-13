// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type xmlDelete struct {
	Id      string       `xml:"id,attr"`
	Data    string       `xml:",chardata"`
	Xml     string       `xml:",innerxml"`
	If      []xmlIf      `xml:"if"`
	For     []xmlFor     `xml:"foreach"`
	OrderBy []xmlOrderBy `xml:"orderBy"`
}

func (xd *xmlDelete) getSql(param map[string]interface{}) (string, string, error) {
	return xmlToSql(xd.Xml, param, xd.If, xd.For, xd.OrderBy, false)
}
