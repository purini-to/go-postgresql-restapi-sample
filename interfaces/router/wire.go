//+build wireinject

package router

import (
	"github.com/google/wire"
)

// RouterSet is routeing initialize group sets.
var RouterSet = wire.NewSet(
	NewRouter,
)
