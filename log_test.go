package nlog_test

import (
	"log"
	"nlog"
	"testing"
)

func TestLogger(t *testing.T) {
	defer nlog.Recovery()
	nlog.SetLogger(&nlog.Logger{
		Level:      0,
		TimeFormat: "",
	})
	nlog.DEBUf("%s log", "debug")
	nlog.INFOf("%s log", "info")
	nlog.WARNf("%s log", "warn")
	nlog.ERROf("%s log", "error")
	nlog.PANICf("%s log", "panic")
	log.Println(123)
}
