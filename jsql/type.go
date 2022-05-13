// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"strconv"
	"strings"
)

// Type sql type
type Type int

const (
	MySql Type = iota
	MSSql
	Oracle
)

// DriverName returns sql driver name
func (t Type) DriverName() string {
	switch t {
	case MySql:
		return "mysql"
	case MSSql:
		return "sqlserver"
	case Oracle:
		return "oci8"
	}
	return ""
}

// Param returns query params string of db Type
func (t Type) Param(i int) string {
	switch t {
	case MySql:
		fallthrough
	default:
		return "?"
	case MSSql:
		return fmt.Sprint("@p", strconv.FormatInt(int64(i+1), 10))
	case Oracle:
		return fmt.Sprint(":", strconv.FormatInt(int64(i), 10))
	}
}

// ParseDBType takes a string db Type and returns the db Type constant.
func ParseDBType(t string) (Type, error) {
	switch strings.ToLower(t) {
	case "mysql":
		return MySql, nil
	case "mssql":
		return MSSql, nil
	case "oracle":
		return Oracle, nil
	}
	return Unknown, errorf(errorNotValidDbType, t)
}
