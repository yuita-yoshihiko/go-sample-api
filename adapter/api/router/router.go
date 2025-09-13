package router

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yuita-yoshihiko/go-sample-api/adapter/api"
)

func NewRouter() *chi.Mux {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
	})
	setupHealthRoutes(r)

	return r
}

func setupHealthRoutes(r chi.Router) {
	healthApi := api.NewHealthApi()
	r.Get("/health", healthApi.FetchHealth)
}
