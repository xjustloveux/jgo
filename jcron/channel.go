// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

type action int

const (
	addJob action = iota
	addSch
	stopSch
	resumeSch
	trigger
	stop
	getSch
	getSchName
	getSchCronExpression
	getSchJob
	getSchData
	getSchDesc
	getSchSts
	getSchNext
	getSchPrev
)

type channel struct {
	action action
	data   interface{}
	err    error
}
