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

func (s *System) Find(ctx context.Context) (*[]model.System, error) {
	var m []model.System
	if err := s.sr.Find(ctx, s.db, &m); err != nil {
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
