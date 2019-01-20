package usecase

import (
	"context"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
	"github.com/purini-to/go-postgresql-restapi-sample/domain/repository"
)

// System is usecase system.
type System struct {
	l  *logger.Logger
	db *db.DB
	sr repository.System
}

// FindID find system by id.
func (s *System) FindID(ctx context.Context, id string) (*model.System, error) {
	return s.sr.FindID(ctx, s.db, id)
}

// Create create system.
func (s *System) Create(ctx context.Context, sys *model.System) error {
	return s.sr.Save(ctx, s.db, sys)
}

// NewSystem create system usecase.
func NewSystem(
	l *logger.Logger,
	db *db.DB,
	sr repository.System,
) *System {
	return &System{
		l:  l,
		db: db,
		sr: sr,
	}
}
