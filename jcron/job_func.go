// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

type jobFunc func(map[string]interface{})

func (j jobFunc) Run(m map[string]interface{}) {
	j(m)
}
