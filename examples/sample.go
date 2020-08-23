package main

import (
	"github.com/expectedsh/errors"
)

func main() {
	t := S{}

	t.S().Log().Error()
}

type S struct {
}

func (s S) S() *errors.Error {
	return errors.Wrap(X(), "s has not pass the test")
}

func X() *errors.Error {
	t := T()

	return errors.Wrap(t, "xdlol").WithField("lol", 123)
}

func T() *errors.Error {
	return errors.New("test %d", 5)
}
