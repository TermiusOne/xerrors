package xerrors

import (
	"fmt"
	"strconv"
)

type XError struct {
	logic bool
	code  int64
	msg   string
}

func New(code int64, text string, logic bool) *XError {
	return &XError{
		code:  code,
		msg:   text,
		logic: logic,
	}
}

func (e XError) Code() int64 {
	return e.code
}

func (e XError) IsLogic() bool {
	return e.logic
}

func (e *XError) WithError(err error) *XError {
	e.msg += ": " + err.Error()
	return e
}

func (e *XError) WithDetails(format string, args ...interface{}) *XError {
	e.msg += ": " + fmt.Sprintf(format, args...)
	return e
}

func (e *XError) Error() string {
	return "code: " + strconv.FormatInt(e.code, 10) + ", error: " + e.msg
}
