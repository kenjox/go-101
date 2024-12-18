package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"../../ui/html/pages/layout.tmpl",
		"../../ui/html/partials/nav.tmpl",
		"../../ui/html/pages/home.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)

	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func (app *application) viewSnippet(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "snippet with id %d", id)
}

func (app *application) newSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for creating snippet"))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Saving snippet"))
}
