//+build wireinject

package logger

import (
	"github.com/google/wire"
)

// LoggerSet is logger initialize group sets.
var LoggerSet = wire.NewSet(
	NewLogger,
)
