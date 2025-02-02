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
	msg := "call log"
	nlog.Skip(1).DEBUf(msg)
	nlog.Skip(1).INFOf(msg)
	nlog.Skip(1).WARNf(msg)
	nlog.Skip(1).ERROf(msg)
}
