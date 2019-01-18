package middleware

import (
	"github.com/purini-to/go-postgresql-restapi-sample/errors"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NoRoute is setup not found router.
type NoRoute struct {
	logger *zap.Logger
}

// Handler handle not found function.
func (n *NoRoute) Handler(c *gin.Context) {
	res := errors.NotFound()
	c.AbortWithStatusJSON(res.GetCode(), res.GetBody())
}

// ProvideNoRoute provide no router filter.
func ProvideNoRoute(lg *zap.Logger) *NoRoute {
	return &NoRoute{
		logger: lg,
	}
}
