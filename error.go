package nlog

import "fmt"

func Catch(err error) error {
	fn, ln := caller(1)
	return fmt.Errorf("%w %s:%d", err, fn, ln)
}
