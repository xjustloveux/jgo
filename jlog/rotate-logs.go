// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"fmt"
	"github.com/xjustloveux/jgo/jtime"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type RotateLogs struct {
	clock         *time.Location
	handler       Handler
	mux           *sync.RWMutex
	fileName      string
	linkName      string
	rotationTime  time.Duration
	current       string
	previous      string
	currentLink   string
	previousLink  string
	rotationSize  int64
	maxAge        time.Duration
	rotationCount int
}

func (r *RotateLogs) Write(p []byte) (n int, err error) {
	r.mux.Lock()
	defer func() {
		r.mux.Unlock()
	}()
	now := time.Now().In(r.clock)
	base := now.Truncate(r.rotationTime)
	fName := filepath.Clean(jtime.FormatString(r.fileName, base))
	if fName != r.current {
		if err = removeFile(r.current); err != nil {
			return 0, err
		}
	}
	r.previous = r.current
	r.current = fName
	var file *logFile
	if file, err = getFile(fName); err != nil {
		return 0, err
	}
	if r.linkName != "" {
		lName := filepath.Clean(jtime.FormatString(r.linkName, base))
		if r.current != r.previous || lName != r.currentLink {
			tmpLink := fmt.Sprint(r.current, "_symlink")
			linkDir := filepath.Dir(lName)
			if _, err = os.Stat(linkDir); err != nil {
				if err = os.MkdirAll(linkDir, 0755); err != nil {
					return 0, err
				}
			}
			if err = os.Symlink(r.current, tmpLink); err != nil {
				return 0, err
			}
			if err = os.Rename(tmpLink, lName); err != nil {
				return 0, err
			}
			r.previousLink = r.currentLink
			r.currentLink = lName
		}
	}
	if n, err = file.write(true, r.rotationSize, r.maxAge, r.rotationCount, r.current, r.currentLink, p); err != nil {
		return n, err
	} else {
		if r.handler != nil {
			go r.handler.Handle(&Event{
				Previous:     r.previous,
				Current:      r.current,
				PreviousLink: r.previousLink,
				CurrentLink:  r.currentLink,
			})
		}
		return n, err
	}
}
