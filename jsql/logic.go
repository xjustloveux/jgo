// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

const (
	And Logic = iota
	Or
)

// Logic clause Logic
type Logic int

// String returns Logic string
func (l Logic) String() string {
	switch l {
	case And:
		return "AND"
	case Or:
		return "OR"
	}
	return ""
}
