package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/purini-to/go-postgresql-restapi-sample/middleware"

	"go.uber.org/zap"

	"github.com/go-chi/chi"
	cw "github.com/go-chi/chi/middleware"
)

// Router is routeing application.
type Router struct {
	logger *zap.Logger
	logMid *middleware.Logger
}

// Mapping handle for url and method.
func (r *Router) Mapping(engine *chi.Mux) {
	engine.Use(cw.RequestID)
	engine.Use(cw.RealIP)
	engine.Use(r.logMid)
	engine.Use(func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, rq *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					r.logger.Error(fmt.Sprintf("%+v", rvr))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, rq)
		}
		return http.HandlerFunc(fn)
	})
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	engine.Use(cw.Timeout(60 * time.Second))

	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hi"))
	})
}

// ProvideRouter provide application route.
func ProvideRouter(lg *zap.Logger, lm *middleware.Logger) *Router {
	return &Router{
		logger: lg,
		logMid: lm,
	}
}
