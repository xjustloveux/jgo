// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jslice

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
)

const (
	errorIndexOut = jError("index %d out of bounds %d")
	//"array index %d is out of range %d"
)

const (
	pkgName = "jslice"
)

// Filter returns slice after filter.
func Filter(i interface{}, f func(interface{}) bool) ([]interface{}, error) {
	if s, err := jcast.SliceInterface(i); err != nil {
		return nil, err
	} else {
		list := make([]interface{}, 0)
		for _, v := range s {
			if f(v) {
				list = append(list, v)
			}
		}
		return list, nil
	}
}

// Insert inserts the specified element at the specified position in this list
func Insert(idx int, i, v interface{}) ([]interface{}, error) {
	if s, err := jcast.SliceInterface(i); err != nil {
		return nil, err
	} else {
		if l := len(s); l < idx {
			return nil, errorf(errorIndexOut, idx, l)
		} else if l == idx || idx < 0 {
			return append(s, v), nil
		} else {
			ns := append(s[:idx+1], s[idx:]...)
			ns[idx] = v
			return ns, nil
		}
	}
}

// InsertAll inserts all the elements in the specified collection into this list at the specified position
func InsertAll(idx int, i, v interface{}) ([]interface{}, error) {
	var err error
	var si, sv []interface{}
	if si, err = jcast.SliceInterface(i); err != nil {
		return nil, err
	}
	if sv, err = jcast.SliceInterface(v); err != nil {
		return nil, err
	}
	if l := len(si); l < idx {
		return nil, errorf(errorIndexOut, idx, l)
	} else if l == idx || idx < 0 {
		return append(si, sv...), nil
	} else {
		return append(si[:idx], append(sv, si[idx:]...)...), nil
	}
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}
