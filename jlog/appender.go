// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"sync"
)

type Appender interface {
	// Call use function name to call the function
	Call(string, ...interface{}) error
	// AddLogger add *logrus.Logger
	AddLogger(*logrus.Logger)
	// ClearLogger clear logger
	ClearLogger()
}

type appender struct {
	list []*logrus.Logger
	mux  *sync.RWMutex
}

func (a appender) Call(funcName string, params ...interface{}) error {
	a.mux.RLock()
	defer func() {
		a.mux.RUnlock()
	}()
	for _, nl := range a.list {
		var e *logrus.Entry
		l := 0
		for _, param := range params {
			switch param.(type) {
			case logrus.Fields:
				e = nl.WithFields(param.(logrus.Fields))
			default:
				l++
			}
		}
		f := reflect.ValueOf(a.getFunc(nl, e)[funcName])
		if fl := f.Type().NumIn(); fl > 1 && l < fl {
			return errors(errorParamsNumOutIndex)
		}
		in := make([]reflect.Value, l)
		count := 0
		for _, param := range params {
			switch param.(type) {
			case logrus.Fields:
			default:
				if param != nil {
					in[count] = reflect.ValueOf(param)
				} else {
					in[count] = reflect.ValueOf("nil")
				}
				count++
			}
		}
		f.Call(in)
	}
	return nil
}

func (a *appender) AddLogger(l *logrus.Logger) {
	a.mux.Lock()
	defer func() {
		a.mux.Unlock()
	}()
	a.list = append(a.list, l)
}

func (a *appender) ClearLogger() {
	a.mux.Lock()
	defer func() {
		a.mux.Unlock()
	}()
	a.list = make([]*logrus.Logger, 0)
}

func (appender) getDefault() *appender {
	return &appender{
		list: make([]*logrus.Logger, 0),
		mux:  new(sync.RWMutex),
	}
}

func (a *appender) getFunc(l *logrus.Logger, e *logrus.Entry) map[string]interface{} {
	if e == nil {
		return map[string]interface{}{
			"Log":     l.Log,
			"Print":   l.Print,
			"Panic":   l.Panic,
			"Fatal":   l.Fatal,
			"Error":   l.Error,
			"Warn":    l.Warn,
			"Info":    l.Info,
			"Debug":   l.Debug,
			"Trace":   l.Trace,
			"Logf":    l.Logf,
			"Printf":  l.Printf,
			"Panicf":  l.Panicf,
			"Fatalf":  l.Fatalf,
			"Errorf":  l.Errorf,
			"Warnf":   l.Warnf,
			"Infof":   l.Infof,
			"Debugf":  l.Debugf,
			"Tracef":  l.Tracef,
			"LogFn":   l.LogFn,
			"PrintFn": l.PrintFn,
			"PanicFn": l.PanicFn,
			"FatalFn": l.FatalFn,
			"ErrorFn": l.ErrorFn,
			"WarnFn":  l.WarnFn,
			"InfoFn":  l.InfoFn,
			"DebugFn": l.DebugFn,
			"TraceFn": l.TraceFn,
			"Logln":   l.Logln,
			"Println": l.Println,
			"Panicln": l.Panicln,
			"Fatalln": l.Fatalln,
			"Errorln": l.Errorln,
			"Warnln":  l.Warnln,
			"Infoln":  l.Infoln,
			"Debugln": l.Debugln,
			"Traceln": l.Traceln,
		}
	} else {
		return map[string]interface{}{
			"Log":     e.Log,
			"Print":   e.Print,
			"Panic":   e.Panic,
			"Fatal":   e.Fatal,
			"Error":   e.Error,
			"Warn":    e.Warn,
			"Info":    e.Info,
			"Debug":   e.Debug,
			"Trace":   e.Trace,
			"Logf":    e.Logf,
			"Printf":  e.Printf,
			"Panicf":  e.Panicf,
			"Fatalf":  e.Fatalf,
			"Errorf":  e.Errorf,
			"Warnf":   e.Warnf,
			"Infof":   e.Infof,
			"Debugf":  e.Debugf,
			"Tracef":  e.Tracef,
			"LogFn":   l.LogFn,
			"PrintFn": l.PrintFn,
			"PanicFn": l.PanicFn,
			"FatalFn": l.FatalFn,
			"ErrorFn": l.ErrorFn,
			"WarnFn":  l.WarnFn,
			"InfoFn":  l.InfoFn,
			"DebugFn": l.DebugFn,
			"TraceFn": l.TraceFn,
			"Logln":   e.Logln,
			"Println": e.Println,
			"Panicln": e.Panicln,
			"Fatalln": e.Fatalln,
			"Errorln": e.Errorln,
			"Warnln":  e.Warnln,
			"Infoln":  e.Infoln,
			"Debugln": e.Debugln,
			"Traceln": e.Traceln,
		}
	}
}
