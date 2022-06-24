// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jevent

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"sync"
)

type Subject struct {
	mux   *sync.RWMutex
	event map[string]Event
}

// Subscribe subscribe Event
func (s *Subject) Subscribe(e Event) Subscription {
	if e == nil {
		return nil
	}
	s.mux.Lock()
	defer func() {
		s.mux.Unlock()
	}()
	if key, err := s.rand(); err != nil {
		return nil
	} else {
		s.event[key] = e
		return &subscription{s, key}
	}
}

// Unsubscribe removes a Subscription from the internal list of subscriptions with key
func (s *Subject) Unsubscribe(key string) {
	s.mux.Lock()
	defer func() {
		s.mux.Unlock()
	}()
	if _, ok := s.event[key]; ok {
		delete(s.event, key)
	}
}

func (s *Subject) Next(args ...interface{}) {
	s.mux.RLock()
	defer func() {
		s.mux.RUnlock()
	}()
	for _, e := range s.event {
		go e(args...)
	}
}

func (s *Subject) rand() (string, error) {
	r := ""
	for r == "" {
		var b16 [16]byte
		var b32 [32]byte
		if _, err := io.ReadFull(rand.Reader, b16[:]); err != nil {
			return "", err
		}
		hex.Encode(b32[:], b16[:])
		r = string(b32[:])
		if _, ok := s.event[r]; ok {
			r = ""
		}
	}
	return r, nil
}
