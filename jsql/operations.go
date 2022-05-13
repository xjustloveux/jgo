// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

// Operations sql operations
type Operations int

const (
	Select Operations = iota
	Insert
	Update
	Delete
	Other
)

// String returns Operations string
func (o Operations) String() string {
	switch o {
	case Select:
		return "Select"
	case Insert:
		return "Insert"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case Other:
		return "Other"
	}
	return "Unknown"
}
