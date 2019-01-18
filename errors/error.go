package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Error is error interface.
type Error interface {
	GetCode() int
	GetBody() *gin.H
}

// ErrorResponse is error response.
type ErrorResponse struct {
	Code int
	Body *gin.H
}

// GetCode get status code.
func (e *ErrorResponse) GetCode() int {
	return e.Code
}

// GetBody get response body.
func (e *ErrorResponse) GetBody() *gin.H {
	return e.Body
}

// Option is option setting.
type Option func(*ErrorResponse)

// WithMsg set message.
func WithMsg(msg string) Option {
	return func(e *ErrorResponse) {
		if e.Body == nil {
			e.Body = &gin.H{"message": msg}
			return
		}
		(*e.Body)["message"] = msg
	}
}

// NotFound create not found response.
func NotFound(ops ...Option) *ErrorResponse {
	res := &ErrorResponse{
		Code: http.StatusNotFound,
		Body: &gin.H{"message": "Path not found"},
	}
	for _, o := range ops {
		o(res)
	}
	return res
}

// InternalServerError create internal server error response
func InternalServerError(ops ...Option) *ErrorResponse {
	res := &ErrorResponse{
		Code: http.StatusInternalServerError,
		Body: &gin.H{"message": "Internal server error"},
	}
	for _, o := range ops {
		o(res)
	}
	return res
}
