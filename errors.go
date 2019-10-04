package xerrors

import (
	"fmt"
)

type XError struct {
	code int
	msg  string
}

func New(code int, text string) *XError {
	return &XError{
		code: code,
		msg:  text,
	}
}

func (e *XError) Error() string {
	return e.msg
}

func (e *XError) Wrap(a ...interface{}) *XError {
	e.msg = fmt.Sprintf("%s: %s", e.msg, fmt.Sprint(a...))
	return e
}

func (e *XError) Wrapf(format string, a ...interface{}) *XError {
	e.msg = fmt.Sprintf("%s: %s", e.msg, fmt.Sprintf(format, a...))
	return e
}

func (e *XError) Code() int {
	return e.code
}
