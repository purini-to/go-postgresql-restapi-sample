//+build wireinject

package middleware

import (
	"github.com/google/wire"
)

// ProvideLogger is middleware initialize group sets.
var MiddleSet = wire.NewSet(
	ProvideLogger,
	ProvideRecoverer,
)
