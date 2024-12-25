package main

import (
	"html/template"
	"path/filepath"

	"github.com/kenjox/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

func newTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("../../ui/html/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)

		files := []string{
			"../../ui/html/pages/layout.tmpl",
			page,
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			return nil, err
		}

		tmpl, err = tmpl.ParseGlob("../../ui/html/partials/*.tmpl")
		if err != nil {
			return nil, err
		}

		cache[filename] = tmpl
	}

	return cache, nil
}
