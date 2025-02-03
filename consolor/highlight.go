package consolor

import "fmt"

type hlColor int

const (
	hlBlack hlColor = iota + 90
	hlRed
	hlGreen
	hlYellow
	hlBlue
	hlMagenta
	hlCyan
	hlWhite
)

func setHlColor(text string, color hlColor) string {
	return fmt.Sprintf("\033[7;%dm%s\033[0m", color, text)
}

func HlBlack(text string) string {
	return setHlColor(text, hlBlack)
}

func HlRed(text string) string {
	return setHlColor(text, hlRed)
}

func HlGreen(text string) string {
	return setHlColor(text, hlGreen)
}

func HlYellow(text string) string {
	return setHlColor(text, hlYellow)
}

func HlBlue(text string) string {
	return setHlColor(text, hlBlue)
}

func HlMagenta(text string) string {
	return setHlColor(text, hlMagenta)
}

func HlCyan(text string) string {
	return setHlColor(text, hlCyan)
}

func HlWhite(text string) string {
	return setHlColor(text, hlWhite)
}
