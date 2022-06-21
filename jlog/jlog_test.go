// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jtime"
	"os"
	"sync"
	"testing"
	"time"
)

type testHandler struct{}

func (h *testHandler) Handle(event *Event) {}

func TestLog(t *testing.T) {
	testErr := "TEST ERROR:"
	SetFormat(jfile.Json)
	SetEnvFileName("")
	inEnvKey := ""
	SetEnvKey(inEnvKey)
	outEnvKey := EnvKey()
	assert.Equal(t, inEnvKey, outEnvKey, fmt.Sprintf("%v != %v", outEnvKey, inEnvKey))
	inEnvVal := ""
	SetEnvVal(inEnvVal)
	outEnvVal := EnvVal()
	assert.Equal(t, inEnvVal, outEnvVal, fmt.Sprintf("%v != %v", outEnvVal, inEnvVal))
	DisableEnv()
	EnableEnv()
	SetLogFunc(func(i ...interface{}) {})
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("../files/test-jlog-error.json")
	if err := Init(); err == nil {
		t.Error(fmt.Sprint(testErr, " Init must be return error"))
	}
	SetFileName("../files/test-jlog.json")
	AddWriter("test", os.Stdout)
	AddHandler("test", &testHandler{})
	o := Output{}.getDefault()
	o.MaxAgeDuration = "Unknown"
	if _, err := NewRotateLogs("", o, make(map[string]string)); err == nil {
		t.Error(fmt.Sprint(testErr, " NewRotateLogs must be return error"))
	}
	o.MaxAgeDuration = "Day"
	o.RotationSizeUnit = "Unknown"
	if _, err := NewRotateLogs("", o, make(map[string]string)); err == nil {
		t.Error(fmt.Sprint(testErr, " NewRotateLogs must be return error"))
	}
	o.RotationSizeUnit = "MB"
	o.RotationTimeDuration = "Unknown"
	if _, err := NewRotateLogs("", o, make(map[string]string)); err == nil {
		t.Error(fmt.Sprint(testErr, " NewRotateLogs must be return error"))
	}
	o.RotationTimeDuration = "Hour"
	o.Handler = "test"
	if _, err := NewRotateLogs("", o, make(map[string]string)); err != nil {
		t.Error(err)
	}
	o.Handler = ""
	o.Clock = "ERROR"
	if _, err := NewRotateLogs("", o, make(map[string]string)); err == nil {
		t.Error(fmt.Sprint(testErr, " NewRotateLogs must be return error"))
	}
	o.Clock = "Local"
	if err := Init(); err != nil {
		t.Error(err)
		return
	}
	Errorln("test")
	inPath := GetParam("path")
	outPath := "../log"
	assert.Equal(t, inPath, outPath, fmt.Sprintf("%v != %v", outPath, inPath))
	Log(InfoLevel, "test")
	Print("test")
	Error("test")
	Warn("test")
	Info("test")
	Debug("test")
	Trace("test")
	Logf(InfoLevel, "%v", "test")
	Printf("%v", "test")
	Errorf("%v", "test")
	Warnf("%v", "test")
	Infof("%v", "test")
	Debugf("%v", "test")
	Tracef("%v", "test")
	LogFn(InfoLevel, logFn)
	PrintFn(logFn)
	ErrorFn(logFn)
	WarnFn(logFn)
	InfoFn(logFn)
	DebugFn(logFn)
	TraceFn(logFn)
	Logln(InfoLevel, "test")
	Println("test")
	Errorln("test")
	Warnln("test")
	Infoln("test")
	Debugln("test")
	Traceln("test")
	Log(InfoLevel, Fields{"test": "test"}, "test", nil)
	Log(InfoLevel, logrus.Fields{"test": "test"}, "test", nil)
	Printf("test")
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		i := 0
		for i < 10 {
			Infoln("thread1")
			i++
			<-time.After(jtime.Millisecond * 123)
		}
	}()
	wg.Add(1)
	go func() {
		defer func() {
			wg.Done()
		}()
		i := 0
		for i < 10 {
			Infoln("thread2")
			i++
			<-time.After(jtime.Millisecond * 234)
		}
	}()
	wg.Wait()
}

func logFn() []interface{} {
	return []interface{}{"test"}
}
