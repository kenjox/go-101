package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "Http network address")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	router := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("../../ui/static/"))

	// Serve static files
	router.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", home) // Restrict this route to exact matches on / only.”
	router.HandleFunc("GET /snippet/view/{id}", viewSnippet)
	router.HandleFunc("GET /snippet/create", newSnippetForm)
	router.HandleFunc("POST /snippet/create", createSnippet)

	logger.Info("Server running on port", "addr", *addr)

	if error := http.ListenAndServe(*addr, router); error != nil {
		logger.Error(error.Error())
		os.Exit(1)
	}
}
