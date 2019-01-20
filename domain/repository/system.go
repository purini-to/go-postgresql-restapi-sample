package repository

import (
	"context"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/db"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"
)

// System is interface system.
type System interface {
	Find(ctx context.Context, db *db.DB, m *[]model.System, o ...O) error
	FindID(ctx context.Context, db *db.DB, id string, m *model.System) error
	Save(ctx context.Context, db *db.DB, sys *model.System) error
}

type O func(o *Option)

type Option struct {
	*model.System
	Limit  uint16
	Offset uint16
}

func WithSystem(s *model.System) O {
	return func(o *Option) {
		o.System = s
	}
}

func WithLimit(l uint16) O {
	return func(o *Option) {
		o.Limit = l
	}
}

func WithOffset(of uint16) O {
	return func(o *Option) {
		o.Offset = of
	}
}

func DefaultOption() *Option {
	return &Option{
		Limit:  100,
		Offset: 0,
	}
}
