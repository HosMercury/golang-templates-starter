package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func (app *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	fs := http.FileServer(http.Dir("./ui/static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fs))

	r.Get("/", app.SnippetsIndex)
	r.Get("/test", app.Test)
	r.Get("/snippets", app.SnippetsIndex)
	r.Get("/snippets/{id}", app.SnippetsView)
	r.Post("/snippets/create", app.SnippetsCreate)

	return r
}
