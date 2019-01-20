//+build wireinject

package handler

import (
	"github.com/google/wire"
)

// HandlerSet is handler initialize group sets.
var HandlerSet = wire.NewSet(
	NewPing,
)
