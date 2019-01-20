package handler

import (
	"context"
	"net/http"

	"github.com/purini-to/go-postgresql-restapi-sample/domain/model"

	"github.com/jinzhu/gorm"

	"github.com/purini-to/go-postgresql-restapi-sample/usecase"

	"go.uber.org/zap"

	"github.com/go-chi/chi"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/api"
)

type key int

const (
	systemKey key = iota
)

// System is system handler.
type System struct {
	l  *logger.Logger
	su *usecase.System
}

// Ctx set system in context.
func (s *System) Ctx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ctx := r.Context()

		sys, err := s.su.FindID(ctx, id)
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				api.Notfound(w)
				return
			}
			api.InternalServerError(w)
			s.l.Error(err.Error(), zap.String("id", id))
			return
		}

		ctx = context.WithValue(ctx, systemKey, sys)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *System) List(w http.ResponseWriter, r *http.Request) {
	systems, err := s.su.Find(r.Context())
	if err != nil {
		api.InternalServerError(w)
		s.l.Error(err.Error())
		return
	}

	api.JSON(w, systems, http.StatusOK)
}

// Get get system by id.
func (s *System) Get(w http.ResponseWriter, r *http.Request) {
	api.JSON(w, r.Context().Value(systemKey), http.StatusOK)
}

// Create create system.
func (s *System) Create(w http.ResponseWriter, r *http.Request) {
	var sys model.System
	if err := api.BindBody(r, &sys); err != nil {
		api.InternalServerError(w)
		s.l.Error(err.Error())
		return
	}

	if err := s.su.Create(r.Context(), &sys); err != nil {
		api.InternalServerError(w)
		s.l.Error(err.Error())
		return
	}

	api.JSON(w, &sys, 200)
}

// NewSystem create system handler.
func NewSystem(
	l *logger.Logger,
	su *usecase.System,
) *System {
	return &System{
		l:  l,
		su: su,
	}
}
