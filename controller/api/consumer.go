package api

import (
	"github.com/purini-to/go-postgresql-restapi-sample/controller"

	"github.com/purini-to/go-postgresql-restapi-sample/errors"

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
func (a *ConsumerAPI) Get(c *gin.Context) (*controller.Response, errors.Error) {
	id := c.Param("id")
	var consumer Consumer
	err := a.db.First(&consumer, "id = ?", id).Error
	if gorm.IsRecordNotFoundError(err) {
		return nil, errors.NotFound()
	} else if err != nil {
		a.logger.Error(err.Error(), zap.String("id", id))
		return nil, errors.InternalServerError()
	}
	return controller.OK(&consumer), nil
}

// ProvideConsumerAPI provide consumer api.
func ProvideConsumerAPI(lg *zap.Logger, db *gorm.DB) *ConsumerAPI {
	return &ConsumerAPI{
		logger: lg,
		db:     db,
	}
}
