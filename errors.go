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

	// Kind represent the type of error
	Kind Kind

	// Message is the information about the error
	Message string

	// metadata is the way to wrap context to the error
	metadata map[string]interface{}

	// the wrapped error
	Err error
}

func Wrap(err error, format string, i ...interface{}) *Error {
	e, ok := err.(*Error)

	newErr := &Error{
		op:       newOperation(),
		Message:  fmt.Sprintf(format, i...),
		metadata: map[string]interface{}{},
		Err:      err,
	}
	if ok {
		newErr.Kind = e.Kind
	}

	return newErr
}

func New(format string, i ...interface{}) *Error {
	return &Error{
		op:       newOperation(),
		Message:  fmt.Sprintf(format, i...),
		metadata: map[string]interface{}{},
		Err:      nil,
	}
}

func NewWithKind(kind Kind, format string, i ...interface{}) *Error {
	return &Error{
		op:       newOperation(),
		Kind:     kind,
		Message:  fmt.Sprintf(format, i...),
		metadata: map[string]interface{}{},
		Err:      nil,
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

func (e *Error) FormatStacktrace() string {
	st := ""

	var tmp error = e
	for tmp != nil {
		x, ok := tmp.(*Error)

		if ok {
			st += fmt.Sprintf("%s:%d %s.%s => %s", x.op.file, x.op.line, x.op.pkg, x.op.function, x.Message)

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

func (e *Error) Stacktrace() []string {
	var st []string

	var tmp error = e
	for tmp != nil {
		x, ok := tmp.(*Error)

		if ok {
			st = append(st, fmt.Sprintf("%s/%s.%s:%d", x.op.file, x.op.pkg, x.op.function, x.op.line))
			tmp = x.Err
		} else {
			break
		}
	}

	return st
}

func (e *Error) WithMessage(format string, i ...interface{}) *Error {
	e.Message = fmt.Sprintf(format, i...)
	return e
}

func (e *Error) WithMetadata(key string, value interface{}) *Error {
	e.metadata[key] = value
	return e
}

func (e *Error) WithKind(kind Kind) *Error {
	e.Kind = kind
	return e
}

func (e *Error) GetMetadata(key string) (value interface{}, ok bool) {
	value, ok = e.metadata[key]
	return value, ok
}

type operation struct {
	pkg      string
	file     string
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
		file:     fileName,
		function: funcName,
		line:     line,
	}
}
