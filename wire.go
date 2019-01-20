//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/core/config"
	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/handler"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/middleware"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/router"
)

func InitializeServer() (*http.Server, func(), error) {
	panic(wire.Build(
		http.HttpSet,
		config.ConfigSet,
		logger.LoggerSet,
		handler.HandlerSet,
		middleware.MiddleSet,
		router.RouterSet,
	))
}
