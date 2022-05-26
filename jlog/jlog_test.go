// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/xjustloveux/jgo/jfile"
	"testing"
)

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
	if err := Init(); err != nil {
		t.Error(err)
		return
	}
	Errorln("test")
	inPath := GetParam("path")
	outPath := "../log"
	assert.Equal(t, inPath, outPath, fmt.Sprintf("%v != %v", outPath, inPath))
	GetAppender()
	if _, err := NewAppender(Default); err == nil {
		t.Error(fmt.Sprint(testErr, " NewAppender must be return error"))
	}
	if ta, err := NewAppender("jlog_test.go"); err != nil {
		t.Error(err)
	} else {
		ta.ClearLogger()
		l := GetLogger(console)
		ta.AddLogger(l)
		ta2 := GetAppender()
		//assert.Equal(t, da, ta, fmt.Sprintf("%v != %v", ta, da))
		assert.Equal(t, ta2, ta, fmt.Sprintf("%v != %v", ta, ta2))
	}
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
}

func logFn() []interface{} {
	return []interface{}{"test"}
}
