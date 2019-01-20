//+build wireinject

package usecase

import (
	"github.com/google/wire"
)

// UsecaseSet is usecase initialize group sets.
var UsecaseSet = wire.NewSet(
	NewSystem,
)
