//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/infrastructure/config"
	"github.com/purini-to/go-postgresql-restapi-sample/infrastructure/logger"
	"github.com/purini-to/go-postgresql-restapi-sample/infrastructure/persistence/datastore"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/handler"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/middleware"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/router"
	"github.com/purini-to/go-postgresql-restapi-sample/usecase"
)

func InitializeServer() (*http.Server, func(), error) {
	panic(wire.Build(
		http.HttpSet,
		db.DBSet,
		config.ConfigSet,
		logger.LoggerSet,
		handler.HandlerSet,
		middleware.MiddleSet,
		router.RouterSet,
		datastore.DatastoreSet,
		usecase.UsecaseSet,
	))
}
