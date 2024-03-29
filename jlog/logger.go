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
	"time"
)

type logger struct {
	logs map[string]LogrusLogger
	mux  *sync.RWMutex
}

func (*logger) getDefault() *logger {
	return &logger{
		logs: make(map[string]LogrusLogger),
		mux:  new(sync.RWMutex),
	}
}

// LogrusLogger returns *logrus.Logger with appender name
func (l *logger) LogrusLogger(name string) *logrus.Logger {
	l.mux.RLock()
	defer func() {
		l.mux.RUnlock()
	}()
	return l.logs[name].Log
}

func (l *logger) Call(funcName, reportCaller string, params ...interface{}) error {
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
				fields := param.(Fields).toLogrusFields()
				if nl.ReportCaller {
					fields["file"] = reportCaller
				}
				e = nl.Log.WithFields(fields)
			case logrus.Fields:
				fields := param.(logrus.Fields)
				if nl.ReportCaller {
					fields["file"] = reportCaller
				}
				e = nl.Log.WithFields(fields)
			default:
				pl++
			}
		}
		if e == nil && nl.ReportCaller {
			e = nl.Log.WithFields(logrus.Fields{"file": reportCaller})
		}
		f := reflect.ValueOf(l.getFunc(nl.Log, e)[funcName])
		if isFormat, err := regexp.MatchString("f$", funcName); err != nil {
			return err
		} else {
			if isFormat {
				if fl := f.Type().NumIn(); fl > 1 && pl < fl-1 {
					return errorStr(errorParamsNumOutIndex)
				}
			} else {
				if fl := f.Type().NumIn(); fl > 1 && pl < fl {
					return errorStr(errorParamsNumOutIndex)
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

func (l *logger) AddLogger(name string, log LogrusLogger) {
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
	l.logs = make(map[string]LogrusLogger)
}

func (l *logger) addLogger(an string, pn string, a *appender, param map[string]string) error {
	log := logrus.New()
	var loc *time.Location
	if len(a.Formatter.Location) > 0 {
		var err error
		if loc, err = time.LoadLocation(a.Formatter.Location); err != nil {
			return err
		}
	}
	var f logrus.Formatter
	switch strings.ToUpper(a.Formatter.Type) {
	default:
		fallthrough
	case "TEXT":
		if loc == nil {
			f = a.Formatter.Text
		} else {
			f = timeFormatter{
				loc: loc,
				log: a.Formatter.Text,
			}
		}
	case "JSON":
		if loc == nil {
			f = a.Formatter.Json
		} else {
			f = timeFormatter{
				loc: loc,
				log: a.Formatter.Json,
			}
		}
	}
	log.SetFormatter(f)
	log.SetLevel(a.getLogrusLevel())
	if a.Output.Name != "" && writer[a.Output.Name] != nil {
		log.SetOutput(writer[a.Output.Name])
	} else if w, err := NewRotateLogs(strings.Trim(pn, ".go"), a.Output, param); err == nil {
		log.SetOutput(w)
	} else {
		return err
	}
	l.logs[an] = LogrusLogger{Log: log, ReportCaller: a.ReportCaller}
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
