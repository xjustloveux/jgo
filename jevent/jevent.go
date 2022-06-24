// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jevent

import "sync"

// New create new *Subject
func New() *Subject {
	return &Subject{
		mux:   new(sync.RWMutex),
		event: make(map[string]Event),
	}
}
