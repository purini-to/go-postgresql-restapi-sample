package repository

import (
	"context"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
)

// System is interface system.
type System interface {
	Find(context.Context, *db.DB, *[]model.System, *model.Query) error
	FindID(context.Context, *db.DB, string, *model.System) error
	Save(context.Context, *db.DB, *model.System) error
	Delete(context.Context, *db.DB, string) error
}
