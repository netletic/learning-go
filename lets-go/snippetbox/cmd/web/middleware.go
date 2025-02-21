package main

import (
	"log/slog"
	"net/http"
)

func commonResponseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		w.Header().Set("Server", "Borant")
		next.ServeHTTP(w, r)
	})
}

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ip     = r.RemoteAddr
			method = r.Method
			host   = r.Host
			uri    = r.URL.RequestURI()
			proto  = r.Proto
		)
		app.logger.Info(
			"received request",
			slog.String("ip", ip),
			slog.String("method", method),
			slog.String("host", host),
			slog.String("uri", uri),
			slog.String("proto", proto),
		)
		next.ServeHTTP(w, r)
	})
}
