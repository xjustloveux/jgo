// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"strconv"
	"strings"
)

type CronExpression struct {
	second  uint64
	minute  uint64
	hour    uint64
	day     uint64
	month   uint64
	weekday uint64
	year    uint64
}

// ParseCronExpression takes a string CronExpression and returns the *CronExpression struct.
func ParseCronExpression(c string) (*CronExpression, error) {
	arr := strings.Split(strings.Trim(c, " "), " ")
	if len(arr) < 6 {
		return nil, errorFmt(errorNotValidCronExpression, c)
	}
	ce := &CronExpression{0, 0, 0, 0, 0, 0, 0}
	if second, err := parseCronExpression(arr[0], 0, 59, false, nil); err != nil {
		return nil, err
	} else {
		if second <= 0 {
			return nil, errorFmt(errorNotValidCronExpression, c)
		}
		ce.second = second
	}
	if minute, err := parseCronExpression(arr[1], 0, 59, false, nil); err != nil {
		return nil, err
	} else {
		if minute <= 0 {
			return nil, errorFmt(errorNotValidCronExpression, c)
		}
		ce.minute = minute
	}
	if hour, err := parseCronExpression(arr[2], 0, 23, false, nil); err != nil {
		return nil, err
	} else {
		if hour <= 0 {
			return nil, errorFmt(errorNotValidCronExpression, c)
		}
		ce.hour = hour
	}
	if day, err := parseCronExpression(arr[3], 1, 31, false, nil); err != nil {
		return nil, err
	} else {
		ce.day = day
	}
	if month, err := parseCronExpression(arr[4], 1, 12, false, parseMonth); err != nil {
		return nil, err
	} else {
		if month <= 0 {
			return nil, errorFmt(errorNotValidCronExpression, c)
		}
		ce.month = month
	}
	if weekday, err := parseCronExpression(arr[5], 0, 6, false, parseWeekday); err != nil {
		return nil, err
	} else {
		ce.weekday = weekday
	}
	if (ce.day <= 0 && ce.weekday <= 0) ||
		(ce.day > 0 && ce.weekday > 0) {
		return nil, errorFmt(errorNotValidCronExpression, c)
	}
	if len(arr) > 6 {
		if year, err := parseCronExpression(arr[6], 0, maxYear-minYear, true, nil); err != nil {
			return nil, err
		} else {
			ce.year = year
		}
	} else {
		for i := 0; i <= maxYear-minYear; i++ {
			ce.year |= 1 << i
		}
	}
	return ce, nil
}

func parseCronExpression(c string, min, max int, year bool, p func(string) int64) (uint64, error) {
	arr := strings.Split(strings.Trim(c, " "), ",")
	if len(arr) <= 0 {
		return 0, errorFmt(errorNotValidCronExpression, c)
	}
	var res uint64
	res = 0
	for _, str := range arr {
		str = strings.Trim(str, " ")
		if str == "?" {
			continue
		}
		if str == "*" {
			for i := min; i <= max; i++ {
				res |= 1 << i
			}
			continue
		}
		{
			var v int64
			hav := false
			if p != nil {
				if v = p(str); v != Unknown {
					hav = true
				}
			}
			if !hav {
				var err error
				if v, err = strconv.ParseInt(str, 10, 64); err == nil {
					hav = true
				}
			}
			if hav {
				var nv int
				if v64, err := strconv.ParseInt(strconv.FormatInt(v, 10), 10, strconv.IntSize); err != nil {
					return 0, err
				} else {
					nv = int(v64)
				}
				if year {
					nv -= minYear
				}
				if nv < min || nv > max {
					return 0, errorFmt(errorNotValidCronExpression, c)
				}
				res |= 1 << nv
				continue
			}
		}
		arr2 := strings.Split(str, "/")
		switch len(arr2) {
		case 1:
			if arr3 := strings.Split(str, "-"); len(arr3) == 2 {
				if sn, en, valid := getCronExpressionSE(arr3, min, max, year, p); valid {
					for i := sn; i <= en; i++ {
						res |= 1 << i
					}
				} else {
					return 0, errorFmt(errorNotValidCronExpression, c)
				}
			} else {
				return 0, errorFmt(errorNotValidCronExpression, c)
			}
		case 2:
			var err error
			var b64 int64
			str2 := strings.Trim(arr2[1], " ")
			if b64, err = strconv.ParseInt(str2, 10, 64); err != nil {
				return 0, errorFmt(errorNotValidCronExpression, c)
			}
			var b int
			var v64 int64
			if v64, err = strconv.ParseInt(strconv.FormatInt(b64, 10), 10, strconv.IntSize); err != nil {
				return 0, err
			} else {
				b = int(v64)
			}
			str2 = strings.Trim(arr2[0], " ")
			var sn, en int
			if str2 == "*" {
				sn = min
				en = max
			} else {
				if arr3 := strings.Split(str2, "-"); len(arr3) == 2 {
					var valid bool
					if sn, en, valid = getCronExpressionSE(arr3, min, max, year, p); !valid {
						return 0, errorFmt(errorNotValidCronExpression, c)
					}
				} else {
					var v int64
					hav := false
					if p != nil {
						if v = p(str2); v != Unknown {
							hav = true
						}
					}
					if !hav {
						if v, err = strconv.ParseInt(str2, 10, 64); err == nil {
							hav = true
						}
					}
					if hav {
						if v64, err = strconv.ParseInt(strconv.FormatInt(v, 10), 10, strconv.IntSize); err != nil {
							return 0, err
						} else {
							sn = int(v64)
						}
						en = max
						if year {
							sn -= minYear
						}
						if sn < min || sn > max {
							return 0, errorFmt(errorNotValidCronExpression, c)
						}
					} else {
						return 0, errorFmt(errorNotValidCronExpression, c)
					}
				}
			}
			for i := sn; i <= en; i += b {
				res |= 1 << i
			}
		default:
			return 0, errorFmt(errorNotValidCronExpression, c)
		}
	}
	return res, nil
}

