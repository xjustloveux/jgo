// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jconf

type jError string

func (e jError) Error() string {
	return string(e)
}
