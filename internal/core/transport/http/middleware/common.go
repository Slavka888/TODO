package middleware

import (
	"context"
	"net/http"
	"time"
	"uuid"

	"github.com/Slavka888/TODO/internal/core/logger"
	"github.com/Slavka888/TODO/internal/core/transport/http/response"
	"go.uber.org/zap"
)

const requestIDHeader = "X-Request-ID"

func RequestID() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)
			if requestID == "" {
				requestID = uuid.New().String()
			}

			r.Header.Set(requestIDHeader, requestID)
			w.Header().Set(requestIDHeader, requestID)

			next.ServeHTTP(w, r)
		})
	}
}

func Logger(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), "log", l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			httpResponseHandler := response.NewHTTPResponseHandler(log, w)

			defer func() {
				if p := recover(); p != nil {
					httpResponseHandler.PanicResponse(p, "during handling http request got unexpected fatal error")
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}


func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			rw := response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				"incoming HTTP request",
				zap.Time("time", before.UTC()),
			)

			next.ServeHTTP(rw, r)

			log.Debug(
				"done HTTP request",
				zap.Int("status code", rw.GetStatusCodeOrPanic()),
				zap.Duration("time", time.Since(before)),
			)
		})
	}
}