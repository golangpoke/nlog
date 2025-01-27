package nlog

import (
	"fmt"
	"github.com/golangpoke/nlog/console"
	"path/filepath"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

type Level int

type Logger struct {
	Level      Level
	TimeFormat string
	out        string
}

var defaultLogger atomic.Pointer[Logger]

func init() {
	defaultLogger.Store(&Logger{
		Level:      LvlDebug,
		TimeFormat: "2006-01-02 15:04:05",
	})
}

func newLogger() *Logger {
	return defaultLogger.Load()
}

func SetLogger(logger *Logger) *Logger {
	if logger.Level != 0 {
		newLogger().Level = logger.Level
	}
	if logger.TimeFormat != "" {
		newLogger().TimeFormat = logger.TimeFormat
	}
	return logger
}

func (l *Logger) output(level Level, msg string, args ...any) *Logger {
	if level < l.Level {
		return l
	}
	tm := time.Now().Format(l.TimeFormat)
	tmc := console.TxWhite(tm)
	lv := mapLevelMessage(level)
	ms := fmt.Sprintf(msg, args...)
	fn, ln := caller(3)
	l.out = strings.Join([]string{tmc, lv, ms, fmt.Sprintf("%s:%d", fn, ln)}, " ")
	return l
}

func (l *Logger) print() {
	if l.out == "" {
		return
	}
	fmt.Println(l.out)
}

func DEBUf(msg string, args ...any) {
	newLogger().output(LvlDebug, msg, args...).print()
}

func INFOf(msg string, args ...any) {
	newLogger().output(LvlInfo, msg, args...).print()
}

func WARNf(msg string, args ...any) {
	newLogger().output(LvlWarn, msg, args...).print()
}

func ERROf(msg string, args ...any) {
	newLogger().output(LvlError, msg, args...).print()
}

func PANICf(msg string, args ...any) {
	panic(newLogger().output(LvlPanic, msg, args...).out)
}

func Recovery() {
	if rec := recover(); rec != nil {
		fmt.Println(rec)
	}
}

func caller(skip int) (fileName string, line int) {
	_, file, line, _ := runtime.Caller(skip)
	fileName = filepath.Base(file)
	return fileName, line
}
