// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jtime

import (
	"fmt"
	"strings"
	"time"
)

const (
	errorNotValidTimeDuration = jError("not a valid time.Duration %q")
)

const (
	ANSIC       = time.ANSIC
	UnixDate    = time.UnixDate
	RubyDate    = time.RubyDate
	RFC822      = time.RFC822
	RFC822Z     = time.RFC822Z
	RFC850      = time.RFC850
	RFC1123     = time.RFC1123
	RFC1123Z    = time.RFC1123Z
	RFC3339     = time.RFC3339
	RFC3339Nano = time.RFC3339Nano
	Kitchen     = time.Kitchen
	Stamp       = time.Stamp
	StampMilli  = time.StampMilli
	StampMicro  = time.StampMicro
	StampNano   = time.StampNano
	ISO8601     = "2006-01-02T15:04:05"
	DateTime    = "2006-01-02 15:04:05"
	DateS       = "2006-01-02 15:04:05 -0700 MST"
	Date        = "2006-01-02"
	Time        = "15:04:05"
)

const (
	Nanosecond  = time.Nanosecond
	Microsecond = time.Microsecond
	Millisecond = time.Millisecond
	Second      = time.Second
	Minute      = time.Minute
	Hour        = time.Hour
	Day         = 24 * time.Hour
	Unknown     = -1
)

const (
	pkgName = "jtime"
)

// ParseTimeDuration takes a string time.Duration and returns the time.Duration constant
func ParseTimeDuration(t string) (time.Duration, error) {
	switch strings.ToLower(t) {
	case "nanosecond":
		return Nanosecond, nil
	case "microsecond":
		return Microsecond, nil
	case "millisecond":
		return Millisecond, nil
	case "second":
		return Second, nil
	case "minute":
		return Minute, nil
	case "hour":
		return Hour, nil
	case "day":
		return Day, nil
	}
	return Unknown, errorf(errorNotValidTimeDuration, t)
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}
