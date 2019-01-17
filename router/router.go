package router

import (
	"github.com/gin-gonic/gin"
	"github.com/purini-to/go-postgresql-restapi-sample/controller/api"
	"github.com/purini-to/go-postgresql-restapi-sample/middleware"
)

// Router is routeing application.
type Router struct {
	corsMid     *middleware.Cors
	logMid      *middleware.Log
	consumerAPI *api.ConsumerAPI
}

// Mapping handle for url and method.
func (r *Router) Mapping(engine *gin.Engine) {
	engine.Use(r.logMid.Handler())
	engine.Use(gin.Recovery())
	engine.Use(r.corsMid.Handler())

	engine.GET("consumers/:id", r.consumerAPI.Get)
}

// ProvideRouter provide application route.
func ProvideRouter(cm *middleware.Cors, lg *middleware.Log, conAPI *api.ConsumerAPI) *Router {
	return &Router{
		corsMid:     cm,
		logMid:      lg,
		consumerAPI: conAPI,
	}
}
