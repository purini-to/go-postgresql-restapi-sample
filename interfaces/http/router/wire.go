//+build wireinject

package router

import (
	"github.com/google/wire"
)

// RouterSet is router initialize group sets.
var RouterSet = wire.NewSet(
	NewRouter,
)
