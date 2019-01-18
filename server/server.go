package server

import (
	"net/http"

	"github.com/spf13/viper"

	"github.com/go-chi/chi"
	"github.com/purini-to/go-postgresql-restapi-sample/router"
)

// Server is application context
type Server struct {
	router *router.Router
	engine *chi.Mux
	config *viper.Viper
}

// Start start application
func (a *Server) Start() error {
	a.router.Mapping(a.engine)
	return http.ListenAndServe(":"+a.config.GetString("PORT"), a.engine)
}

// ProvideServer provide Server instance
func ProvideServer(en *chi.Mux, r *router.Router, c *viper.Viper) *Server {
	return &Server{
		router: r,
		engine: en,
		config: c,
	}
}
