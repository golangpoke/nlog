package nlog

import "nlog/console"

const (
	LvlDebug Level = iota
	LvlInfo
	LvlWarn
	LvlError
	LvlPanic
)

func mapLevelMessage(level Level) (string, string) {
	switch level {
	case LvlDebug:
		return "debug", console.TxMagenta("[DEBU]")
	case LvlInfo:
		return "info", console.TxGreen("[INFO]")
	case LvlWarn:
		return "warning", console.TxYellow("[WARN]")
	case LvlError:
		return "error", console.TxRed("[ERRO]")
	case LvlPanic:
		return "panic", console.HlRed("[PANIC]")
	}
	return "", ""
}
