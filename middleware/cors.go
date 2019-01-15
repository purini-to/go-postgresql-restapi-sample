package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Cors is setup cors.
type Cors struct {
	logger *zap.Logger
}

// Handler handle cors function.
func (c *Cors) Handler() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AddAllowMethods("DELETE")
	config.AllowAllOrigins = true
	c.logger.Info("initialize cors middleware.")
	return cors.New(config)
}

// ProvideCors provide cors filter.
func ProvideCors(lg *zap.Logger) *Cors {
	return &Cors{
		logger: lg,
	}
}
