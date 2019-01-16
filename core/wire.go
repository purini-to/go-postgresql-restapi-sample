//+build wireinject

package core

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
)

// CoreSet is core initialize group sets.
var CoreSet = wire.NewSet(
	logger.ProvideLogger,
)
