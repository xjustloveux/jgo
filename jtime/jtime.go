// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jtime

import (
	"fmt"
	"strconv"
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

func FormatString(str string, t time.Time) string {
	arr := strings.Split(str, "%%")
	format := []string{
		"dddd", "ddd", "dd", "d", "DDD", "DD", "D",
		"ffff", "fff", "ff", "f",
		"FFFF", "FFF", "FF", "F",
		"g",
		"hh", "h", "HH", "H",
		"kk", "k", "KK", "K",
		"l",
		"mm", "m", "MMMM", "MMM", "MM", "M",
		"ss", "s",
		"tt", "t",
		"w", "W",
		"yyyy", "yyy", "yy", "y",
		"zzzz", "zzz", "zz", "z", "Z",
	}
	v := ""
	for i, s := range arr {
		fs := s
		if i > 0 {
			v = fmt.Sprint("%", v)
		}
		for _, f := range format {
			fs = formatString(f, fs, t)
		}
		v = fmt.Sprint(v, fs)
	}
	return v
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func formatString(format, str string, t time.Time) string {
	k := "%"
	switch format {
	case "d":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("2"))
	case "dd":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("02"))
	case "ddd":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Weekday().String()[:3])
	case "dddd":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Weekday().String())
	case "D":
		fallthrough
	case "DD":
		fallthrough
	case "DDD":
		l := len(format)
		s := fmt.Sprintf("%v", t.YearDay())
		for len(s) < l {
			s = fmt.Sprint("0", s)
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "f":
		fallthrough
	case "ff":
		fallthrough
	case "fff":
		fallthrough
	case "ffff":
		l := len(format)
		s := fmt.Sprintf("%v", t.Nanosecond())
		for len(s) < l {
			s = fmt.Sprint(s, "0")
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s[:l])
	case "F":
		fallthrough
	case "FF":
		fallthrough
	case "FFF":
		fallthrough
	case "FFFF":
		l := len(format)
		s := fmt.Sprintf("%v", t.Nanosecond())
		for len(s) < l {
			s = fmt.Sprint(s, "0")
		}
		s = s[:l]
		if i, err := strconv.ParseInt(s, 0, 64); err != nil {
			fmt.Println(err)
			str = strings.ReplaceAll(str, fmt.Sprint(k, format), "")
		} else if i != 0 {
			str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
		} else {
			str = strings.ReplaceAll(str, fmt.Sprint(k, format), "")
		}
	case "g":
		var s string
		if t.Year() <= 0 {
			s = "B.C."
		} else {
			s = "A.D."
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "h":
		fallthrough
	case "hh":
		s := fmt.Sprintf("%v", t.Hour()%12)
		for len(s) < len(format) {
			s = fmt.Sprint("0", s)
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "H":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), fmt.Sprintf("%v", t.Hour()))
	case "HH":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("15"))
	case "k":
		fallthrough
	case "kk":
		s := fmt.Sprintf("%v", (t.Hour()+1)%12)
		for len(s) < len(format) {
			s = fmt.Sprint("0", s)
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "K":
		fallthrough
	case "KK":
		s := fmt.Sprintf("%v", t.Hour()+1)
		for len(s) < len(format) {
			s = fmt.Sprint("0", s)
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "l":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Location().String())
	case "m":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), fmt.Sprintf("%v", t.Minute()))
	case "mm":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("04"))
	case "M":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("1"))
	case "MM":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("01"))
	case "MMM":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Month().String()[:3])
	case "MMMM":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Month().String())
	case "s":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), fmt.Sprintf("%v", t.Second()))
	case "ss":
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("05"))
	case "t":
		fallthrough
	case "tt":
		l := len(format)
		var s string
		if t.Hour() < 12 {
			s = "AM"
		} else {
			s = "PM"
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s[:l])
	case "w":
		_, w := t.ISOWeek()
		s := fmt.Sprintf("%v", w)
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "W":
		s := fmt.Sprintf("%v", (int)(t.Weekday()))
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "y":
		fallthrough
	case "yy":
		fallthrough
	case "yyy":
		fallthrough
	case "yyyy":
		t.Zone()
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), t.Format("2006")[:len(format)])
	case "z":
		fallthrough
	case "zz":
		fallthrough
	case "zzz":
		l := len(format)
		_, sec := t.Zone()
		negative := sec < 0
		if negative {
			sec = sec * -1
		}
		min := sec / 60
		hour := min / 60
		s := fmt.Sprintf("%v", hour)
		for len(s) < l && len(s) < 3 {
			s = fmt.Sprint("0", s)
		}
		if l >= 3 {
			m := fmt.Sprintf("%v", min)
			for len(m) < 2 {
				m = fmt.Sprint("0", m)
			}
			s = fmt.Sprint(s, ":", m)
		}
		if negative {
			s = fmt.Sprint("-", s)
		}
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "zzzz":
		_, sec := t.Zone()
		s := fmt.Sprintf("%v", sec)
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	case "Z":
		s, _ := t.Zone()
		str = strings.ReplaceAll(str, fmt.Sprint(k, format), s)
	}
	return str
}
