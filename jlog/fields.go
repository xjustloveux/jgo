// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import "github.com/sirupsen/logrus"

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

func (f Fields) toLogrusFields() logrus.Fields {
	var m map[string]interface{}
	m = f
	return m
}
