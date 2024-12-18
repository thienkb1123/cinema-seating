package errors

import (
	"errors"
	"fmt"
	"net/http"
)

const (
	// Internal Server Errors
	CodeInternalServerError = 5001 // Internal server error

	// Bad Request Errors
	CodeBadRequest = 4001 // Bad request

	// NotFound Errors
	CodeNotFound = 4041 // Not found

	// Forbidden Errors
	CodeForbidden = 4031 // Forbidden

	// Request Timeout Errors
	CodeRequestTimeout = 4081 // Request timeout
)

var (
	ErrorBadRequest          = errors.New("bad request")
	ErrorNotFound            = errors.New("not Found")
	ErrorForbidden           = errors.New("forbidden")
	ErrorInternalServerError = errors.New("internal server error")
	ErrorRequestTimeoutError = errors.New("request timeout")
)

// Error struct
type Error struct {
	ErrStatusCode int    `json:"-"`
	ErrCode       int    `json:"code,omitempty"`
	ErrMessage    string `json:"message,omitempty"`
	ErrCauses     any    `json:"-"`
}

// Error  Error() interface method
func (e Error) Error() string {
	return fmt.Sprintf("code: %d - message: %s - errors: %v", e.ErrCode, e.ErrMessage, e.ErrCauses)
}

func (e Error) StatusCode() int {
	return e.ErrStatusCode
}

// Code get error code
func (e Error) Code() int {
	return e.ErrCode
}

func (e Error) Message() string {
	return e.ErrMessage
}

// Causes Error Causes
func (e Error) Causes() any {
	return e.ErrCauses
}

func New(text string) error {
	return errors.New(text)
}

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func As(err error, target any) bool {
	return errors.As(err, &target)
}

// NewError New Error
func NewError(statusCode, code int, message string, causes any) *Error {
	return &Error{
		ErrStatusCode: statusCode,
		ErrCode:       code,
		ErrMessage:    message,
		ErrCauses:     causes,
	}
}

// NewBadRequestError New Bad Request Error
func NewBadRequestError(causes any) *Error {
	return &Error{
		ErrStatusCode: http.StatusBadRequest,
		ErrCode:       CodeBadRequest,
		ErrMessage:    ErrorBadRequest.Error(),
		ErrCauses:     causes,
	}
}

// NewNotFoundError New Not Found Error
func NewNotFoundError(causes any) *Error {
	return &Error{
		ErrStatusCode: http.StatusNotFound,
		ErrCode:       CodeNotFound,
		ErrMessage:    ErrorNotFound.Error(),
		ErrCauses:     causes,
	}
}

// New Forbidden Error
func NewForbiddenError(causes any) *Error {
	return &Error{
		ErrStatusCode: http.StatusForbidden,
		ErrCode:       CodeForbidden,
		ErrMessage:    ErrorForbidden.Error(),
		ErrCauses:     causes,
	}
}

// New Internal Server Error
func NewInternalServerError(causes any) *Error {
	return &Error{
		ErrStatusCode: http.StatusInternalServerError,
		ErrCode:       CodeInternalServerError,
		ErrMessage:    ErrorInternalServerError.Error(),
		ErrCauses:     causes,
	}
}

type withMessage struct {
	cause error
	msg   string
}

func (w *withMessage) Error() string { return w.msg + ": " + w.cause.Error() }
func (w *withMessage) Cause() error  { return w.cause }

// WithMessage annotates err with a new message.
// If err is nil, WithMessage returns nil.
func WithMessage(err error, message string) error {
	if err == nil {
		return nil
	}
	return &withMessage{
		cause: err,
		msg:   message,
	}
}
