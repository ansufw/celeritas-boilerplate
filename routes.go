package main

import (
	"boilerplate/public"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (a *application) routes() *chi.Mux {

	a.get("/", a.handlers.Home)

	fileServer := http.FileServer(http.FS(public.Public))
	a.App.Routes.Handle("/public/*", http.StripPrefix("/public/", fileServer))

	return a.App.Routes
}
