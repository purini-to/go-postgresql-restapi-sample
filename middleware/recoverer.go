package middleware

import (
	"fmt"
	"net/http"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"
)

// Recoverer write request log.
type Recoverer func(next http.Handler) http.Handler

// ProvideRecoverer prodive request log middleware.
func ProvideRecoverer(l *logger.Logger) Recoverer {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rvr := recover(); rvr != nil {
					l.Error(fmt.Sprintf("%+v", rvr))
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
