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
	mux := app.Routes()
	logger.Info("starting server", slog.String("addr", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
