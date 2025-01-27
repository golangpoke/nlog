package nlog_test

import (
	"github.com/golangpoke/nlog"
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	defer nlog.Recovery()
	nlog.SetLogger(&nlog.Logger{
		Level: nlog.LvlWarn,
	})
	nlog.DEBUf("this is %s %d log", "debug", 1)
	nlog.INFOf("this is %s %d log", "info", 2)
	nlog.WARNf("this is %s %d log", "warn", 3)
	nlog.ERROf("this is %s %d log", "err", 4)
	nlog.PANICf("this is %s %d log", "panic", 5)
	log.Println(123)
}
