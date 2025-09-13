package main

import (
	"net/http"

	"github.com/yuita-yoshihiko/go-sample-api/adapter/api/router"
)

func main() {
	// if err := env.Parse(&config.Conf); err != nil {
	// 	panic(err)
	// }

	if err := http.ListenAndServe(":80", router.NewRouter()); err != nil {
		panic(err)
	}
}
