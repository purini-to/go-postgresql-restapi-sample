package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AbortNotFound is not found response.
func AbortNotFound(c *gin.Context, msg string) {
	if msg == "" {
		msg = "path not found"
	}
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
		"message": "path not found",
	})
}
