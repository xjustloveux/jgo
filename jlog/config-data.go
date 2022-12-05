// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

type configData struct {
	Log *configPack
}

type configPack struct {
	Params   map[string]string
	Default  []string
	Appender map[string]map[string]interface{}
	Logs     []*logs
	appender map[string]*appender
}
