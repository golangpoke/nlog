package nlog

import (
	"context"
	"fmt"
	"log/slog"
)

type skipLogger int

func Skip(skip int) skipLogger {
	return skipLogger(skip)
}

func (sk skipLogger) DEBUf(msg string, args ...any) {
	skipLog(LvlDebug, sk, msg, args...)
}

func (sk skipLogger) INFOf(msg string, args ...any) {
	skipLog(LvlInfo, sk, msg, args...)
}

func (sk skipLogger) WARNf(msg string, args ...any) {
	skipLog(LvlWarn, sk, msg, args...)
}

func (sk skipLogger) ERROf(msg string, args ...any) {
	skipLog(LvlError, sk, msg, args...)
}

func (sk skipLogger) PANICf(msg string, args ...any) {
	skipLog(LvlPanic, sk, msg, args...)
}

func skipLog(level Level, skip skipLogger, msg string, args ...any) skipLogger {
	msg = fmt.Sprintf(msg, args...)
	fn, ln := caller(int(skip))
	msg = fmt.Sprintf("%s %s:%d", msg, fn, ln)
	slog.Default().Log(context.Background(), level.Level(), msg)
	return skip
}
