package nlog

import (
	"fmt"
	"github.com/golangpoke/nlog/console"
)

func toJsonString(args ...string) (s string) {
	for i := range args {
		if i%2 == 0 {
			args[i] = console.DpBold(args[i]) + ":"
			s += fmt.Sprintf("  %s ", args[i])
		} else {
			s += fmt.Sprintf("%s\n", args[i])
		}
	}
	return fmt.Sprintf("%s\n%s%s", console.TxWhite("{"), s, console.TxWhite("}"))
}
