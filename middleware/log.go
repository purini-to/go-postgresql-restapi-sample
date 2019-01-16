package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/purini-to/go-postgresql-restapi-sample/context"
	"go.uber.org/zap"
)

// Log is setup log.
type Log struct {
	logger *zap.Logger
}

// Handler handle log function.
func (c *Log) Handler() gin.HandlerFunc {
	c.logger.Info("initialize log middleware.")
	if !context.IsProduction() {
		return gin.Logger()
	}

	return func(ctx *gin.Context) {
		// Start timer
		start := time.Now()
		path := ctx.Request.URL.Path
		raw := ctx.Request.URL.RawQuery
		if raw != "" {
			path = path + "?" + raw
		}

		// Process request
		ctx.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)
		clientIP := ctx.ClientIP()
		method := ctx.Request.Method
		statusCode := ctx.Writer.Status()
		err := ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		msg := fmt.Sprintf("%s %s", method, path)

		c.logger.Info(msg,
			zap.String("url", path),
			zap.Time("time", end),
			zap.Duration("latency", latency),
			zap.String("method", method),
			zap.String("clientIP", clientIP),
			zap.Int("statusCode", statusCode),
			zap.String("err", err),
		)
	}
}

// ProvideLog provide log filter.
func ProvideLog(lg *zap.Logger) *Log {
	return &Log{
		logger: lg,
	}
}
