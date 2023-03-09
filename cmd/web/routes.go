package main

import (
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// r.Use(middleware.Recoverer)
	// r.Use(middleware.URLFormat)
	// r.Use(render.SetContentType(render.ContentTypeJSON))

	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	r.Get("/", app.SnippetsIndex)
	r.Get("/test", app.Test)
	r.Get("/snippets", app.SnippetsIndex)
	r.Post("/snippets/create", app.SnippetsCreate)
	r.Get("/snippets/{id}", app.SnippetsView)

	return r
}
