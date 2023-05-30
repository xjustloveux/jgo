// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"fmt"
	"github.com/xjustloveux/jgo/jslice"
	"github.com/xjustloveux/jgo/jtime"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

type logFile struct {
	name string
	mux  *sync.RWMutex
	file *os.File
}

func (l *logFile) open(lock bool) error {
	if lock {
		l.mux.Lock()
		defer func() {
			l.mux.Unlock()
		}()
	}
	if l.file == nil {
		dir := filepath.Dir(l.name)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return err
		}
		if fi, err := os.OpenFile(l.name, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644); err != nil {
			return err
		} else {
			l.file = fi
		}
	}
	return nil
}

func (l *logFile) close(lock bool) error {
	if lock {
		l.mux.Lock()
		defer func() {
			l.mux.Unlock()
		}()
	}
	if l.file != nil {
		if err := l.file.Close(); err != nil {
			return err
		}
		l.file = nil
	}
	return nil
}

func (l *logFile) rotation(lock bool) error {
	if lock {
		l.mux.Lock()
		defer func() {
			l.mux.Unlock()
		}()
	}
	count := 1
	var name string
	for {
		name = fmt.Sprintf("%s.%d", l.name, count)
		if _, err := os.Stat(name); err != nil {
			break
		}
		count++
	}
	if err := l.close(false); err != nil {
		return err
	}
	if err := os.Rename(l.name, name); err != nil {
		return err
	}
	if err := l.open(false); err != nil {
		return err
	}
	return nil
}

func (l *logFile) remove(lock bool, sn string, age time.Duration, count int, current, currentLink string) error {
	if lock {
		l.mux.Lock()
		defer func() {
			l.mux.Unlock()
		}()
	}
	if age > 0 || count > 0 {
		format := jtime.FormatList()
		for _, v := range format {
			sn = strings.ReplaceAll(sn, fmt.Sprint("%", v), "*")
		}
		sn = fmt.Sprint(sn, "*")
		if matches, err := filepath.Glob(sn); err != nil {
			return err
		} else {
			matches = l.sortByModTime(matches)
			cutoff := time.Now().Add(-1 * age)
			remove := make([]string, 0)
			for _, name := range matches {
				if name == current || name == currentLink {
					continue
				}
				if strings.HasSuffix(name, "_symlink") {
					continue
				}
				var fi os.FileInfo
				if fi, err = os.Stat(name); err != nil {
					continue
				}
				var fl os.FileInfo
				if fl, err = os.Lstat(name); err != nil {
					continue
				}
				if fl.Mode()&os.ModeSymlink == os.ModeSymlink {
					continue
				}
				if age > 0 && cutoff.After(fi.ModTime()) {
					remove = append(remove, name)
					continue
				}
				if count > 0 && len(matches)-len(remove) > count {
					remove = append(remove, name)
					continue
				}
			}
			if len(remove) > 0 {
				for _, path := range remove {
					if err = os.Remove(path); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (l *logFile) sortByModTime(list []string) []string {
	type tempInfo struct {
		time time.Time
		name string
	}
	nl := make([]interface{}, 0)
	for _, name := range list {
		if fi, err := os.Stat(name); err != nil {
			continue
		} else {
			idx := 0
			for i, f := range nl {
				if f.(tempInfo).time.After(fi.ModTime()) {
					break
				} else {
					idx = i + 1
				}
			}
			if nl, err = jslice.Insert(idx, nl, tempInfo{fi.ModTime(), name}); err != nil {
				continue
			}
		}
	}
	newList := make([]string, len(nl))
	for i, ti := range nl {
		newList[i] = ti.(tempInfo).name
	}
	return newList
}

func (l *logFile) write(lock bool, sn string, size int64, age time.Duration, count int, current, currentLink string, p []byte) (n int, err error) {
	if lock {
		l.mux.Lock()
		defer func() {
			l.mux.Unlock()
		}()
	}
	if size > 0 {
		var fi os.FileInfo
		if fi, err = os.Stat(l.name); err != nil {
			return 0, err
		} else if size <= fi.Size() {
			if err = l.rotation(false); err != nil {
				return 0, err
			}
		}
	}
	if err = l.remove(false, sn, age, count, current, currentLink); err != nil {
		return 0, err
	}
	if l.file == nil {
		if err = l.open(false); err != nil {
			return 0, err
		}
	}
	return l.file.Write(p)
}
