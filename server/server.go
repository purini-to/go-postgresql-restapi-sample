package server

import (
	"github.com/gin-gonic/gin"
	"github.com/purini-to/go-postgresql-restapi-sample/router"
)

// Server is application context
type Server struct {
	router *router.Router
	engine *gin.Engine
}

// Start start application
func (a *Server) Start() error {
	a.router.Mapping(a.engine)
	return a.engine.Run()
}

// ProvideServer provide Server instance
func ProvideServer(en *gin.Engine, r *router.Router) *Server {
	return &Server{
		router: r,
		engine: en,
	}
}
