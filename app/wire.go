package app

import (
	"github.com/google/wire"
)

// AppSet is app initialize group sets.
var AppSet = wire.NewSet(
	ProvideApp,
)
