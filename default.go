package nlog

import (
	"context"
	"fmt"
	"github.com/golangpoke/nlog/consolor"
	"log/slog"
	"runtime"
)

type Level int

func (l Level) Level() slog.Level {
	return slog.Level(l)
}

const (
	LvlDebug Level = -4
	LvlInfo  Level = 0
	LvlWarn  Level = 4
	LvlError Level = 8
	LvlPanic Level = 12
)

type logHandler struct {
	level      slog.Level
	timeFormat string
}

func (lh *logHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= lh.level
}

func (lh *logHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return lh
}

func (lh *logHandler) WithGroup(name string) slog.Handler {
	return lh
}

func SetDefault(level Level) {
	slog.SetDefault(slog.New(Default(level)))
}

func Default(level Level) *logHandler {
	return &logHandler{
		level:      level.Level(),
		timeFormat: "2006-01-02 15:04:05",
	}
}

func (lh *logHandler) Handle(ctx context.Context, record slog.Record) error {
	tm := record.Time.Format(lh.timeFormat)
	tm = consolor.TxWhite(tm)
	lv := ""
	curStack := ""
	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key == FileNameKey {
			curStack += fmt.Sprintf("%s", attr.Value)
		}
		if attr.Key == LineNumberKey {
			curStack += fmt.Sprintf(":%s", attr.Value)
		}
		return true
	})
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
		lv = consolor.HlRed("[PANI]")
	}
	s := fmt.Sprintf("%s %s %s %s\n\n", tm, curStack, lv, record.Message)
	if record.Level == LvlPanic.Level() {
		panic(s)
	}
	fmt.Print(s)
	return nil
}

func DEBUf(msg string, args ...any) {
	defaultLogWithStack(LvlDebug, msg, args...)
}

func INFOf(msg string, args ...any) {
	defaultLogWithStack(LvlInfo, msg, args...)
}

func WARNf(msg string, args ...any) {
	defaultLogWithStack(LvlWarn, msg, args...)
}

func ERROf(msg string, args ...any) {
	defaultLogWithStack(LvlError, msg, args...)
}

func PANICf(msg string, args ...any) {
	defaultLogWithStack(LvlPanic, msg, args...)
}

func Recovery() {
	if rec := recover(); rec != nil {
		fmt.Println(rec)
	}
}

const (
	FileNameKey   = "fileName"
	LineNumberKey = "lineNumber"
	ErrorKey      = "error"
)

func defaultLogWithStack(level Level, msg string, args ...any) {
	fn, ln := caller(2)
	msg = fmt.Sprintf(msg, args...)
	slog.Default().LogAttrs(context.Background(), level.Level(), msg,
		slog.String(FileNameKey, fn),
		slog.Int(LineNumberKey, ln),
	)
}

func caller(skip int) (string, int) {
	_, file, line, _ := runtime.Caller(skip + 1)
	return file, line
}
