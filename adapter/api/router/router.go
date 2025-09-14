package router

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/yuita-yoshihiko/go-sample-api/adapter/api"
	"github.com/yuita-yoshihiko/go-sample-api/adapter/database"
	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
	"github.com/yuita-yoshihiko/go-sample-api/usecase"
)

func SetupRoutes(dbutil db.DBUtils) *chi.Mux {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	r := chi.NewRouter()
	r.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		setupUserRoutes(r, dbutil)
	})
	setupHealthRoutes(r)

	return r
}

func setupHealthRoutes(r chi.Router) {
	healthApi := api.NewHealthApi()
	r.Get("/health", healthApi.FetchHealth)
}

func setupUserRoutes(r chi.Router, dbutil db.DBUtils) {
	userUseCase := usecase.NewUserUseCase(
		database.NewUserRepository(dbutil),
	)
	handler := api.NewUserApi(userUseCase)
	r.Get("/users/{id}", handler.Fetch)
}
