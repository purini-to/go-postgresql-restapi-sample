package router

import (
	"net/http"
	"time"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/middleware"

	"github.com/go-chi/chi"
	cw "github.com/go-chi/chi/middleware"
)

// Router is routeing application.
type Router struct {
	logger *logger.Logger
	logMid middleware.Logger
	recMid middleware.Recoverer
}

// Mapping handle for url and method.
func (r *Router) Mapping(engine *chi.Mux) {
	engine.Use(cw.RequestID)
	engine.Use(cw.RealIP)
	engine.Use(r.logMid)
	engine.Use(r.recMid)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	engine.Use(cw.Timeout(60 * time.Second))

	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
}

// NewRouter create application route.
func NewRouter(
	lg *logger.Logger,
	lm middleware.Logger,
	rm middleware.Recoverer,
) *Router {
	return &Router{
		logger: lg,
		logMid: lm,
		recMid: rm,
	}
}
