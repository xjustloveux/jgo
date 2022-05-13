// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

type Format string

const (
	Json Format = "Json"
	Yml  Format = "Yml"
	Yaml Format = "Yaml"
	Toml Format = "Toml"
	Text Format = "Text"
	Xml  Format = "Xml"
)

func (f Format) String() string {
	return string(f)
}
