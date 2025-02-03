package nlog

import (
	"context"
	"fmt"
	"github.com/golangpoke/nlog/consolor"
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
)

type Level int

func (l Level) Level() slog.Level {
	return slog.Level(l)
}

const (
	LvlDebug Level = iota + 1
	LvlInfo
	LvlWarn
	LvlError
	LvlPanic
)

type logHandler struct {
	level      slog.Level
	handler    slog.Handler
	logger     *log.Logger
	timeFormat string
}

func SetDefault() {
	slog.SetDefault(slog.New(&logHandler{
		level:      LvlInfo.Level(),
		handler:    slog.NewTextHandler(os.Stdout, nil),
		logger:     log.New(os.Stdout, "", 0),
		timeFormat: "2006-01-02 15:04:05",
	}))
}

func (lh *logHandler) Handle(ctx context.Context, record slog.Record) error {
	tm := record.Time.Format(lh.timeFormat)
	tm = consolor.TxWhite(tm)
	lv := ""
	switch record.Level {
	case LvlDebug.Level():
		lv = consolor.TxMagenta("[DEBU]")
	case LvlInfo.Level():
		lv = consolor.TxGreen("[INFO]")
	case LvlWarn.Level():
		lv = consolor.TxYellow("[WARN]")
	case LvlError.Level():
		lv = consolor.TxRed("[ERRO]")
	case LvlPanic.Level():
		lv = consolor.HlRed("[PANIC]")
		panic(fmt.Sprintf("%s %s %s", tm, lv, record.Message))
		return nil
	}
	lh.logger.Println(tm, lv, record.Message)
	return nil
}

func (lh *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return lh.handler.WithAttrs(attrs)
}

func (lh *logHandler) WithGroup(name string) slog.Handler {
	return lh.handler.WithGroup(name)
}

func (lh *logHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return lh.level <= level
}

func DEBUf(msg string, args ...any) {
	defaultLog(LvlDebug, msg, args...)
}

func INFOf(msg string, args ...any) {
	defaultLog(LvlInfo, msg, args...)
}

func WARNf(msg string, args ...any) {
	defaultLog(LvlWarn, msg, args...)
}

func ERROf(msg string, args ...any) {
	defaultLog(LvlError, msg, args...)
}

func PANICf(msg string, args ...any) {
	defaultLog(LvlPanic, msg, args...)
}

func Recovery() {
	if rec := recover(); rec != nil {
		fmt.Println(rec)
	}
}

func defaultLog(level Level, msg string, args ...any) {
	msg = fmt.Sprintf(msg, args...)
	fn, ln := caller(2)
	msg = fmt.Sprintf("%s %s:%d", msg, fn, ln)
	slog.Default().Log(context.Background(), level.Level(), msg)
}

func caller(skip int) (fileName string, line int) {
	_, file, line, _ := runtime.Caller(skip + 1)
	fileName = filepath.Base(file)
	return fileName, line
}
