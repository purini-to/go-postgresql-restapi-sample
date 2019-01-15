package middleware

import (
	"github.com/google/wire"
)

// MiddleSet is middleware initialize group sets.
var MiddleSet = wire.NewSet(
	ProvideCors,
)
