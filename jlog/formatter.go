// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"github.com/sirupsen/logrus"
	"github.com/xjustloveux/jgo/jtime"
)

type formatter struct {
	Type string
	Text *logrus.TextFormatter
	Json *logrus.JSONFormatter
}

func (*formatter) getDefault() *formatter {
	return &formatter{
		Type: "TEXT",
		Text: &logrus.TextFormatter{TimestampFormat: jtime.DateTime},
		Json: &logrus.JSONFormatter{TimestampFormat: jtime.DateTime},
	}
}
