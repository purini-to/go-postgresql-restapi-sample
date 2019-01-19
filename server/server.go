package server

import (
	"net/http"

	"github.com/purini-to/go-postgresql-restapi-sample/core/config"

	"github.com/go-chi/chi"
	"github.com/purini-to/go-postgresql-restapi-sample/router"
)

// Server is application context
type Server struct {
	router *router.Router
	engine *chi.Mux
	config *config.Config
}

// Start start application
func (a *Server) Start() error {
	a.router.Mapping(a.engine)
	return http.ListenAndServe(":"+a.config.GetString("PORT"), a.engine)
}

// ProvideServer provide Server instance
func ProvideServer(en *chi.Mux, r *router.Router, c *config.Config) *Server {
	return &Server{
		router: r,
		engine: en,
		config: c,
	}
}
