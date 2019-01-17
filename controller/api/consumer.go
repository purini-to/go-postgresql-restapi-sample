package api

import (
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ConsumerAPI is consumer api.
type ConsumerAPI struct {
	logger *zap.Logger
	db     *gorm.DB
}

// Consumer is consumer model.
type Consumer struct {
	ID   string `json:"id" gorm:"primary_key"`
	Name string `json:"name,omitempty"`
	Mail string `json:"mail,omitempty"`
}

// Get is get consumer by id.
func (a *ConsumerAPI) Get(c *gin.Context) {
	id := c.Param("id")
	var consumer Consumer
	if err := a.db.First(&consumer, "id = ?", id).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}
	c.JSON(http.StatusOK, &consumer)
}

// ProvideConsumerAPI provide consumer api.
func ProvideConsumerAPI(lg *zap.Logger, db *gorm.DB) *ConsumerAPI {
	return &ConsumerAPI{
		logger: lg,
		db:     db,
	}
}
