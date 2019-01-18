package server

import (
	"github.com/go-chi/chi"
)

// ProvideEngine provide web framework engine.
func ProvideEngine() *chi.Mux {
	return chi.NewRouter()
}
