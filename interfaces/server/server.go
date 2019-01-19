package server

import (
	"net/http"

	"github.com/purini-to/go-postgresql-restapi-sample/core/config"

	"github.com/go-chi/chi"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/router"
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

// NewServer create Server instance
func NewServer(en *chi.Mux, r *router.Router, c *config.Config) *Server {
	return &Server{
		router: r,
		engine: en,
		config: c,
	}
}
