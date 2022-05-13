// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

type Job interface {
	// Run executes the underlying function
	Run(map[string]interface{})
}

type job struct {
	name string
	j    Job
}
