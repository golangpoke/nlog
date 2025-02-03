package nlog_test

import (
	"errors"
	"github.com/golangpoke/nlog"
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	nlog.SetDefault()
	defer nlog.Recovery()
	nlog.DEBUf("%s log", "debug")
	nlog.INFOf("%s log", "info")
	CallTest()
	nlog.WARNf("%s log", "warn")
	nlog.ERROf("%s log", "err")
	nlog.PANICf("%s log", "panic")
	log.Println(123)
}

func TestError(t *testing.T) {
	nlog.SetDefault()
	defer nlog.Recovery()
	err := ErrorTest()
	if err != nil {
		nlog.ERROf(err.Error())
	}
	nlog.NoSource().ERROf("err:%v", err)
	if errors.Is(err, os.ErrNotExist) {
		nlog.INFOf("same error")
	}
}

func ErrorTest() error {
	err := os.ErrNotExist
	if err != nil {
		return nlog.Catch(err)
	}
	return nil
}

func CallTest() {
	msg := "log"
	nlog.Skip(1).INFOf("%s %s", "option", msg)
	nlog.NoSource().WARNf("%s %s", "option", msg)
}
