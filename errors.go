package xerrors

import (
	"fmt"
	"strconv"
)

type Error interface {
	Code() int64
	IsLogic() bool
	WithError(err error) Error
	WithDetails(format string, args ...interface{}) Error
	Error() string
}

type xError struct {
	logic bool
	code  int64
	msg   string
}

func New(code int64, text string, logic bool) Error {
	return &xError{
		code:  code,
		msg:   text,
		logic: logic,
	}
}

func (e xError) Code() int64 {
	return e.code
}

func (e xError) IsLogic() bool {
	return e.logic
}

func (e xError) WithError(err error) Error {
	e.msg += ": " + err.Error()
	return e
}

func (e xError) WithDetails(format string, args ...interface{}) Error {
	e.msg += ": " + fmt.Sprintf(format, args...)
	return e
}

func (e xError) Error() string {
	return "code: " + strconv.FormatInt(e.code, 10) + ", error: " + e.msg
}
