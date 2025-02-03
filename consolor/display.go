package consolor

import "fmt"

type display int

const (
	dpBold      display = 1
	dpUnderLine display = 4
)

func setDisplay(display display, text string) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", display, text)
}

func DpBold(text string) string {
	return setDisplay(dpBold, text)
}

func DpUnderLine(text string) string {
	return setDisplay(dpUnderLine, text)
}
