package middleware

import (
	"fmt"
	"net/http"
	"time"

	cw "github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

// Logger write request log.
type Logger func(next http.Handler) http.Handler

// ProvideLogger prodive request log middleware.
func ProvideLogger(l *zap.Logger) {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, rq *http.Request) {
			ww := cw.NewWrapResponseWriter(w, rq.ProtoMajor)
			t1 := time.Now()
			defer func() {
				method := rq.Method
				url := rq.URL
				requestID, _ := rq.Context().Value(cw.RequestIDKey).(string)
				l.Info(
					fmt.Sprintf("%s %s %s", method, rq.URL, rq.Proto),
					zap.Int("status", ww.Status()),
					zap.String("url", url.String()),
					zap.String("proto", rq.Proto),
					zap.String("ip", rq.RemoteAddr),
					zap.Int("byte", ww.BytesWritten()),
					zap.Duration("ts", time.Since(t1)),
					zap.String("reqID", requestID),
					zap.String("reqID", requestID),
				)
			}()
			next.ServeHTTP(ww, rq)
		}
		return http.HandlerFunc(fn)
	}
}
