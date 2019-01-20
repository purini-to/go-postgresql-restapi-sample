package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/purini-to/go-postgresql-restapi-sample/core/logger"

	cw "github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

// Logger write request log.
type Logger func(next http.Handler) http.Handler

// NewLogger prodive request log middleware.
func NewLogger(l *logger.Logger) Logger {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := cw.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()
			defer func() {
				method := r.Method
				url := r.URL
				requestID, _ := r.Context().Value(cw.RequestIDKey).(string)
				l.Info(
					fmt.Sprintf("%s %s %s", method, r.URL, r.Proto),
					zap.Int("status", ww.Status()),
					zap.String("url", url.String()),
					zap.String("proto", r.Proto),
					zap.String("ip", r.RemoteAddr),
					zap.Int("byte", ww.BytesWritten()),
					zap.Duration("latency", time.Since(t1)),
					zap.String("reqID", requestID),
				)
			}()
			next.ServeHTTP(ww, r)
		}
		return http.HandlerFunc(fn)
	}
}
