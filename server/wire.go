//+build wireinject

package server

import (
	"github.com/google/wire"
)

// ServerSet is server initialize group sets.
var ServerSet = wire.NewSet(
	ProvideServer,
	ProvideEngine,
)
