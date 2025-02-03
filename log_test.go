package nlog_test

import (
	"github.com/golangpoke/nlog"
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	nlog.Use()
	defer nlog.Recovery()
	nlog.DEBUf("%s log", "debug")
	nlog.INFOf("%s log", "info")
	CallTest()
	nlog.WARNf("%s log", "warn")
	nlog.ERROf("%s log", "err")
	nlog.PANICf("%s log", "panic")
	log.Println(123)
}

func CallTest() {
	msg := "log"
	nlog.Option(
		nlog.Skip(1),
	).INFOf("%s %s", "option", msg)
	nlog.Option(
		nlog.NoSource(),
	).WARNf("%s %s", "option", msg)
}