func getCronExpressionSE(arr []string, min, max int, year bool, p func(string) int64) (sn, en int, valid bool) {
	var err error
	var sn64, en64 int64
	has := false
	hae := false
	if p != nil {
		if sn64 = p(arr[0]); sn64 != Unknown {
			has = true
		}
		if en64 = p(arr[1]); en64 != Unknown {
			hae = true
		}
	}
	if !has {
		if sn64, err = strconv.ParseInt(arr[0], 10, 64); err != nil {
			return 0, 0, false
		}
	}
	if !hae {
		if en64, err = strconv.ParseInt(arr[1], 10, 64); err != nil {
			return 0, 0, false
		}
	}
	var v64 int64
	if v64, err = strconv.ParseInt(strconv.FormatInt(sn64, 10), 10, strconv.IntSize); err != nil {
		return 0, 0, false
	} else {
		sn = int(v64)
	}
	if v64, err = strconv.ParseInt(strconv.FormatInt(en64, 10), 10, strconv.IntSize); err != nil {
		return 0, 0, false
	} else {
		en = int(v64)
	}
	if year {
		sn -= minYear
		en -= minYear
	}
	if sn < min || sn > max || en < min || en > max || sn > en {
		return 0, 0, false
	}
	return sn, en, true
}

func parseMonth(str string) int64 {
	switch strings.ToLower(str) {
	case "january":
		fallthrough
	case "jan":
		return 1
	case "february":
		fallthrough
	case "feb":
		return 2
	case "march":
		fallthrough
	case "mar":
		return 3
	case "april":
		fallthrough
	case "apr":
		return 4
	case "may":
		return 5
	case "june":
		fallthrough
	case "jun":
		return 6
	case "july":
		fallthrough
	case "jul":
		return 7
	case "august":
		fallthrough
	case "aug":
		return 8
	case "september":
		fallthrough
	case "sep":
		return 9
	case "october":
		fallthrough
	case "oct":
		return 10
	case "november":
		fallthrough
	case "nov":
		return 11
	case "december":
		fallthrough
	case "dec":
		return 12
	default:
		return Unknown
	}
}

func parseWeekday(str string) int64 {
	switch strings.ToLower(str) {
	case "sunday":
		fallthrough
	case "sun":
		return 0
	case "monday":
		fallthrough
	case "mon":
		return 1
	case "tuesday":
		fallthrough
	case "tue":
		return 2
	case "wednesday":
		fallthrough
	case "wed":
		return 3
	case "thursday":
		fallthrough
	case "thu":
		return 4
	case "friday":
		fallthrough
	case "fri":
		return 5
	case "saturday":
		fallthrough
	case "sat":
		return 6
	default:
		return Unknown
	}
}
