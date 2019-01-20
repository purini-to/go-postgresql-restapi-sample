package http

import (
	"github.com/go-chi/chi"
)

// NewEngine create web framework engine.
func NewEngine() *chi.Mux {
	return chi.NewRouter()
}
