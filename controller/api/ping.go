package api

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Ping is ping api.
type Ping struct {
	logger *zap.Logger
	db     *gorm.DB
}

// Message is message response.
type Message struct {
	Message string `json:"message"`
}

type Consumer struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name,omitempty"`
	Mail string `json:"mail,omitempty"`
}

// Ping echo ping.
func (p *Ping) Ping(c *gin.Context) {
	var consumer Consumer
	p.db.First(&consumer, "id = ?", "aaa")
	c.JSON(http.StatusOK, &consumer)
}

// ProvidePing provide ping api.
func ProvidePing(lg *zap.Logger, db *gorm.DB) *Ping {
	return &Ping{
		logger: lg,
		db:     db,
	}
}
