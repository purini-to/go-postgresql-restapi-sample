//+build wireinject

package http

import (
	"github.com/google/wire"
)

// HttpSet is http initialize group sets.
var HttpSet = wire.NewSet(
	NewServer,
	NewEngine,
)
