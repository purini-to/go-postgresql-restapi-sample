//+build wireinject

package logger

import (
	"go.uber.org/zap"
	lg "github.com/purini-to/go-postgresql-restapi-sample/core/logger"
	"github.com/google/wire"
)

// LoggerSet is logger initialize group sets.
var LoggerSet = wire.NewSet(
	NewLogger,
	wire.Bind(new(lg.Logger), new(zap.Logger))
)
