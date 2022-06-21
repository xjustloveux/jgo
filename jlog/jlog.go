// Copyright 2022 JaJa All rights reserved.
// Use of this source code is governed by a MIT-style.
// license that can be found in the LICENSE file.

package jlog

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/xjustloveux/jgo/jcast"
	"github.com/xjustloveux/jgo/jconf"
	"github.com/xjustloveux/jgo/jfile"
	"github.com/xjustloveux/jgo/jruntime"
	"github.com/xjustloveux/jgo/jtime"
	"io"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	errorParamsNumOutIndex = jError("the number of params is out of index")
	errorLoggerNil         = jError("logger is nil")
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
	cd      *configData
	logFunc func(...interface{})
	mux     = new(sync.RWMutex)
	fileMap = make(map[string]*logFile)
	writer  = make(map[string]io.Writer)
	logMap  = make(map[string]*logger)
	handler = make(map[string]Handler)
)

func init() {
	SetFileName(fileName)
	writer[console] = os.Stdout
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

// NewRotateLogs create new rotate logs
func NewRotateLogs(name string, o *Output, param map[string]string) (*RotateLogs, error) {
	var err error
	var loc *time.Location
	if loc, err = time.LoadLocation(o.Clock); err != nil {
		return nil, err
	}
	var h Handler
	if o.Handler == "" {
		h = nil
	} else {
		h = handler[o.Handler]
	}
	p := o.P
	linkName := o.LinkName
	for pk, pv := range param {
		p = strings.ReplaceAll(p, fmt.Sprint("${", pk, "}"), jcast.String(pv))
		linkName = strings.ReplaceAll(linkName, fmt.Sprint("${", pk, "}"), jcast.String(pv))
	}
	p = strings.ReplaceAll(p, fmt.Sprint("${Program}"), name)
	linkName = strings.ReplaceAll(linkName, fmt.Sprint("${Program}"), name)
	var d time.Duration
	if d, err = jtime.ParseTimeDuration(o.RotationTimeDuration); err != nil {
		return nil, err
	}
	rotationTime := o.RotationTime * d
	var unit jfile.SizeUnit
	if unit, err = jfile.ParseSizeUnit(o.RotationSizeUnit); err != nil {
		return nil, err
	}
	rotationSize := o.RotationSize * unit.ToInt64()
	if d, err = jtime.ParseTimeDuration(o.MaxAgeDuration); err != nil {
		return nil, err
	}
	maxAge := o.MaxAge * d
	return &RotateLogs{
		clock:         loc,
		handler:       h,
		mux:           new(sync.RWMutex),
		fileName:      p,
		linkName:      linkName,
		rotationTime:  rotationTime,
		current:       "",
		previous:      "",
		currentLink:   "",
		previousLink:  "",
		rotationSize:  rotationSize,
		maxAge:        maxAge,
		rotationCount: o.RotationCount,
	}, nil
}

// AddWriter add io.Writer with name
func AddWriter(name string, w io.Writer) {
	if w != nil {
		writer[name] = w
	}
}

// AddHandler add Handler with name
func AddHandler(name string, h Handler) {
	if h != nil {
		handler[name] = h
	}
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
	if err := createLogger(); err != nil {
		return err
	}
	return nil
}

// GetParam returns conf.Params with key
func GetParam(key string) string {
	return cd.Params[key]
}

// GetLogger returns *logger with program file or package name.
// if not found logger then default return console logger.
func GetLogger(args ...string) *logger {
	if len(args) > 0 {
		for _, k := range args {
			if l := logMap[k]; l != nil {
				return l
			}
		}
	} else {
		if l := logMap[jruntime.GetCallerProgramName()]; l != nil {
			return l
		}
		if l := logMap[fmt.Sprint(pkgKey, jruntime.GetCallerPkgName())]; l != nil {
			return l
		}
	}
	return logMap[Default]
}

// Log log
func Log(level Level, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = logrus.Level(level)
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Print print
func Print(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Panic print panic level message
func Panic(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Fatal print fatal level message
func Fatal(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Error print error level message
func Error(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Warn print warn level message
func Warn(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Info print info level message
func Info(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Debug print debug level message
func Debug(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Trace print trace level message
func Trace(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Logf log
func Logf(level Level, format string, args ...interface{}) {
	na := make([]interface{}, len(args)+2)
	na[0] = logrus.Level(level)
	na[1] = format
	for i, v := range args {
		na[i+2] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Printf print
func Printf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Panicf print panic level message
func Panicf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Fatalf print fatal level message
func Fatalf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Errorf print error level message
func Errorf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Warnf print warn level message
func Warnf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Infof print info level message
func Infof(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Debugf print debug level message
func Debugf(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Tracef print trace level message
func Tracef(format string, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = format
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// LogFn log
func LogFn(level Level, fn LogFunction) {
	args := []interface{}{logrus.Level(level), logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// PrintFn print
func PrintFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// PanicFn print panic level message
func PanicFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// FatalFn print fatal level message
func FatalFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// ErrorFn print error level message
func ErrorFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// WarnFn print warn level message
func WarnFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// InfoFn print info level message
func InfoFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// DebugFn print debug level message
func DebugFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// TraceFn print trace level message
func TraceFn(fn LogFunction) {
	args := []interface{}{logrus.LogFunction(fn)}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Logln log
func Logln(level Level, args ...interface{}) {
	na := make([]interface{}, len(args)+1)
	na[0] = logrus.Level(level)
	for i, v := range args {
		na[i+1] = v
	}
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), na...)
}

// Println print
func Println(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Panicln print panic level message
func Panicln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Fatalln print fatal level message
func Fatalln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Errorln print error level message
func Errorln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Warnln print warn level message
func Warnln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Infoln print info level message
func Infoln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Debugln print debug level message
func Debugln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
}

// Traceln print trace level message
func Traceln(args ...interface{}) {
	loggerCall(jruntime.GetCallerProgramName(), jruntime.GetFuncName(), jruntime.GetCallerPkgName(), args...)
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

func createLogger() error {
	cd.appender = make(map[string]*appender)
	for k, v := range cd.Appender {
		cd.appender[k] = appender{}.getDefault()
		if err := jfile.Convert(v, cd.appender[k]); err != nil {
			return err
		}
	}
	if cd.appender[console] == nil {
		cd.appender[console] = appender{}.getDefault()
		cd.appender[console].Level = "Debug"
		cd.appender[console].Output.Name = console
	}
	cd.Logs = append(cd.Logs, &logs{
		Program:  []string{Default},
		Appender: []string{console},
	})
	for _, cl := range cd.Logs {
		for _, pv := range cl.Program {
			var pn string
			if pv == Default || strings.Index(pv, pkgKey) == 0 {
				pn = pv
			} else {
				pn = fmt.Sprint(strings.Trim(pv, ".go"), ".go")
			}
			log := logMap[pn]
			if log == nil {
				log = logger{}.getDefault()
				logMap[pn] = log
			}
			for _, av := range cl.Appender {
				if a := cd.appender[av]; a != nil {
					if err := log.addLogger(av, a, cd.Params); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func loggerCall(pn, fn, pkg string, args ...interface{}) {
	l := GetLogger(pn, fmt.Sprint(pkgKey, pkg))
	if l == nil {
		fmtPrintln(errors(errorLoggerNil))
	} else {
		if err := l.Call(fn, args...); err != nil {
			fmtPrintln(err)
		}
	}
}

func getFile(name string) (*logFile, error) {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if fileMap[name] == nil {
		l := &logFile{name: name, mux: new(sync.RWMutex), file: nil}
		if err := l.open(true); err != nil {
			return nil, err
		}
		fileMap[name] = l
		return l, nil
	}
	return fileMap[name], nil
}

func removeFile(name string) error {
	mux.Lock()
	defer func() {
		mux.Unlock()
	}()
	if fileMap[name] != nil {
		if err := fileMap[name].close(true); err != nil {
			return err
		}
	}
	return nil
}
