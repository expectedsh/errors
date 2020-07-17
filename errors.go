package errors

import (
	"fmt"
	"path"
	"runtime"
	"strings"
)

type Error struct {
	// op is filled by default
	op *operation

	// StatusCode is only for final httpCall
	StatusCode int

	// Message is the information about the error
	Message string

	// Metadata is the way to wrap context to the error
	Metadata map[string]interface{}

	// the wrapped error
	Err error
}

func Wrap(err error, format string, i ...interface{}) *Error {
	return &Error{
		op:         newOperation(),
		StatusCode: -1,
		Message:    fmt.Sprintf(format, i...),
		Err:        err,
	}
}

func New(format string, i ...interface{}) *Error {
	return &Error{
		op:         newOperation(),
		StatusCode: 0,
		Message:    fmt.Sprintf(format, i...),
		Metadata:   nil,
		Err:        nil,
	}
}

func (e *Error) Error() string {
	var errors []string

	if e.Message != "" {
		errors = append(errors, e.Message)
	}

	tmp := e.Err

	for tmp != nil {
		x, ok := tmp.(*Error)

		fmt.Println(ok)
		if ok {
			if x.Message != "" {
				errors = append(errors, x.Message)
			} else {
				errors = append(errors, x.Error())
			}
			tmp = x.Err
		} else {
			errors = append(errors, tmp.Error())
			break
		}
	}

	finalStr := ""
	for i := len(errors) - 1; i >= 0; i-- {
		finalStr += errors[i]
		if i != 0 {
			finalStr += ": "
		}
	}

	return finalStr
}

func (e *Error) Stacktrace() string {
	st := ""

	var tmp error = e
	for tmp != nil {
		x, ok := tmp.(*Error)

		if ok {
			st += fmt.Sprintf("%s.%s:%d => %s", x.op.pkg, x.op.function, x.op.line, x.Message)

			if x.Err != nil {
				st += "\n  "
			}

			tmp = x.Err
		} else {
			st += tmp.Error()
			break
		}
	}

	return st
}

func (e *Error) WithStatusCode(statusCode int) *Error {
	e.StatusCode = statusCode
	return e
}

func (e *Error) WithMessage(format string, i ...interface{}) *Error {
	e.Message = fmt.Sprintf(format, i...)
	return e
}

type operation struct {
	pkg      string
	fileName string
	function string
	line     int
}

func newOperation() *operation {
	pc, file, line, _ := runtime.Caller(2)
	_, fileName := path.Split(file)
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	packageName := ""
	funcName := parts[pl-1]

	if parts[pl-2][0] == '(' {
		funcName = parts[pl-2] + "." + funcName
		packageName = strings.Join(parts[0:pl-2], ".")
	} else {
		packageName = strings.Join(parts[0:pl-1], ".")
	}

	return &operation{
		pkg:      packageName,
		fileName: fileName,
		function: funcName,
		line:     line,
	}
}
