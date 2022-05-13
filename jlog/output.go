// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import "time"

type output struct {
	P                    string
	UTC                  bool
	LinkName             string
	MaxAge               time.Duration
	MaxAgeDuration       string
	RotationTime         time.Duration
	RotationTimeDuration string
	RotationSize         int64
	RotationSizeUnit     string
	RotationCount        uint
}

func (o output) getDefault() *output {
	return &output{
		P:                    "",
		UTC:                  false,
		LinkName:             "",
		MaxAge:               365,
		MaxAgeDuration:       "Day",
		RotationTime:         24,
		RotationTimeDuration: "Hour",
		RotationSize:         10,
		RotationSizeUnit:     "MB",
		RotationCount:        0,
	}
}
