// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

// SchInfo config schedule information
type SchInfo struct {
	// Name schedule name
	Name string
	// Cron cron expression
	Cron string
	// JobName job name
	JobName string
	// JobData job data
	JobData map[string]interface{}
	// Desc job description
	Desc string
	// Status schedule status(stop or run)
	Status string
}
