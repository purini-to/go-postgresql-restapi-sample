package handler

import (
	"context"
	"net/http"

	"github.com/imdario/mergo"

	"github.com/asaskevich/govalidator"

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
	l  logger.Logger
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

// List get system list
func (s *System) List(w http.ResponseWriter, r *http.Request) {
	var q model.Query
	api.BindQuery(r, &q)
	_, err := govalidator.ValidateStruct(&q)
	if err != nil {
		api.BadRequest(w, api.WithError(err)...)
		return
	}

	systems, err := s.su.Find(r.Context(), &q)
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
		api.BadRequest(w, api.WithError(err)...)
		return
	}
	if _, err := govalidator.ValidateStruct(&sys); err != nil {
		api.BadRequest(w, api.WithError(err)...)
		return
	}

	if err := s.su.Create(r.Context(), &sys); err != nil {
		api.InternalServerError(w)
		s.l.Error(err.Error(), zap.Any("system", &sys))
		return
	}

	api.JSON(w, &sys, http.StatusOK)
}

// Put update system.
func (s *System) Put(w http.ResponseWriter, r *http.Request) {
	var sys model.System
	if err := api.BindBody(r, &sys); err != nil {
		api.BadRequest(w, api.WithError(err)...)
		return
	}
	if _, err := govalidator.ValidateStruct(&sys); err != nil {
		api.BadRequest(w, api.WithError(err)...)
		return
	}
	v := r.Context().Value(systemKey)
	sysDB, ok := v.(*model.System)
	if !ok {
		api.InternalServerError(w)
		s.l.Error("context value can't cast type", zap.Any("src", v), zap.String("type", "model.System"))
		return
	}
	if err := mergo.MergeWithOverwrite(sysDB, sys); err != nil {
		api.InternalServerError(w)
		s.l.Error("db system can't merge request system",
			zap.Any("db system", &sysDB), zap.Any("request system", &sys))
		return
	}

	if err := s.su.Update(r.Context(), sysDB); err != nil {
		api.InternalServerError(w)
		s.l.Error(err.Error(), zap.Any("system", &sys))
		return
	}

	api.JSON(w, sysDB, http.StatusOK)
}

// Delete delete system.
func (s *System) Delete(w http.ResponseWriter, r *http.Request) {
	v := r.Context().Value(systemKey)
	sys, ok := v.(*model.System)
	if !ok {
		api.InternalServerError(w)
		s.l.Error("context value can't cast type", zap.Any("src", v), zap.String("type", "model.System"))
		return
	}

	if err := s.su.Delete(r.Context(), sys); err != nil {
		api.InternalServerError(w)
		s.l.Error(err.Error(), zap.Any("system", &sys))
		return
	}

	api.JSON(w, "", http.StatusNoContent)
}

// NewSystem create system handler.
func NewSystem(
	l logger.Logger,
	su *usecase.System,
) *System {
	return &System{
		l:  l,
		su: su,
	}
}
