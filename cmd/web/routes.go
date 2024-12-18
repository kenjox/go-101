package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Serve static files
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /{$}", app.home) // Restrict this route to exact matches on / only.”
	mux.HandleFunc("GET /snippet/view/{id}", app.viewSnippet)
	mux.HandleFunc("GET /snippet/create", app.newSnippetForm)
	mux.HandleFunc("POST /snippet/create", app.createSnippet)

	return mux
}
