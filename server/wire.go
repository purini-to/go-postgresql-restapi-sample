//+build wireinject

package server

import (
	"github.com/google/wire"
)

func InitializeApp() (*App, func(), error) {
	panic(wire.Build(ProvideApp))
}
