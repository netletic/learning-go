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
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// logging
	slogHandlerOpts := &slog.HandlerOptions{
		// AddSource: true, // adds line number where error occurred
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, slogHandlerOpts))

	app := &application{
		logger: logger,
	}

	// mux
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
