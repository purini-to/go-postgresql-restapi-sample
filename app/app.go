package app

import (
	"github.com/purini-to/go-postgresql-restapi-sample/server"
)

// App is application instance.
type App struct {
	server *server.Server
}

// Start start application.
func (app *App) Start() error {
	return app.server.Start()
}

// ProvideApp provide app instance.
func ProvideApp(s *server.Server) *App {
	return &App{
		server: s,
	}
}
