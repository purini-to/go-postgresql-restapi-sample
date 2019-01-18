package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/purini-to/go-postgresql-restapi-sample/errors"
)

// HandleFunc is return response function.
type HandleFunc func(c *gin.Context) (*Response, errors.Error)

// Response is response.
type Response struct {
	Code int
	Body interface{}
}

// WrapHandle wrap handle func.
// return gin.HandleFunc
func WrapHandle(h HandleFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		res, err := h(c)
		if err != nil {
			c.AbortWithStatusJSON(err.GetCode(), err.GetBody())
			return
		}
		c.JSON(res.Code, res.Body)
	}
}

// OK create ok resuponse.
func OK(body interface{}) *Response {
	return &Response{
		Code: http.StatusOK,
		Body: body,
	}
}
