package httpx

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	chimw "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	"github.com/aji/chi-go-template/internal/config"
	"github.com/aji/chi-go-template/internal/http/middleware"
	"github.com/aji/chi-go-template/internal/observability"
)

type RouterDeps struct {
	Cfg     *config.Config
	Logger  *observability.Channels
	Metrics bool
}

func NewRouter(d RouterDeps) http.Handler {
	r := chi.NewRouter()
	r.Use(chimw.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(chimw.RequestID)
	r.Use(chimw.Compress(5))
	if d.Logger != nil {
		r.Use(middleware.RequestLogger(d.Logger.App))
	}

	if d.Cfg.App.Env == "development" {
		r.Use(cors.Handler(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders: []string{"Accept", "Authorization", "Content-Type"},
		}))
	}

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	r.Get("/readyz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ready"))
	})

	if d.Metrics {
		r.Handle("/metrics", observability.MetricsHandler())
	}

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"message":"pong"}`))
		})
	})

	return r
}
