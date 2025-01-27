package nlog

import "github.com/golangpoke/nlog/console"

const (
	LvlDebug Level = iota + 1
	LvlInfo
	LvlWarn
	LvlError
	LvlPanic
)

func mapLevelMessage(level Level) string {
	switch level {
	case LvlDebug:
		return console.TxMagenta("[DEBU]")
	case LvlInfo:
		return console.TxGreen("[INFO]")
	case LvlWarn:
		return console.TxYellow("[WARN]")
	case LvlError:
		return console.TxRed("[ERRO]")
	case LvlPanic:
		return console.HlRed("[PANIC]")
	}
	return ""
}
