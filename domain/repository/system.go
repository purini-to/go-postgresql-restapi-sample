package repository

import (
	"context"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
)

// System is interface system.
type System interface {
	Find(ctx context.Context, db *db.DB, m *[]model.System, q *model.Query) error
	FindID(ctx context.Context, db *db.DB, id string, m *model.System) error
	Save(ctx context.Context, db *db.DB, sys *model.System) error
}
