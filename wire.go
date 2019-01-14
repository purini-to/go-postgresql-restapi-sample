//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/server"
)

func InitializeApp() (*server.App, func(), error) {
	panic(wire.Build(
		server.ProvideApp,
	))
}
