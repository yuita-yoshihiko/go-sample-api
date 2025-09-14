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
		setupPostRoutes(r, dbutil)
	})

	return r
}

func setupUserRoutes(r chi.Router, dbutil db.DBUtils) {
	userUseCase := usecase.NewUserUseCase(
		database.NewUserRepository(dbutil),
	)
	handler := api.NewUserApi(userUseCase)
	r.Get("/users/{id}", handler.Fetch)
	r.Get("/users/{id}/posts", handler.FetchWithPosts)
	r.Post("/users", handler.Create)
	r.Put("/users/{id}", handler.Update)
	r.Delete("/users/{id}", handler.Delete)
}

func setupPostRoutes(r chi.Router, dbutil db.DBUtils) {
	postUseCase := usecase.NewPostUseCase(
		database.NewPostRepository(dbutil),
	)
	handler := api.NewPostApi(postUseCase)
	r.Get("/posts/users/{user_id}", handler.FetchByUserID)
	r.Get("/posts/users/{user_id}/comments", handler.FetchByUserIDWithComments)
}
