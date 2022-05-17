// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jcron

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jtime"
	"testing"
	"time"
)

func TestSortSch(t *testing.T) {
	now := time.Now()
	s1 := schedule{next: now}
	s2 := schedule{next: now.Add(jtime.Second)}
	list := sortSch{&s1, &s2}
	inLen := list.Len()
	outLen := 2
	assert.Equal(t, outLen, inLen, fmt.Sprintf("%v != %v", inLen, outLen))
	list.Swap(0, 1)
	assert.Equal(t, &s2, list[0], fmt.Sprintf("%v != %v", list[0], &s2))
	assert.Equal(t, &s1, list[1], fmt.Sprintf("%v != %v", list[1], &s1))
	less := list.Less(0, 1)
	assert.Equal(t, false, less, fmt.Sprintf("%v != %v", less, false))
	list[0].next = time.Time{}
	less = list.Less(0, 1)
	assert.Equal(t, false, less, fmt.Sprintf("%v != %v", less, false))
	list[0].next = now
	list[1].next = time.Time{}
	less = list.Less(0, 1)
	assert.Equal(t, true, less, fmt.Sprintf("%v != %v", less, true))
}
