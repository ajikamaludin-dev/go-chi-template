package middleware

import (
	"bytes"
	"net/http"
	"regexp"
	"time"

	chimw "github.com/go-chi/chi/v5/middleware"
	"log/slog"
)

func RequestLogger(log *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			ww := chimw.NewWrapResponseWriter(w, r.ProtoMajor)
			var respBuf bytes.Buffer
			ww.Tee(&respBuf)
			next.ServeHTTP(ww, r)

			if log != nil {
				body := redact(respBuf.String())
				if len(body) > 512 {
					body = body[:512] + "...(truncated)"
				}
				log.Info("http request",
					slog.String("method", r.Method),
					slog.String("path", r.URL.RequestURI()),
					slog.Int("status", ww.Status()),
					slog.Duration("duration", time.Since(start)),
					slog.String("remote", r.RemoteAddr),
					slog.Int("bytes", ww.BytesWritten()),
					slog.String("response", body),
				)
			}
		})
	}
}

var sensitiveRe = regexp.MustCompile(`(?i)("?(password|token|secret|api[_-]?key|access[_-]?token|refresh[_-]?token|authorization)"?\s*[:=]\s*)"?[^",}\s]*"?`)

func redact(s string) string {
	return sensitiveRe.ReplaceAllString(s, `${1}"***"`)
}