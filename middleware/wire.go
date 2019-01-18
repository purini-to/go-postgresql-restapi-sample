//+build wireinject

package middleware

import (
	"github.com/google/wire"
)

// CoreSet is core initialize group sets.
var MiddleSet = wire.NewSet(
	ProvideLogger,
)
