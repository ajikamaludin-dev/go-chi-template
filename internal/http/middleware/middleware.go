package middleware

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Timeout(d time.Duration) func(http.Handler) http.Handler {
	return middleware.Timeout(d)
}

func Recoverer(next http.Handler) http.Handler {
	return middleware.Recoverer(next)
}
