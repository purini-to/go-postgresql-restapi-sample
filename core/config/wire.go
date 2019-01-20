//+build wireinject

package config

import (
	"github.com/google/wire"
)

// ConfigSet is config initialize group sets.
var ConfigSet = wire.NewSet(
	NewConfig,
)
