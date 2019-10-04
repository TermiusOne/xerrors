package xerrors

import (
	"fmt"
)

type XError interface {
	Code() int
	Error() string
	Wrap(a ...interface{}) XError
	Wrapf(format string, a ...interface{}) XError
}

type baseError struct {
	code int
	msg  string
}

func New(code int, text string) XError {
	return &baseError{
		code: code,
		msg:  text,
	}
}

func (e *baseError) Error() string {
	return e.msg
}

func (e *baseError) Wrap(a ...interface{}) XError {
	e.msg = fmt.Sprintf("%s: %s", e.msg, fmt.Sprint(a...))
	return e
}

func (e *baseError) Wrapf(format string, a ...interface{}) XError {
	e.msg = fmt.Sprintf("%s: %s", e.msg, fmt.Sprintf(format, a...))
	return e
}

func (e *baseError) Code() int {
	return e.code
}
