package router

import (
	"time"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/handler"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/middleware"

	"github.com/go-chi/chi"
	cw "github.com/go-chi/chi/middleware"
)

// Router is routeing application.
type Router struct {
	logger    *logger.Logger
	logMid    middleware.Logger
	recMid    middleware.Recoverer
	pingH     *handler.Ping
	notfoundH *handler.Notfound
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

	engine.NotFound(r.notfoundH.Handler)

	engine.Get("/ping", r.pingH.Ping)
}

// NewRouter create application route.
func NewRouter(
	lg *logger.Logger,
	lm middleware.Logger,
	rm middleware.Recoverer,
	ph *handler.Ping,
	nh *handler.Notfound,
) *Router {
	return &Router{
		logger:    lg,
		logMid:    lm,
		recMid:    rm,
		pingH:     ph,
		notfoundH: nh,
	}
}
