package repository

import (
	"context"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
)

// System is interface system.
type System interface {
	FindID(ctx context.Context, db *db.DB, id string) (*model.System, error)
	Save(ctx context.Context, db *db.DB, sys *model.System) error
}
