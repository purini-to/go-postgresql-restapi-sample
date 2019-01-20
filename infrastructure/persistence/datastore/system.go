package datastore

import (
	"context"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"
)

// System is system datastore.
type System struct{}

// FindID find system by id.
func (s *System) FindID(ctx context.Context, db *db.DB, id string) (*model.System, error) {
	var sys model.System
	err := db.First(&sys, "id = ?", id).Error
	return &sys, err
}

// Save save system.
func (s *System) Save(ctx context.Context, db *db.DB, sys *model.System) error {
	if len(sys.ID) == 0 {
		return db.Create(sys).Error
	}

	return db.Update(sys).Error
}

// NewSystem create system datastore.
func NewSystem() *System {
	return &System{}
}
