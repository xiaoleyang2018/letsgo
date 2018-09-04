// Package logger is a simple wrap for logrus.
package logger

import (
	"fmt"
	"io"
	"os"
	"path"
	"runtime"
	"time"

	colorable "github.com/mattn/go-colorable"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

// Parse setting logs.
func Parse(path, mode, format string) error {
	var out io.Writer = os.Stderr

	if mode == "prod" {
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		out = f
	}

	log = logrus.New()

	if format == "json" {
		log.Formatter = &logrus.JSONFormatter{}
	}
	log.Out = out
	log.AddHook(&ContextHook{callDepth: 6, token: "default"})
	log.SetLevel(logrus.DebugLevel)

	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})
	logrus.SetOutput(colorable.NewColorableStdout()) // for Windows

	return nil
}

// ContextHook is custrom context hook which implements logrus.Hook.
type ContextHook struct {
	token     string
	callDepth int
}

// Levels .
func (hook ContextHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Fire .
func (hook *ContextHook) Fire(entry *logrus.Entry) error {
	if pc, file, line, ok := runtime.Caller(hook.callDepth); ok {
		funcName := runtime.FuncForPC(pc).Name()
		entry.Data["date"] = time.Now().Format("20060102-15:04:05.0000")
		entry.Data["position"] = fmt.Sprintf("%s:%d", path.Base(file), line)
		entry.Data["func"] = path.Base(funcName)
		entry.Data["pid"] = os.Getpid()
		// entry.Data["token"] = hook.token
	}

	return nil
}

// Errorf logs a message at level Error.
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Infof logs a message at level Info.
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Printf logs a message at level Info.
func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

// Warnf logs a message at level Warn.
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Fatalf logs a message at level Fatal.
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Debugf logs a message at level Debug.
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Panicf logs a message at level Panic.
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// Warningf logs a message at level Warn.
func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
}
