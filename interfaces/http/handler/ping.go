package handler

import (
	"net/http"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/api"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
)

// Ping is echo ping.
type Ping struct {
	l *logger.Logger
}

// Ping return pong.
func (p *Ping) Ping(w http.ResponseWriter, r *http.Request) {
	api.JSON(w, &api.H{"message": "test"}, http.StatusOK)
}

// NewPing create ping handler.
func NewPing(l *logger.Logger) *Ping {
	return &Ping{
		l: l,
	}
}
