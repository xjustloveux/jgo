// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jsql

import "fmt"

type Param struct {
	Logic       Logic
	Col         string
	Val         interface{}
	Opr         Operators
	ParamsLogic Logic
	Params      []*Param
}

// AddParam add Param to this Param.Params array
func (p *Param) AddParam(param *Param) {
	p.Params = append(p.Params, param)
}

func (p *Param) getClauseAndParams(t Type, params []interface{}) (string, []interface{}, error) {
	clause := ""
	if p.Col != "" {
		if opr, pm, err := p.Opr.getClauseAndParams(t, p.Val, params); err != nil {
			return "", nil, err
		} else {
			clause = fmt.Sprint(" ", p.Logic.String(), " ", p.Col, opr)
			params = pm
		}
	}
	if len(p.Params) > 0 {
		clause = fmt.Sprint(clause, " ", p.ParamsLogic.String(), " (1 = 1")
		for _, param := range p.Params {
			if opr, pm, err := param.getClauseAndParams(t, params); err != nil {
				return "", nil, err
			} else {
				clause = fmt.Sprint(clause, opr)
				params = pm
			}
		}
		clause = fmt.Sprint(clause, ")")
	}
	return clause, params, nil
}
