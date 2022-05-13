// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

type SchInfo struct {
	Name    string
	Cron    string
	JobName string
	JobData map[string]interface{}
	Desc    string
}
