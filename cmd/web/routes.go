package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/toastsandwich/bookings/pkg/config"
	"github.com/toastsandwich/bookings/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	// USING PAT
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	mux := chi.NewRouter()
	mux.Use(middleware.Recoverer)
	// mux.Use(midware.WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/whoami", handlers.Repo.WhoAmI)
	return mux
}
