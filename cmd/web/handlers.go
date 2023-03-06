package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/go-chi/chi/v5"
)

func (app *application) SnippetsIndex(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
		app.errorLog.Print(err.Error())
	}

	err = ts.ExecuteTemplate(w, "base", ts)

	if err != nil {
		app.errorLog.Print(err.Error())
	}
}

func (app *application) SnippetsView(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	w.Write([]byte(id))
}

func (app *application) SnippetsCreate(w http.ResponseWriter, r *http.Request) {
	// var s models.Snippet

	// err := json.NewDecoder(r.Body).Decode(&s)

	// if err != nil {
	// 	log.Print(err.Error())
	// }

	// log.Print((s))

	// title := r.Body.Get("title")
	// content := r.PostForm.Get("content")

	// fmt.Fprintf(w, "here"+title+"\n"+content+"\n")
	r.ParseForm()
	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")

	stmt := "INSERT INTO snippets (title, content) VALUES ($1, $2);"

	result, err := app.pool.Exec(r.Context(), stmt, title, content)

	if err != nil {
		log.Print(err.Error())
	}

	log.Print(result)
	// fmt.Fprintf(w, "Snippet: %+v", s)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.errorLog.Print("Not Found")
}

func (app *application) Test(w http.ResponseWriter, r *http.Request) {
	var greeting string
	err := app.pool.QueryRow(context.Background(), "select 'Test'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	w.Write([]byte(greeting))
}
