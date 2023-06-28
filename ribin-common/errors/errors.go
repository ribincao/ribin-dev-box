package errors

import "fmt"

type Error struct {
	error
	Code    int32
	Message string
}

func (e *Error) Error() string {
	if e == nil {
		return ""
	}
	return fmt.Sprintf("code : %d, msg : %s", e.Code, e.Message)
}

func New(code int32, msg string) *Error {
	return &Error{
		Code:    code,
		Message: msg,
	}
}
