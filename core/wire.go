//+build wireinject

package core

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/core/config"
	"github.com/purini-to/go-postgresql-restapi-sample/core/db"
	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
)

// CoreSet is core initialize group sets.
var CoreSet = wire.NewSet(
	logger.NewLogger,
	db.NewDB,
	config.NewConfig,
)
