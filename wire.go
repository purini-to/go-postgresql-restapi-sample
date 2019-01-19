//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/core"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/middleware"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/router"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/server"
)

func InitializeServer() (*server.Server, func(), error) {
	panic(wire.Build(
		server.ServerSet,
		core.CoreSet,
		router.RouterSet,
		middleware.MiddleSet,
	))
}
