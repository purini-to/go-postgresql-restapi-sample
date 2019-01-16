package server

import "github.com/gin-gonic/gin"

// ProvideEngine provide web framework engine.
func ProvideEngine() *gin.Engine {
	return gin.New()
}
