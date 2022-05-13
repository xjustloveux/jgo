// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

type sortSch []*schedule

func (s sortSch) Len() int {
	return len(s)
}

func (s sortSch) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortSch) Less(i, j int) bool {
	if s[i].next.IsZero() {
		return false
	}
	if s[j].next.IsZero() {
		return true
	}
	return s[i].next.Before(s[j].next)
}
