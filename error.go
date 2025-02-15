package nlog

import (
	"errors"
	"fmt"
	"runtime"
)

type Error interface {
	error
	StackTrace() string
	Cause() error
}

type stackError struct {
	err    error
	stacks []string
}

func (e *stackError) Error() string {
	s := e.err.Error()
	s += e.StackTrace()
	return s
}

func (e *stackError) StackTrace() (s string) {
	l := len(e.stacks) - 1
	for i := l; i >= 0; i-- {
		if i%2 != 0 {
			s += "\n[" + e.stacks[i] + "]"
		} else {
			s += "\n\t" + e.stacks[i]
		}
	}
	return s
}

func (e *stackError) Cause() error {
	return e.err
}

func UnWrap(err error) error {
	var e *stackError
	ok := errors.As(err, &e)
	if ok {
		return e.err
	}
	return err
}

func Wrap(err error) (e *stackError) {
	funcName, curStack := loadStack(0)
	stack := []string{curStack, funcName}
	if !errors.As(err, &e) {
		return &stackError{
			err:    err,
			stacks: stack,
		}
	}
	e.stacks = append(e.stacks, stack...)
	return e
}

func loadStack(skip int) (string, string) {
	pc, fl, ln, _ := runtime.Caller(skip + 2)
	fn := runtime.FuncForPC(pc).Name()
	stack := fmt.Sprintf("%s:%d", fl, ln)
	return fn, stack
}
