// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type xmlSql struct {
	Select []xmlSelect `xml:"select"`
	Insert []xmlInsert `xml:"insert"`
	Update []xmlUpdate `xml:"update"`
	Delete []xmlDelete `xml:"delete"`
	Other  []xmlOther  `xml:"other"`
}
