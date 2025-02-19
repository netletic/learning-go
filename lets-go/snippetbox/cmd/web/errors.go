package main

import (
	"log/slog"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		host   = r.Host
		uri    = r.URL.RequestURI()
		status = http.StatusInternalServerError // 500
	)

	app.logger.Error(
		err.Error(),
		slog.String("method", method),
		slog.String("host", host),
		slog.String("uri", uri),
		slog.Int("status", status),
	)

	http.Error(w, http.StatusText(status), status)
}

// func (app *application) clientError(w http.ResponseWriter, r *http.Request, status int) {
// 	var (
// 		method = r.Method
// 		host   = r.Host
// 		uri    = r.URL.RequestURI()
// 	)

// 	app.logger.Error(
// 		http.StatusText(status),
// 		slog.String("method", method),
// 		slog.String("host", host),
// 		slog.String("uri", uri),
// 		slog.Int("status", status),
// 	)

// 	http.Error(w, http.StatusText(status), status)
// }
