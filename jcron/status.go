// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import "strings"

type Status int

const (
	Stop Status = iota
	Run
	SyncWait
	Unknown = -1
)

func parseStatus(str string) Status {
	switch strings.ToLower(str) {
	case "stop":
		return Stop
	case "run":
		return Run
	case "syncwait":
		return SyncWait
	default:
		return Unknown
	}
}
