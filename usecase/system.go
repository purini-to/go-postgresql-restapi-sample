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

// Find find system.
func (s *System) Find(ctx context.Context, q *model.Query) (*[]model.System, error) {
	var m []model.System
	if err := s.sr.Find(ctx, s.db, &m, q); err != nil {
		return nil, err
	}

	return &m, nil
}

// FindID find system by id.
func (s *System) FindID(ctx context.Context, id string) (*model.System, error) {
	var m model.System
	if err := s.sr.FindID(ctx, s.db, id, &m); err != nil {
		return nil, err
	}

	return &m, nil
}

// Create create system.
func (s *System) Create(ctx context.Context, sys *model.System) error {
	return s.sr.Save(ctx, s.db, sys)
}

// Update update system.
func (s *System) Update(ctx context.Context, sys *model.System) error {
	return s.sr.Save(ctx, s.db, sys)
}

// Delete delete system.
func (s *System) Delete(ctx context.Context, sys *model.System) error {
	return s.sr.Delete(ctx, s.db, sys.ID)
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
