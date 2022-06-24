// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jevent

import (
	"fmt"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	subject := New()
	sub1 := subject.Subscribe(func(i ...interface{}) {
		fmt.Println("sub1")
		fmt.Println(i...)
	})
	sub2 := subject.Subscribe(func(i ...interface{}) {
		fmt.Println("sub2")
		fmt.Println(i...)
	})
	fmt.Println(sub1.Key())
	fmt.Println(sub2.Key())
	subject.Next("run next")
	<-time.After(time.Second)
	sub1.Unsubscribe()
	subject.Next("run next")
	<-time.After(time.Second)
	subject.Unsubscribe(sub2.Key())
}
