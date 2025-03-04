package middleware_logger

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func MyLogger(log *slog.Logger) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		log = log.With(slog.String("component", "middleware/logger"))
		log.Info("logger middleware enabled")

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			curLog := log.With(slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("user_agent", r.UserAgent()),
				slog.String("request_id", middleware.GetReqID(r.Context())),
			)
			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
			t1 := time.Now()

			defer func() {
				curLog.Info("request completed",
					slog.Int("status", ww.Status()),
					slog.Int("bytes", ww.BytesWritten()),
					slog.Duration("request time", time.Since(t1)),
				)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
