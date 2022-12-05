// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"github.com/sirupsen/logrus"
	"github.com/xjustloveux/jgo/jtime"
	"time"
)

type formatter struct {
	Type     string
	Location string
	Text     *logrus.TextFormatter
	Json     *logrus.JSONFormatter
}

func (*formatter) getDefault() *formatter {
	return &formatter{
		Type:     "TEXT",
		Location: "",
		Text:     &logrus.TextFormatter{TimestampFormat: jtime.DateTime},
		Json:     &logrus.JSONFormatter{TimestampFormat: jtime.DateTime},
	}
}

type timeFormatter struct {
	loc *time.Location
	log logrus.Formatter
}

func (tf timeFormatter) Format(e *logrus.Entry) ([]byte, error) {
	e.Time = e.Time.In(tf.loc)
	return tf.log.Format(e)
}
