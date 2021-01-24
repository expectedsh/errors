package main

import (
	"github.com/expectedsh/errors"
)

func main() {
	t := Structure{}

	println(t.SuperCaller().FormatStacktrace())
}

type Structure struct {
}

func (s Structure) SuperCaller() *errors.Error {
	return errors.Wrap(FirstChild(), "error while declaring fields")
}

func FirstChild() *errors.Error {
	t := SecondChildTheOler()

	return errors.Wrap(t, "error always returned :)").
		WithField("lol", 123).
		WithField("xxxxxxxxxxxlol", "2345")
}

func SecondChildTheOler() *errors.Error {
	return errors.New("I am just failing: %d", 5)
}
