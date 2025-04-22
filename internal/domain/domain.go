package domain

import (
	"fmt"
)

type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

func (e *Error) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s: %v", e.Op, e.Message, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Op, e.Message)
}

func (e *Error) Unwrap() error {
	return e.Err
}

func NewError(code, message, op string, err error) *Error {
	return &Error{
		Code:    code,
		Message: message,
		Op:      op,
		Err:     err,
	}
}

func ArgumentsInvalid(op string, err error) *Error {
	return NewError("arguments_invalid", "Arguments is invalid", op, err)
}

func ErrNotFound(op string, err error) *Error {
	return NewError("not_found", "Resource not found", op, err)
}

func ErrInvalidArgument(op, message string, args ...interface{}) *Error {
	return NewError(
		"invalid_argument",
		fmt.Sprintf(message, args...),
		op,
		nil,
	)
}

func ErrInternal(op, message string, cause error) *Error {
	return NewError(
		"internal",
		message,
		op,
		cause,
	)
}
