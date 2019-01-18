package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// NoRoute is setup not found router.
type NoRoute struct {
	logger *zap.Logger
}

// Handler handle not found function.
func (n *NoRoute) Handler(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "path not found",
	})
}

// ProvideNoRoute provide no router filter.
func ProvideNoRoute(lg *zap.Logger) *NoRoute {
	return &NoRoute{
		logger: lg,
	}
}
