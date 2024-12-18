package errors

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// HTTP Parser of error string messages returns RestError
func HTTPParseErrors(err error) *Error {
	switch {
	case errors.Is(err, context.DeadlineExceeded):
		return NewError(http.StatusRequestTimeout, CodeRequestTimeout, ErrorRequestTimeoutError.Error(), err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewError(http.StatusBadRequest, CodeBadRequest, ErrorBadRequest.Error(), err)
	case strings.Contains(err.Error(), "required"):
		fmt.Println("required")
		return NewError(http.StatusBadRequest, CodeBadRequest, err.Error(), err)
	case strings.Contains(err.Error(), "EOF"):
		return NewError(http.StatusBadRequest, CodeBadRequest, ErrorBadRequest.Error(), err)
	default:
		var restErr *Error
		if errors.As(err, &restErr) {
			return restErr
		}
		return NewInternalServerError(err)
	}
}
