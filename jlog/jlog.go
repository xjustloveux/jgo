// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"fmt"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jruntime"
	"github.com/xjustloveux/jgo/jtime"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	errorParamsNumOutIndex = jError("the number of params is out of index")
	errorAppenderNil       = jError("appender is nil")
	errorAppenderExists    = jError("appender already exists for this program name")
)

const (
	pkgName  = "jlog"
	fileName = "log.json"
	console  = "console"
	Default  = "default"
	pkgKey   = "pkg:"
)

var (
	conf    = jconf.New()
	mux     = new(sync.RWMutex)
	cd      *configData
	logFunc func(...interface{})
	apMap   map[string]*logrus.Logger
	ap      map[string]*appender
)

func init() {
	SetFileName(fileName)
}

// SetFormat set config format
func SetFormat(f jfile.Format) {
	conf.SetFormat(f)
}

// SetFileName set config file name
func SetFileName(name string) {
	conf.SetFileName(name)
}

// SetEnvFileName set config env file name
func SetEnvFileName(name string) {
	conf.SetEnvFileName(name)
}

// EnvKey returns env key
func EnvKey() string {
	return conf.EnvKey()
}

// SetEnvKey set env key
func SetEnvKey(key string) {
	conf.SetEnvKey(key)
}

// EnvVal returns env value
func EnvVal() string {
	return conf.EnvVal()
}

// SetEnvVal set env value
func SetEnvVal(val string) {
	conf.SetEnvVal(val)
}

// EnableEnv enable env
func EnableEnv() {
	conf.EnableEnv()
}

// DisableEnv disable env
func DisableEnv() {
	conf.DisableEnv()
}

// SetLogFunc set fmt.Println log function
func SetLogFunc(f func(...interface{})) {
	logFunc = f
}

// Init initialize
func Init() error {
	if err := conf.Load(); err != nil {
		return err
	}
	cd = &configData{Debug: false}
	if err := conf.Convert(cd); err != nil {
		return err
	}
	if err := createAppender(); err != nil {
		return err
	}
	return nil
}

// GetParam returns conf.Params with key
func GetParam(key string) interface{} {
	return cd.Params[key]
}

// GetAppender returns appender for logs with program file name.
// if not found appender then default return console appender.
func GetAppender(args ...string) Appender {
	mux.RLock()
	defer func() {
		mux.RUnlock()
	}()
	if len(args) > 0 {
		for _, k := range args {
			if a := ap[k]; a != nil {
				return a
			}
		}
	} else {
		if a := ap[jruntime.GetCallerProgramName()]; a != nil {
			return a
		}
		if a := ap[fmt.Sprint(pkgKey, jruntime.GetCallerPkgName())]; a != nil {
			return a
		}
	}
	return ap[Default]
}

// NewAppender new appender of program name
func NewAppender(name string) (Appender, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if ap[name] != nil {
		return nil, errors(errorAppenderExists)
	}
	ap[name] = appender{}.getDefault()
	return ap[name], nil
}

// GetLogger returns *logrus.Logger with appender name
func GetLogger(name string) *logrus.Logger {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	return apMap[name]
}

