// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jruntime

import (
	"runtime"
	"strings"
)

// GetFuncName returns function name
func GetFuncName(sep ...string) string {
	if pc, _, _, ok := runtime.Caller(1); ok {
		s := "."
		if len(sep) > 0 {
			s = sep[0]
		}
		path := runtime.FuncForPC(pc).Name()
		pathArr := strings.Split(path, s)
		if pl := len(pathArr); pl > 0 {
			return pathArr[pl-1]
		}
		return path
	}
	return ""
}

// GetCallerName returns caller function name
func GetCallerName(sep ...string) string {
	if pc, _, _, ok := runtime.Caller(2); ok {
		s := "."
		if len(sep) > 0 {
			s = sep[0]
		}
		path := runtime.FuncForPC(pc).Name()
		pathArr := strings.Split(path, s)
		if pl := len(pathArr); pl > 0 {
			return pathArr[pl-1]
		}
		return path
	}
	return ""
}

// GetCallerProgramName returns caller program name
func GetCallerProgramName(sep ...string) string {
	if _, path, _, ok := runtime.Caller(2); ok {
		s := "/"
		if len(sep) > 0 {
			s = sep[0]
		}
		pathArr := strings.Split(path, s)
		if pl := len(pathArr); pl > 0 {
			return pathArr[pl-1]
		}
		return path
	}
	return ""
}
