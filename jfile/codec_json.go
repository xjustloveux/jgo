// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

import (
	"encoding/json"
)

type jsonCodec struct{}

func (jsonCodec) Encode(m map[string]interface{}) ([]byte, error) {
	return json.MarshalIndent(m, "", "  ")
}

func (jsonCodec) Decode(b []byte, m map[string]interface{}) error {
	return json.Unmarshal(b, &m)
}
