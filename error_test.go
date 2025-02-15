package nlog_test

import (
	"errors"
	"fmt"
	"github.com/golangpoke/nlog"
	"testing"
)

func TestWrap(t *testing.T) {
	fmt.Println(err3())
	if errors.Is(nlog.UnWrap(err3()), err1) {
		fmt.Println("same error")
	}
}

var err1 = errors.New("err1")

func err2() error {
	return nlog.Wrap(err1)
}

func err3() error {
	return nlog.Wrap(err2())
}

func BenchmarkWrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = nlog.Wrap(err1)
	}
}
