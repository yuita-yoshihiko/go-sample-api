package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/caarlos0/env/v11"
	"github.com/yuita-yoshihiko/go-sample-api/adapter/api/router"
	"github.com/yuita-yoshihiko/go-sample-api/config"
	"github.com/yuita-yoshihiko/go-sample-api/infrastructure/db"
)

func main() {
	if err := env.Parse(&config.Conf); err != nil {
		panic(err)
	}

	// slogの初期化
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	database, err := db.Init()
	if err != nil {
		panic(err)
	}
	dbManager := db.NewDBManager(database)
	dbUtil := db.NewDBUtil(database)

	r := router.SetupRoutes(dbUtil, dbManager)

	if err := http.ListenAndServe(":80", r); err != nil {
		panic(err)
	}
}
