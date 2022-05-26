// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import "strings"

type tag string

const (
	tagDao     tag = "dao"
	tagText    tag = "text"
	tagSelect  tag = "select"
	tagInsert  tag = "insert"
	tagUpdate  tag = "update"
	tagDelete  tag = "delete"
	tagOther   tag = "other"
	tagIf      tag = "if"
	tagForeach tag = "foreach"
	tagOrderBy tag = "orderBy"
	tagUnknown tag = "Unknown"
)

func (t tag) String() string {
	return string(t)
}

func ParseTag(str string) tag {
	switch strings.ToLower(str) {
	case "dao":
		return tagDao
	case "text":
		return tagText
	case "select":
		return tagSelect
	case "insert":
		return tagInsert
	case "update":
		return tagUpdate
	case "delete":
		return tagDelete
	case "other":
		return tagOther
	case "if":
		return tagIf
	case "foreach":
		return tagForeach
	case "orderby":
		return tagOrderBy
	default:
		return tagUnknown
	}
}
