// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	Equal Operators = iota
	NotEqual
	In
	NotIn
	Between
	NotBetween
	IsNull
	IsNotNull
	Like
	SLike
	ELike
	Greater
	GreaterThanOrEqual
	Less
	LessThanOrEqual
)

// Operators clause Operators
type Operators int

// String returns Operators string
func (o Operators) String() string {
	switch o {
	case Equal:
		return "Equal"
	case NotEqual:
		return "NotEqual"
	case In:
		return "In"
	case NotIn:
		return "NotIn"
	case Between:
		return "Between"
	case NotBetween:
		return "NotBetween"
	case IsNull:
		return "IsNull"
	case IsNotNull:
		return "IsNotNull"
	case Like:
		return "Like"
	case SLike:
		return "SLike"
	case ELike:
		return "ELike"
	case Greater:
		return "Greater"
	case GreaterThanOrEqual:
		return "GreaterThanOrEqual"
	case Less:
		return "Less"
	case LessThanOrEqual:
		return "LessThanOrEqual"
	default:
		return "Unknown"
	}
}

// ParseOperators takes a string Operators and returns the Operators constant.
func ParseOperators(o string) (Operators, error) {
	switch strings.ToLower(o) {
	case "equal":
		return Equal, nil
	case "notequal":
		return NotEqual, nil
	case "in":
		return In, nil
	case "notin":
		return NotIn, nil
	case "between":
		return Between, nil
	case "notbetween":
		return NotBetween, nil
	case "isnull":
		return IsNull, nil
	case "isnotnull":
		return IsNotNull, nil
	case "like":
		return Like, nil
	case "slike":
		return SLike, nil
	case "elike":
		return ELike, nil
	case "greater":
		return Greater, nil
	case "greaterthanorequal":
		return GreaterThanOrEqual, nil
	case "less":
		return Less, nil
	case "lessthanorequal":
		return LessThanOrEqual, nil
	}
	return Unknown, errorf(errorNotValidOperators, o)
}

func (o Operators) getClauseAndParams(t Type, val interface{}, params []interface{}) (string, []interface{}, error) {
	switch o {
	case Equal:
		return fmt.Sprint(" = ", t.Param(len(params))), append(params, val), nil
	case NotEqual:
		return fmt.Sprint(" != ", t.Param(len(params))), append(params, val), nil
	case In:
		if vs, ok := val.([]interface{}); ok {
			if len(vs) > 0 {
				var ps string
				np := params
				for i, v := range vs {
					if i == 0 {
						ps = fmt.Sprint(ps, t.Param(len(np)))
						np = append(np, v)
					} else {
						ps = fmt.Sprint(ps, ", ", t.Param(len(np)))
						np = append(np, v)
					}
				}
				return fmt.Sprint(" IN (", ps, ")"), np, nil
			} else {
				return "", nil, errorf(errorOprValLenZero, o.String())
			}
		} else {
			return "", nil, errorf(errorOprValTypeNotInterfaceSlice, o.String(), reflect.TypeOf(val))
		}
	case NotIn:
		if vs, ok := val.([]interface{}); ok {
			if len(vs) > 0 {
				var ps string
				np := params
				for i, v := range vs {
					if i == 0 {
						ps = fmt.Sprint(ps, t.Param(len(np)))
						np = append(np, v)
					} else {
						ps = fmt.Sprint(ps, ", ", t.Param(len(np)))
						np = append(np, v)
					}
				}
				return fmt.Sprint(" NOT IN (", ps, ")"), np, nil
			} else {
				return "", nil, errorf(errorOprValLenZero, o.String())
			}
		} else {
			return "", nil, errorf(errorOprValTypeNotInterfaceSlice, o.String(), reflect.TypeOf(val))
		}
	case Between:
		if vs, ok := val.([]interface{}); ok {
			if len(vs) == 2 {
				var ps string
				np := params
				ps = fmt.Sprint(ps, t.Param(len(np)))
				np = append(np, vs[0])
				ps = fmt.Sprint(ps, " AND ", t.Param(len(np)))
				np = append(np, vs[1])
				return fmt.Sprint(" BETWEEN ", ps), np, nil
			} else {
				return "", nil, errorf(errorOprValLenNot2, o.String())
			}
		} else {
			return "", nil, errorf(errorOprValTypeNotInterfaceSlice, o.String(), reflect.TypeOf(val))
		}
	case NotBetween:
		if vs, ok := val.([]interface{}); ok {
			if len(vs) == 2 {
				var ps string
				np := params
				ps = fmt.Sprint(ps, t.Param(len(np)))
				np = append(np, vs[0])
				ps = fmt.Sprint(ps, " AND ", t.Param(len(np)))
				np = append(np, vs[1])
				return fmt.Sprint(" NOT BETWEEN ", ps), np, nil
			} else {
				return "", nil, errorf(errorOprValLenNot2, o.String())
			}
		} else {
			return "", nil, errorf(errorOprValTypeNotInterfaceSlice, o.String(), reflect.TypeOf(val))
		}
	case IsNull:
		return " IS NULL", params, nil
	case IsNotNull:
		return " IS NOT NULL", params, nil
	case Like:
		if s, ok := val.(string); ok {
			return fmt.Sprint(" LIKE ", t.Param(len(params))), append(params, fmt.Sprint("%", s, "%")), nil
		} else {
			return "", nil, errorf(errorOprValTypeNotString, o.String(), reflect.TypeOf(val))
		}
	case SLike:
		if s, ok := val.(string); ok {
			return fmt.Sprint(" LIKE ", t.Param(len(params))), append(params, fmt.Sprint(s, "%")), nil
		} else {
			return "", nil, errorf(errorOprValTypeNotString, o.String(), reflect.TypeOf(val))
		}
	case ELike:
		if s, ok := val.(string); ok {
			return fmt.Sprint(" LIKE ", t.Param(len(params))), append(params, fmt.Sprint("%", s)), nil
		} else {
			return "", nil, errorf(errorOprValTypeNotString, o.String(), reflect.TypeOf(val))
		}
	case Greater:
		return fmt.Sprint(" > ", t.Param(len(params))), append(params, val), nil
	case GreaterThanOrEqual:
		return fmt.Sprint(" >= ", t.Param(len(params))), append(params, val), nil
	case Less:
		return fmt.Sprint(" < ", t.Param(len(params))), append(params, val), nil
	case LessThanOrEqual:
		return fmt.Sprint(" <= ", t.Param(len(params))), append(params, val), nil
	default:
		return "", nil, errors(errorUnknownOpr)
	}
}
