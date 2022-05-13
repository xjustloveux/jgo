// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"github.com/sirupsen/logrus"
)

type logger struct {
	Level     string
	Formatter *formatter
	Output    *output
}

func (l logger) getLogrusLevel() logrus.Level {
	if _, err := logrus.ParseLevel(l.Level); err != nil {
		l.Level = "info"
		fmtPrintln(err)
	}
	lv, _ := logrus.ParseLevel(l.Level)
	return lv
}

func (l logger) getDefault() *logger {
	return &logger{
		Level:     "info",
		Formatter: formatter{}.getDefault(),
		Output:    output{}.getDefault(),
	}
}
