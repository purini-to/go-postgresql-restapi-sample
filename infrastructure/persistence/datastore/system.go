package datastore

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/purini-to/go-postgresql-restapi-sample/domain/repository"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"
)

func with(o *repository.Option, d *db.DB) *gorm.DB {
	return d.Where(o.System).Limit(o.Limit).Offset(o.Offset)
}

// System is system datastore.
type System struct{}

func (s *System) Find(ctx context.Context, db *db.DB, m *[]model.System, o ...repository.O) error {
	op := repository.DefaultOption()

	for _, o := range o {
		o(op)
	}

	d := with(op, db)

	return d.Find(m).Error
}

// FindID find system by id.
func (s *System) FindID(ctx context.Context, db *db.DB, id string, m *model.System) error {
	return db.First(m, "id = ?", id).Error
}

// Save save system.
func (s *System) Save(ctx context.Context, db *db.DB, m *model.System) error {
	if len(m.ID) == 0 {
		return db.Create(m).Error
	}

	return db.Update(m).Error
}

// NewSystem create system datastore.
func NewSystem() *System {
	return &System{}
}
