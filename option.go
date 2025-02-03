package nlog

import (
	"context"
	"fmt"
	"log/slog"
)

type optionLogger struct {
	skip      int
	addSource bool
}

type option func(ol *optionLogger)

func Option(option ...option) *optionLogger {
	ol := &optionLogger{
		skip:      0,
		addSource: true,
	}
	for _, o := range option {
		o(ol)
	}
	return ol
}

func Skip(skip int) option {
	return func(ol *optionLogger) {
		ol.skip = skip
	}
}

func NoSource() option {
	return func(ol *optionLogger) {
		ol.addSource = false
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
	if ol.addSource {
		fn, ln := caller(ol.skip + 2)
		msg = fmt.Sprintf("%s %s:%d", msg, fn, ln)
	}
	slog.Default().Log(context.Background(), level.Level(), msg)
}
