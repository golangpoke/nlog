package nlog

import (
	"fmt"
	"nlog/console"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type Level int

type Logger struct {
	Level      Level
	TimeFormat string
	AddSource  bool
}

func defaultLogger() *Logger {
	return &Logger{
		Level:      LvlDebug,
		TimeFormat: "2006-01-02 15:04:05",
		AddSource:  true,
	}
}

func SetLogger(logger *Logger) *Logger {
	if logger.Level != defaultLogger().Level {
		defaultLogger().Level = logger.Level
	}
	if logger.TimeFormat != defaultLogger().TimeFormat {
		defaultLogger().TimeFormat = logger.TimeFormat
	}
	return logger
}

func (l *Logger) output(level Level, msg string, args ...any) string {
	tm := time.Now().Format(l.TimeFormat)
	tmc := console.TxWhite(tm)
	lv, lvm := mapLevelMessage(level)
	ms := fmt.Sprintf(msg, args...)
	msc := console.TxCyan(ms)
	fn, ln, fc := caller(3)
	mp := toJsonString(
		"time", tm,
		"level", lv,
		"message", ms,
		"source", fmt.Sprintf("%s:%d %s", fn, ln, fc),
	)
	return strings.Join([]string{tmc, lvm, msc, mp}, " ")
}

func DEBUf(msg string, args ...any) {
	fmt.Println(defaultLogger().output(LvlDebug, msg, args...))
}

func INFOf(msg string, args ...any) {
	fmt.Println(defaultLogger().output(LvlInfo, msg, args...))
}

func WARNf(msg string, args ...any) {
	fmt.Println(defaultLogger().output(LvlWarn, msg, args...))
}

func ERROf(msg string, args ...any) {
	fmt.Println(defaultLogger().output(LvlError, msg, args...))
}

func PANICf(msg string, args ...any) {
	panic(defaultLogger().output(LvlPanic, msg, args...))
}

func Recovery() {
	if rec := recover(); rec != nil {
		fmt.Println(rec)
	}
}

func caller(skip int) (fileName string, line int, funcName string) {
	pc, file, line, _ := runtime.Caller(skip)
	funcName = runtime.FuncForPC(pc).Name()
	fileName = filepath.Base(file)
	return fileName, line, funcName
}
