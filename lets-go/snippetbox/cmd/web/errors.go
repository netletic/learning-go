package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		ip     = r.RemoteAddr
		method = r.Method
		host   = r.Host
		uri    = r.URL.EscapedPath()
		proto  = r.Proto
		scheme = "http"
		status = http.StatusInternalServerError // 500
	)

	traceID, ok := r.Context().Value("TraceID").(string)
	if !ok {
		traceID = ""
	}

	if r.TLS != nil {
		scheme = "https"
	}

	app.logger.Error(
		err.Error(),
		slog.String("ip", ip),
		slog.String("scheme", scheme),
		slog.String("method", method),
		slog.String("host", host),
		slog.String("uri", uri),
		slog.String("proto", proto),
		slog.Int("status", status),
		slog.String("trace_id", traceID),
	)

	msg := fmt.Sprintf("%s (X-Trace-Id: %s)", http.StatusText(status), traceID)
	http.Error(w, msg, status)
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
