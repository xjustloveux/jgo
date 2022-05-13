// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"sort"
	"strconv"
	"strings"
)

type CronExpression struct {
	second  []int
	minute  []int
	hour    []int
	day     []int
	month   []int
	weekday []int
	year    []int
}

// ParseCronExpression takes a string CronExpression and returns the *CronExpression struct.
func ParseCronExpression(c string) (*CronExpression, error) {
	arr := strings.Split(strings.Trim(c, " "), " ")
	if len(arr) < 6 {
		return nil, errorf(errorNotValidCronExpression, c)
	}
	ce := &CronExpression{}
	if second, err := parseCronExpression(arr[0], 0, 59); err != nil {
		return nil, err
	} else {
		if len(second) <= 0 {
			return nil, errorf(errorNotValidCronExpression, c)
		}
		ce.second = second
	}
	if minute, err := parseCronExpression(arr[1], 0, 59); err != nil {
		return nil, err
	} else {
		if len(minute) <= 0 {
			return nil, errorf(errorNotValidCronExpression, c)
		}
		ce.minute = minute
	}
	if hour, err := parseCronExpression(arr[2], 0, 23); err != nil {
		return nil, err
	} else {
		if len(hour) <= 0 {
			return nil, errorf(errorNotValidCronExpression, c)
		}
		ce.hour = hour
	}
	if day, err := parseCronExpression(arr[3], 1, 31); err != nil {
		return nil, err
	} else {
		ce.day = day
	}
	if month, err := parseCronExpression(arr[4], 1, 12); err != nil {
		return nil, err
	} else {
		if len(month) <= 0 {
			return nil, errorf(errorNotValidCronExpression, c)
		}
		ce.month = month
	}
	if weekday, err := parseCronExpression(arr[5], 0, 6); err != nil {
		return nil, err
	} else {
		ce.weekday = weekday
	}
	if (len(ce.day) <= 0 && len(ce.weekday) <= 0) ||
		(len(ce.day) > 0 && len(ce.weekday) > 0) {
		return nil, errorf(errorNotValidCronExpression, c)
	}
	if len(arr) > 6 {
		if year, err := parseCronExpression(arr[6], minYear, maxYear); err != nil {
			return nil, err
		} else {
			ce.year = year
		}
	} else {
		for i := minYear; i <= maxYear; i++ {
			ce.year = addCronExpressionSlice(ce.year, i)
		}
	}
	return ce, nil
}

func parseCronExpression(c string, min, max int) ([]int, error) {
	arr := strings.Split(strings.Trim(c, " "), ",")
	if len(arr) <= 0 {
		return nil, errorf(errorNotValidCronExpression, c)
	}
	res := make([]int, 0)
	for _, str := range arr {
		str = strings.Trim(str, " ")
		if str == "?" {
			continue
		}
		if str == "*" {
			for i := min; i <= max; i++ {
				res = addCronExpressionSlice(res, i)
			}
			continue
		}
		if v, err := strconv.ParseInt(str, 10, 64); err == nil {
			nv := int(v)
			if nv < min || nv > max {
				return nil, errorf(errorNotValidCronExpression, c)
			}
			res = addCronExpressionSlice(res, int(v))
			continue
		}
		arr2 := strings.Split(str, "/")
		switch len(arr2) {
		case 1:
			if arr3 := strings.Split(str, "-"); len(arr3) == 2 {
				if sn, en, valid := getCronExpressionSE(arr3, min, max); valid {
					for i := sn; i <= en; i++ {
						res = addCronExpressionSlice(res, i)
					}
				} else {
					return nil, errorf(errorNotValidCronExpression, c)
				}
			} else {
				return nil, errorf(errorNotValidCronExpression, c)
			}
		case 2:
			var err error
			var b64 int64
			str2 := strings.Trim(arr2[1], " ")
			if b64, err = strconv.ParseInt(str2, 10, 64); err != nil {
				return nil, errorf(errorNotValidCronExpression, c)
			}
			b := int(b64)
			str2 = strings.Trim(arr2[0], " ")
			var sn, en int
			if str2 == "*" {
				sn = min
				en = max
			} else {
				var sn64 int64
				if arr3 := strings.Split(str2, "-"); len(arr3) == 2 {
					var valid bool
					if sn, en, valid = getCronExpressionSE(arr3, min, max); !valid {
						return nil, errorf(errorNotValidCronExpression, c)
					}
				} else if sn64, err = strconv.ParseInt(str2, 10, 64); err != nil {
					return nil, errorf(errorNotValidCronExpression, c)
				} else {
					sn = int(sn64)
					en = max
				}
			}
			for i := sn; i <= en; i += b {
				res = addCronExpressionSlice(res, i)
			}
		default:
			return nil, errorf(errorNotValidCronExpression, c)
		}
	}
	sort.Ints(res)
	return res, nil
}

func addCronExpressionSlice(s []int, n int) []int {
	add := true
	for _, v := range s {
		if v == n {
			add = false
			break
		}
	}
	if add {
		return append(s, n)
	}
	return s
}

func getCronExpressionSE(arr []string, min, max int) (sn, en int, valid bool) {
	var err error
	var sn64, en64 int64
	if sn64, err = strconv.ParseInt(arr[0], 10, 64); err != nil {
		return 0, 0, false
	}
	if en64, err = strconv.ParseInt(arr[1], 10, 64); err != nil {
		return 0, 0, false
	}
	sn = int(sn64)
	en = int(en64)
	if sn < min || sn > max || en < min || en > max || sn > en {
		return 0, 0, false
	}
	return sn, en, true
}
