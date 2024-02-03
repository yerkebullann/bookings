package main

import (
	"MyFirstApp/pkg/config"
	"MyFirstApp/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routers(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}
