package main

import (
	"fmt"
	"strings"

	"github.com/expectedsh/errors"
)

func main() {
	t := S{}

	fmt.Println(t.S().FormatStacktrace())
	fmt.Println(strings.Join(t.S().Stacktrace(), ","))
	fmt.Println(t.S().Error())
}

type S struct {
}

func (s S) S() *errors.Error {
	return errors.Wrap(X(), "s has not pass the test")
}

func X() *errors.Error {
	t := T()

	return errors.Wrap(t, "xdlol")
}

func T() *errors.Error {
	return errors.New("test %d", 5)
}
