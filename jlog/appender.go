// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import "github.com/sirupsen/logrus"

type appender struct {
	Level     string
	Formatter *formatter
	Output    *Output
}

func (*appender) getDefault() *appender {
	return &appender{
		Level:     "info",
		Formatter: (&formatter{}).getDefault(),
		Output:    (&Output{}).getDefault(),
	}
}

func (a *appender) getLogrusLevel() logrus.Level {
	if _, err := logrus.ParseLevel(a.Level); err != nil {
		a.Level = "info"
		subject.Next(err)
	}
	lv, _ := logrus.ParseLevel(a.Level)
	return lv
}
