// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jevent

type Subscription interface {
	// Unsubscribe removes a Subscription from the internal list of subscriptions
	Unsubscribe()
	// Key return key
	Key() string
}

type subscription struct {
	subject *Subject
	key     string
}

func (s *subscription) Unsubscribe() {
	s.subject.Unsubscribe(s.key)
}

func (s *subscription) Key() string {
	return s.key
}
