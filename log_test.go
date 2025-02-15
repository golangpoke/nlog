package nlog_test

import (
	"github.com/golangpoke/nlog"
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	defer nlog.Recovery()
	nlog.SetDefault(nlog.LvlDebug)
	nlog.DEBUf("%s log", "debug")
	nlog.INFOf("%s log", "info")
	nlog.WARNf("%s log", "warn")
	nlog.ERROf("%v", err3())
	nlog.PANICf("%s log", "panic")
	log.Println(123)
}

func TestOptionLogger(t *testing.T) {
	defer nlog.Recovery()
	nlog.SetDefault(nlog.LvlDebug)
	nlog.NoSource().INFOf("%s log", "info")
	SkipTest()
	nlog.ERROf("%v", err3())
}

func SkipTest() {
	nlog.Skip(1).INFOf("%s log", "info")
}
