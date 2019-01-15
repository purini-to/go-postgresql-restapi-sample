package router

import (
	"github.com/gin-gonic/gin"
	"github.com/purini-to/go-postgresql-restapi-sample/controller/api"
	"github.com/purini-to/go-postgresql-restapi-sample/middleware"
)

// Router is routeing application.
type Router struct {
	corsMid *middleware.Cors
	ping    *api.Ping
}

// Mapping handle for url and method.
func (r *Router) Mapping(engine *gin.Engine) {
	engine.Use(r.corsMid.Handler())

	engine.GET("ping", r.ping.Ping)
}

// ProvideRouter provide application route.
func ProvideRouter(cm *middleware.Cors, ping *api.Ping) *Router {
	return &Router{
		corsMid: cm,
		ping:    ping,
	}
}
