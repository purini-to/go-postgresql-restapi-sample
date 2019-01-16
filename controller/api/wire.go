//+build wireinject

package api

import (
	"github.com/google/wire"
)

// APISet is api initialize group sets.
var APISet = wire.NewSet(
	ProvidePing,
)
