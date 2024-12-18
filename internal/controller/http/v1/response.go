package v1

import (
	"cinema-seating/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	MessageOK = "Success"

	// Success Codes
	CodeSuccess        = 1000 // Success
	CodeCreatedSuccess = 1001 // Created successfully
)

type response struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Result  any    `json:"result,omitempty"`
}

func IsAnyNil(i any) bool {
	return i == nil
}

// newResponse creates a new instance of Response.
func newResponse(code int, mess string, data any) response {
	r := response{
		Code:    code,
		Message: mess,
	}

	if !IsAnyNil(data) {
		r.Result = data
	}

	return r
}

func bareResponse(c *gin.Context, statusCode, code int, mess string, data any) {
	r := newResponse(code, mess, data)
	c.JSON(statusCode, r)
}

// errorResponse returns an error response.
func errorResponse(c *gin.Context, err error) {
	e := errors.HTTPParseErrors(err)
	r := newResponse(e.Code(), e.Message(), nil)
	c.AbortWithStatusJSON(e.StatusCode(), r)
}

// successResponseWithOK returns a success response with status OK and data. Ex: 200
func successResponseWithOK(c *gin.Context, data any) {
	r := newResponse(CodeSuccess, MessageOK, data)
	c.JSON(http.StatusOK, r)
}

// successResponseWithCreated returns a success response with status Created. Ex: 201
func successResponseWithCreated(c *gin.Context) {
	r := newResponse(CodeCreatedSuccess, MessageOK, nil)
	c.JSON(http.StatusCreated, r)
}