// Log log
func Log(level Level, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = logrus.Level(level)
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Print print
func Print(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Panic print panic level message
func Panic(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Fatal print fatal level message
func Fatal(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Error print error level message
func Error(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Warn print warn level message
func Warn(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Info print info level message
func Info(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Debug print debug level message
func Debug(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Trace print trace level message
func Trace(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Logf log
func Logf(level Level, format string, args ...interface{}) {
	na := make([]interface{}, len(args)+2)
	na[0] = logrus.Level(level)
	na[1] = format
	for i, v := range args {
		na[i+2] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Printf print
func Printf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Panicf print panic level message
func Panicf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Fatalf print fatal level message
func Fatalf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Errorf print error level message
func Errorf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Warnf print warn level message
func Warnf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Infof print info level message
func Infof(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Debugf print debug level message
func Debugf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Tracef print trace level message
func Tracef(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// LogFn log
func LogFn(level Level, fn LogFunction) {
	args := []interface{}{logrus.Level(level), logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// PrintFn print
func PrintFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// PanicFn print panic level message
func PanicFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// FatalFn print fatal level message
func FatalFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// ErrorFn print error level message
func ErrorFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// WarnFn print warn level message
func WarnFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// InfoFn print info level message
func InfoFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// DebugFn print debug level message
func DebugFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// TraceFn print trace level message
func TraceFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Logln log
func Logln(level Level, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = logrus.Level(level)
	for i, v := range args {
		na[i+1] = v
	}
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Println print
func Println(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Panicln print panic level message
func Panicln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Fatalln print fatal level message
func Fatalln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Errorln print error level message
func Errorln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Warnln print warn level message
func Warnln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Infoln print info level message
func Infoln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Debugln print debug level message
func Debugln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Traceln print trace level message
func Traceln(args ...interface{}) {
	appenderCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

func errors(e jError) error {
	return jError(fmt.Sprint(pkgName, ": ", e.Error()))
}

func fmtPrintln(args ...interface{}) {
	if cd.Debug {
		fmt.Println(args...)
	}
	if logFunc != nil {
		logFunc(args...)
	}
}

func createAppender() error {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	apMap = make(map[string]*logrus.Logger)
	ap = make(map[string]*appender)
	l := logrus.New()
	l.SetFormatter(&logrus.TextFormatter{TimestampFormat: jtime.DateTime})
	l.SetOutput(os.Stdout)
	l.SetLevel(logrus.DebugLevel)
	apMap[console] = l
	var err error
	for k, v := range cd.Appender {
		var m map[string]interface{}
		if m, err = jcast.StringMapInterface(v); err != nil {
			return err
		}
		nl := logger{}.getDefault()
		if err = jfile.Convert(m, nl); err != nil {
			return err
		}
		for pk, pv := range cd.Params {
			nl.Output.P = strings.ReplaceAll(nl.Output.P, fmt.Sprint("${", pk, "}"), jcast.String(pv))
			nl.Output.LinkName = strings.ReplaceAll(nl.Output.LinkName, fmt.Sprint("${", pk, "}"), jcast.String(pv))
		}
		options := make([]rotatelogs.Option, 0)
		if nl.Output.UTC {
			options = append(options, rotatelogs.WithClock(rotatelogs.UTC))
		}
		if len(nl.Output.LinkName) > 0 {
			options = append(options, rotatelogs.WithLinkName(nl.Output.LinkName))
		}
		var d time.Duration
		if d, err = jtime.ParseTimeDuration(nl.Output.MaxAgeDuration); nl.Output.MaxAge > 0 && err == nil {
			options = append(options, rotatelogs.WithMaxAge(nl.Output.MaxAge*d))
		} else if nl.Output.RotationCount > 0 {
			options = append(options, rotatelogs.WithRotationCount(nl.Output.RotationCount))
		}
		if d, err = jtime.ParseTimeDuration(nl.Output.RotationTimeDuration); nl.Output.RotationTime > 0 && err == nil {
			options = append(options, rotatelogs.WithRotationTime(nl.Output.RotationTime*d))
		}
		var unit jfile.SizeUnit
		if unit, err = jfile.ParseSizeUnit(nl.Output.RotationSizeUnit); nl.Output.RotationSize > 0 && err == nil {
			options = append(options, rotatelogs.WithRotationSize(nl.Output.RotationSize*unit.ToInt64()))
		}
		l = logrus.New()
		switch strings.ToUpper(nl.Formatter.Type) {
		case "TEXT":
			l.SetFormatter(nl.Formatter.Text)
		case "JSON":
			l.SetFormatter(nl.Formatter.Json)
		}
		l.SetLevel(nl.getLogrusLevel())
		var out *rotatelogs.RotateLogs
		if out, err = rotatelogs.New(nl.Output.P, options...); err != nil {
			return err
		}
		l.SetOutput(out)
		apMap[k] = l
	}
	ap[Default] = appender{}.getDefault()
	ap[Default].AddLogger(apMap[console])
	for _, lv := range cd.Logs {
		for _, pv := range lv.Program {
			var pn string
			if pv == Default || strings.Index(pv, pkgKey) == 0 {
				pn = pv
			} else {
				pn = fmt.Sprint(strings.Trim(pv, ".go"), ".go")
			}
			ap[pn] = appender{}.getDefault()
			for _, av := range lv.Appender {
				if l = apMap[av]; l != nil {
					add := true
					for _, cl := range ap[pn].list {
						if cl == l {
							add = false
							break
						}
					}
					if add {
						ap[pn].AddLogger(l)
					}
				}
			}
		}
	}
	return nil
}

func appenderCall(pn, fn, pkg string, args ...interface{}) {
	a := GetAppender(pn, fmt.Sprint(pkgKey, pkg))
	if a == nil {
		fmtPrintln(errors(errorAppenderNil))
	} else {
		if err := a.Call(fn, args...); err != nil {
			fmtPrintln(err)
		}
	}
}
