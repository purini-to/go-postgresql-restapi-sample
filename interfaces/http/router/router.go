package router

import (
	"net/http"
	"time"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/api"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/handler"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"

	"github.com/purini-to/go-postgresql-restapi-sample/interfaces/http/middleware"

	"github.com/go-chi/chi"
	cw "github.com/go-chi/chi/middleware"
)

// Router is routeing application.
type Router struct {
	logger  logger.Logger
	logMid  middleware.Logger
	recMid  middleware.Recoverer
	systemH *handler.System
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

	engine.NotFound(func(w http.ResponseWriter, r *http.Request) {
		api.Notfound(w)
	})

	engine.Route("/systems", func(rt chi.Router) {
		rt.Get("/", r.systemH.List)
		rt.Post("/", api.WrapAllowContentType(r.systemH.Create, api.ContentTypeJSON))

		rt.Route("/{id}", func(rt chi.Router) {
			rt.Use(r.systemH.Ctx)
			rt.Get("/", r.systemH.Get)
			rt.Put("/", api.WrapAllowContentType(r.systemH.Put, api.ContentTypeJSON))
			rt.Delete("/", r.systemH.Delete)
		})
	})
}

// NewRouter create application route.
func NewRouter(
	lg logger.Logger,
	lm middleware.Logger,
	rm middleware.Recoverer,
	sh *handler.System,
) *Router {
	return &Router{
		logger:  lg,
		logMid:  lm,
		recMid:  rm,
		systemH: sh,
	}
}
