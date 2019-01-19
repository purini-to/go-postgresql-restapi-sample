//+build wireinject

package middleware

import (
	"github.com/google/wire"
)

// NewLogger is middleware initialize group sets.
var MiddleSet = wire.NewSet(
	NewLogger,
	NewRecoverer,
)
