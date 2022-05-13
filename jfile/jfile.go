// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"encoding/json"
	"fmt"
	"github.com/xjustloveux/jgo/jcast"
	"io"
	"os"
	"reflect"
)

const (
	errorNotValidSizeUnit = jError("not a valid size unit %q")
	errorNotFoundCodec    = jError("codec not fount")
)

const (
	pkgName = "jfile"
)

func init() {
	RegisterCodec(Json.String(), jsonCodec{})
}

// Load file to byte array
func Load(name string) (bytes []byte, err error) {
	var file *os.File
	if file, err = os.Open(name); err != nil {
		return nil, err
	}
	defer func() {
		if e := file.Close(); e != nil {
			err = e
		}
	}()
	return io.ReadAll(file)
}

// Convert map[string]interface{} convert to v
func Convert(m map[string]interface{}, v interface{}) error {
	var err error
	for mk, mv := range m {
		if reflect.TypeOf(mv).Kind() == reflect.Map {
			if m[mk], err = toStringMap(mv); err != nil {
				return err
			}
		}
	}
	var b []byte
	if b, err = Encode(Json.String(), m); err != nil {
		return err
	}
	return json.Unmarshal(b, &v)
}

func errorf(e jError, args ...interface{}) error {
	return fmt.Errorf(fmt.Sprint(pkgName, ": ", e.Error()), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}

func toStringMap(i interface{}) (map[string]interface{}, error) {
	var err error
	var m map[string]interface{}
	if m, err = jcast.StringMapInterface(i); err != nil {
		return nil, err
	} else {
		for k, v := range m {
			if reflect.TypeOf(v).Kind() == reflect.Map {
				if m[k], err = toStringMap(v); err != nil {
					return nil, err
				}
			}
		}
	}
	return m, nil
}
