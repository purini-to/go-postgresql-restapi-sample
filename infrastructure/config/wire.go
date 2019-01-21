//+build wireinject

package config

import (
	"github.com/spf13/viper"
	cg "github.com/purini-to/go-postgresql-restapi-sample/core/config"
	"github.com/google/wire"
)

// ConfigSet is config initialize group sets.
var ConfigSet = wire.NewSet(
	NewConfig,
	wire.Bind(new(cg.Config), new(viper.Viper))
)
