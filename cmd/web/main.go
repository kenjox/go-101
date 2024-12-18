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

	logger.Info("Server running on port", "addr", *addr)

	if error := http.ListenAndServe(*addr, app.routes()); error != nil {
		logger.Error(error.Error())
		os.Exit(1)
	}
}
