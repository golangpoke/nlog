package consolor

import "fmt"

type txColor int

const (
	txBlack txColor = iota + 30
	txRed
	txGreen
	txYellow
	txBlue
	txMagenta
	txCyan
	txWhite
)

func setTxColor(text string, txColor txColor) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", txColor, text)
}

func TxBlack(text string) string {
	return setTxColor(text, txBlack)
}

func TxRed(text string) string {
	return setTxColor(text, txRed)
}

func TxGreen(text string) string {
	return setTxColor(text, txGreen)
}

func TxYellow(text string) string {
	return setTxColor(text, txYellow)
}

func TxBlue(text string) string {
	return setTxColor(text, txBlue)
}

func TxMagenta(text string) string {
	return setTxColor(text, txMagenta)
}

func TxCyan(text string) string {
	return setTxColor(text, txCyan)
}

func TxWhite(text string) string {
	return setTxColor(text, txWhite)
}
