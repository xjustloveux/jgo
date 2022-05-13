// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

type configData struct {
	Debug    bool
	Params   map[string]interface{}
	Appender map[string]interface{}
	Logs     []*logs
}
