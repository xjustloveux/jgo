// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

type configData struct {
	Db *configPack
}

type configPack struct {
	DaoPath    string
	Default    string
	DataSource map[string]interface{}
}
