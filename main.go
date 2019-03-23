package main

import (
	"log"
	"net/http"

	"go-project/internal/module/ah"
	"go-project/internal/pkg/config"
	"go-project/internal/pkg/httpserver"
	"go-project/internal/pkg/middleware"
	"go-project/internal/pkg/sqldb"
)

func main() {
	config.Load("config")

	err := sqldb.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sqldb.Close()

	mux := http.NewServeMux()

	// mount modules
	ah.Mount(mux)

	httpserver.Server{
		Addr: ":8080",
		Handler: middleware.Chain(
			sqldb.Middleware(),
		)(mux),
	}.ListenAndServe()
}
