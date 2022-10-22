// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jconf

import (
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jfile"
	"reflect"
)

const (
	errorFileNameEmpty    = jError("file name is empty")
	errorArgNil           = jError("arg is nil")
	errorValNilOfArg      = jError("the value is nil of %q")
	errorIndexOut         = jError("index %d out of bounds %d")
	errorValNotFoundOfArg = jError("the value not found of %q")
)

const (
	pkgName = "jconf"
	envKey  = "jEnv"
	dash    = "-"
)

func New() Config {
	return &config{
		format:      jfile.Json,
		fileName:    "",
		root:        "./config/",
		envFileName: "",
		data:        make(map[string]interface{}),
		env:         true,
		envKey:      "",
		envVal:      "",
	}
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}

func overwriteMap(m1, m2 map[interface{}]interface{}) interface{} {
	nm := make(map[interface{}]interface{})
	if m1 != nil {
		for k, v := range m1 {
			nm[k] = jcast.Value(v)
		}
	}
	if m2 != nil {
		for k, v := range m2 {
			if nm[k] == nil {
				nm[k] = jcast.Value(v)
			} else if v != nil {
				def := false
				if reflect.TypeOf(v).Kind() == reflect.Map && reflect.TypeOf(nm[k]).Kind() == reflect.Map {
					var err error
					var vimi, nmimi map[interface{}]interface{}
					if vimi, err = jcast.InterfaceMapInterface(v); err != nil {
						def = true
					}
					if nmimi, err = jcast.InterfaceMapInterface(nm[k]); err != nil {
						def = true
					}
					if !def {
						nm[k] = overwriteMap(nmimi, vimi)
						def = true
					}
				} else if reflect.TypeOf(v).Kind() == reflect.Slice && reflect.TypeOf(nm[k]).Kind() == reflect.Slice {
					var err error
					var vsi, nmsi []interface{}
					if vsi, err = jcast.SliceInterface(v); err != nil {
						def = true
					}
					if nmsi, err = jcast.SliceInterface(nm[k]); err != nil {
						def = true
					}
					if !def {
						nm[k] = overwriteSlice(nmsi, vsi)
						def = true
					}
				}
				if !def {
					nm[k] = jcast.Value(v)
				}
			}
		}
	}
	return nm
}

func overwriteSlice(s1, s2 []interface{}) []interface{} {
	var ns []interface{}
	if s1 != nil {
		ns = make([]interface{}, len(s1))
		for i, v := range s1 {
			ns[i] = v
		}
	} else {
		ns = make([]interface{}, 0)
	}
	if s2 != nil {
		for i, v := range s2 {
			if i < len(ns) {
				if ns[i] == nil {
					ns[i] = jcast.Value(v)
				} else if v != nil {
					def := false
					if reflect.TypeOf(v).Kind() == reflect.Map && reflect.TypeOf(ns[i]).Kind() == reflect.Map {
						var err error
						var vimi, nsimi map[interface{}]interface{}
						if vimi, err = jcast.InterfaceMapInterface(v); err != nil {
							def = true
						}
						if nsimi, err = jcast.InterfaceMapInterface(ns[i]); err != nil {
							def = true
						}
						if !def {
							ns[i] = overwriteMap(nsimi, vimi)
							def = true
						}
					} else if reflect.TypeOf(v).Kind() == reflect.Slice && reflect.TypeOf(ns[i]).Kind() == reflect.Slice {
						var err error
						var vsi, nssi []interface{}
						if vsi, err = jcast.SliceInterface(v); err != nil {
							def = true
						}
						if nssi, err = jcast.SliceInterface(ns[i]); err != nil {
							def = true
						}
						if !def {
							ns[i] = overwriteSlice(nssi, vsi)
							def = true
						}
					}
					if !def {
						ns[i] = jcast.Value(v)
					}
				}
			} else {
				ns = append(ns, v)
			}
		}
	}
	return ns
}
