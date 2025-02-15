package nlog

import (
	"context"
	"fmt"
	"log/slog"
)

type optionLogger struct {
	skip     int
	noSource bool
}

func Skip(skip int) *optionLogger {
	return &optionLogger{
		skip: skip,
	}
}

func NoSource() *optionLogger {
	return &optionLogger{
		noSource: true,
	}
}

func (ol *optionLogger) DEBUf(msg string, args ...any) {
	optionLog(ol, LvlDebug, msg, args...)
}

func (ol *optionLogger) INFOf(msg string, args ...interface{}) {
	optionLog(ol, LvlInfo, msg, args...)
}

func (ol *optionLogger) WARNf(msg string, args ...interface{}) {
	optionLog(ol, LvlWarn, msg, args...)
}

func (ol *optionLogger) ERROf(msg string, args ...interface{}) {
	optionLog(ol, LvlError, msg, args...)
}

func (ol *optionLogger) PANICf(msg string, args ...interface{}) {
	optionLog(ol, LvlPanic, msg, args...)
}

func optionLog(ol *optionLogger, level Level, msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	if !ol.noSource {
		fn, ln := caller(ol.skip + 2)
		msg = fmt.Sprintf("%s %s:%d", msg, fn, ln)
	}
	slog.Default().Log(context.Background(), level.Level(), msg)
}
