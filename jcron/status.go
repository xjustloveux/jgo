// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"strings"
)

type Status int

const (
	Stop Status = iota
	Run
	SyncWait
	Unknown = -1
)

// ParseStatus takes a string Status and returns the Status constant.
func ParseStatus(s string) (Status, error) {
	switch strings.ToLower(s) {
	case "stop":
		return Stop, nil
	case "run":
		return Run, nil
	case "syncwait":
		return SyncWait, nil
	}
	return Unknown, errorf(errorNotValidStatus, s)
}
