//+build wireinject

package db

import (
	"github.com/google/wire"
)

// DBSet is database initialize group sets.
var DBSet = wire.NewSet(
	NewDB,
)
