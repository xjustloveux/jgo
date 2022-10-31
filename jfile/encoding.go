// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"sync"
)

var (
	codec = make(map[string]Codec)
	mux   sync.RWMutex
)

// RegisterCodec registers Codec
func RegisterCodec(format string, c Codec) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	codec[format] = c
}

// GetCodec returns Codec
func GetCodec(format string) (Codec, error) {
	mux.RLock()
	defer func() {
		mux.RUnlock()
	}()
	if c, ok := codec[format]; ok {
		return c, nil
	} else {
		return nil, errorStr(errorNotFoundCodec)
	}
}

// Encode encode
func Encode(format string, m map[string]interface{}) ([]byte, error) {
	if c, err := GetCodec(format); err != nil {
		return nil, err
	} else {
		return c.Encode(m)
	}
}

// Decode decode
func Decode(format string, b []byte, m map[string]interface{}) error {
	if c, err := GetCodec(format); err != nil {
		return err
	} else {
		return c.Decode(b, m)
	}
}
