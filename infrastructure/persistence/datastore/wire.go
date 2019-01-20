//+build wireinject

package datastore

import (
	"github.com/purini-to/go-postgresql-restapi-sample/domain/repository"
	"github.com/google/wire"
)

// DatastoreSet is datastore initialize group sets.
var DatastoreSet = wire.NewSet(
	NewSystem,
	wire.Bind(new(repository.System), new(System))
)
