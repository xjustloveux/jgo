// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jfile

type Codec interface {
	//Encode encode
	Encode(map[string]interface{}) ([]byte, error)
	//Decode decode
	Decode([]byte, map[string]interface{}) error
}
