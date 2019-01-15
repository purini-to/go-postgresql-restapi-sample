package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Ping is ping api.
type Ping struct {
	logger *zap.Logger
}

// Message is message response.
type Message struct {
	Message string `json:"message"`
}

// Ping echo ping.
func (p *Ping) Ping(c *gin.Context) {
	p.logger.Info("ping")
	c.JSON(http.StatusOK, &Message{Message: "ping"})
}

// ProvidePing provide ping api.
func ProvidePing(lg *zap.Logger) *Ping {
	return &Ping{
		logger: lg,
	}
}
