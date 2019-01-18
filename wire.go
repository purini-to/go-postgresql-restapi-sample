//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/app"
	"github.com/purini-to/go-postgresql-restapi-sample/core"
	"github.com/purini-to/go-postgresql-restapi-sample/middleware"
	"github.com/purini-to/go-postgresql-restapi-sample/router"
	"github.com/purini-to/go-postgresql-restapi-sample/server"
)

func InitializeApp() (*app.App, func(), error) {
	panic(wire.Build(
		app.AppSet,
		server.ServerSet,
		core.CoreSet,
		router.RouterSet,
		middleware.MiddleSet,
	))
}
