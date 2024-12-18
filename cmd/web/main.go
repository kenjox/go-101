package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Serve static files
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", app.home) // Restrict this route to exact matches on / only.”
	router.HandleFunc("GET /snippet/view/{id}", app.viewSnippet)
	router.HandleFunc("GET /snippet/create", app.newSnippetForm)
	router.HandleFunc("POST /snippet/create", app.createSnippet)

	logger.Info("Server running on port", "addr", *addr)

	if error := http.ListenAndServe(*addr, router); error != nil {
		logger.Error(error.Error())
		os.Exit(1)
	}
}
