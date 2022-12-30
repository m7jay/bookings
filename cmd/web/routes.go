package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/m7jay/bookings/pkg/config"
	"github.com/m7jay/bookings/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	// return mux

	mux := chi.NewRouter()
	//	mux.Use(middleware.Recoverer, middleware.CleanPath, middleware.Logger, middleware.RequestID)
	//	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
