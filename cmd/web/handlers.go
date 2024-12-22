package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/kenjox/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

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

	data := templateData{
		Snippets: snippets,
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)

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

	snippet, err := app.snippets.GetById(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	files := []string{
		"../../ui/html/pages/layout.tmpl",
		"../../ui/html/partials/nav.tmpl",
		"../../ui/html/pages/view.tmpl",
	}

	tmpl, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	data := templateData{
		Snippet: snippet,
	}

	err = tmpl.ExecuteTemplate(w, "layout", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v", snippet)
}

func (app *application) newSnippetForm(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form for creating snippet"))
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)

	id, err := app.snippets.Insert("Testing", "Testing content", 3)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
	w.Write([]byte("Saving snippet"))
}
