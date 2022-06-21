// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"github.com/sirupsen/logrus"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

type logger struct {
	logs map[string]*logrus.Logger
	mux  *sync.RWMutex
}

func (logger) getDefault() *logger {
	return &logger{
		logs: make(map[string]*logrus.Logger),
		mux:  new(sync.RWMutex),
	}
}

// LogrusLogger returns *logrus.Logger with appender name
func (l *logger) LogrusLogger(name string) *logrus.Logger {
	l.mux.RLock()
	defer func() {
		l.mux.RUnlock()
	}()
	return l.logs[name]
}

func (l *logger) Call(funcName string, params ...interface{}) error {
	l.mux.RLock()
	defer func() {
		l.mux.RUnlock()
	}()
	for _, nl := range l.logs {
		var e *logrus.Entry
		pl := 0
		for _, param := range params {
			switch param.(type) {
			case Fields:
				e = nl.WithFields(param.(Fields).toLogrusFields())
			case logrus.Fields:
				e = nl.WithFields(param.(logrus.Fields))
			default:
				pl++
			}
		}
		f := reflect.ValueOf(l.getFunc(nl, e)[funcName])
		if isFormat, err := regexp.MatchString("f$", funcName); err != nil {
			return err
		} else {
			if isFormat {
				if fl := f.Type().NumIn(); fl > 1 && pl < fl-1 {
					return errors(errorParamsNumOutIndex)
				}
			} else {
				if fl := f.Type().NumIn(); fl > 1 && pl < fl {
					return errors(errorParamsNumOutIndex)
				}
			}
		}
		in := make([]reflect.Value, pl)
		count := 0
		for _, param := range params {
			switch param.(type) {
			case Fields:
			case logrus.Fields:
			default:
				if param != nil {
					in[count] = reflect.ValueOf(param)
				} else {
					in[count] = reflect.ValueOf("<nil>")
				}
				count++
			}
		}
		f.Call(in)
	}
	return nil
}

func (l *logger) AddLogger(name string, log *logrus.Logger) {
	l.mux.Lock()
	defer func() {
		l.mux.Unlock()
	}()
	l.logs[name] = log
}

func (l *logger) ClearLogger() {
	l.mux.Lock()
	defer func() {
		l.mux.Unlock()
	}()
	l.logs = make(map[string]*logrus.Logger)
}

func (l *logger) addLogger(name string, a *appender, param map[string]string) error {
	log := logrus.New()
	switch strings.ToUpper(a.Formatter.Type) {
	case "TEXT":
		log.SetFormatter(a.Formatter.Text)
	case "JSON":
		log.SetFormatter(a.Formatter.Json)
	}
	log.SetLevel(a.getLogrusLevel())
	if a.Output.Name != "" && writer[a.Output.Name] != nil {
		log.SetOutput(writer[a.Output.Name])
	} else if w, err := NewRotateLogs(name, a.Output, param); err == nil {
		log.SetOutput(w)
	} else {
		return err
	}
	l.logs[name] = log
	return nil
}

func (l *logger) getFunc(log *logrus.Logger, e *logrus.Entry) map[string]interface{} {
	if e == nil {
		return map[string]interface{}{
			"Log":     log.Log,
			"Print":   log.Print,
			"Panic":   log.Panic,
			"Fatal":   log.Fatal,
			"Error":   log.Error,
			"Warn":    log.Warn,
			"Info":    log.Info,
			"Debug":   log.Debug,
			"Trace":   log.Trace,
			"Logf":    log.Logf,
			"Printf":  log.Printf,
			"Panicf":  log.Panicf,
			"Fatalf":  log.Fatalf,
			"Errorf":  log.Errorf,
			"Warnf":   log.Warnf,
			"Infof":   log.Infof,
			"Debugf":  log.Debugf,
			"Tracef":  log.Tracef,
			"LogFn":   log.LogFn,
			"PrintFn": log.PrintFn,
			"PanicFn": log.PanicFn,
			"FatalFn": log.FatalFn,
			"ErrorFn": log.ErrorFn,
			"WarnFn":  log.WarnFn,
			"InfoFn":  log.InfoFn,
			"DebugFn": log.DebugFn,
			"TraceFn": log.TraceFn,
			"Logln":   log.Logln,
			"Println": log.Println,
			"Panicln": log.Panicln,
			"Fatalln": log.Fatalln,
			"Errorln": log.Errorln,
			"Warnln":  log.Warnln,
			"Infoln":  log.Infoln,
			"Debugln": log.Debugln,
			"Traceln": log.Traceln,
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
			"LogFn":   log.LogFn,
			"PrintFn": log.PrintFn,
			"PanicFn": log.PanicFn,
			"FatalFn": log.FatalFn,
			"ErrorFn": log.ErrorFn,
			"WarnFn":  log.WarnFn,
			"InfoFn":  log.InfoFn,
			"DebugFn": log.DebugFn,
			"TraceFn": log.TraceFn,
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
