//+build wireinject

package controller

import (
	"github.com/google/wire"
	"github.com/purini-to/go-postgresql-restapi-sample/controller/api"
)

// ControllerSet is controller initialize group sets.
var ControllerSet = wire.NewSet(
	api.ProvideConsumerAPI,
)
