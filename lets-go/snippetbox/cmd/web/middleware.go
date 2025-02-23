package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/justinas/nosurf"
)

func injectSecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Security-Policy", "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com")
		w.Header().Set("Referrer-Policy", "origin-when-cross-origin")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "0")
		next.ServeHTTP(w, r)
	})
}

func generateTraceID() string {
	bytes := make([]byte, 8) // 8 bytes (16 hex chars)
	_, err := rand.Read(bytes)
	if err != nil {
		return "0000000000000000" // Fallback in case of an error
	}
	return hex.EncodeToString(bytes)
}

type responseWriter struct {
	http.ResponseWriter
	traceID string
	status  int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.Header().Set("X-Trace-ID", rw.traceID)
	rw.ResponseWriter.WriteHeader(code)
}

func (app *application) logHTTPExchange(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var (
			ip     = r.RemoteAddr
			method = r.Method
			host   = r.Host
			uri    = r.URL.EscapedPath()
			proto  = r.Proto
			scheme = "http"
		)

		traceID, ok := r.Context().Value(traceIDContextKey).(string)
		if !ok {
			traceID = ""
		}

		if r.TLS != nil {
			scheme = "https"
		}

		// inject traceID into the response header
		rw := &responseWriter{ResponseWriter: w, traceID: traceID, status: 0}
		next.ServeHTTP(rw, r)
		app.logger.Info(
			"http_exchange",
			slog.Time("start", start),
			slog.Time("end", time.Now()),
			slog.Int64("durationMs", time.Since(start).Milliseconds()),
			slog.String("ip", ip),
			slog.String("method", method),
			slog.String("scheme", scheme),
			slog.String("host", host),
			slog.String("uri", uri),
			slog.String("proto", proto),
			slog.Int("status", rw.status),
			slog.String("trace_id", traceID),
		)
	})
}

func (app *application) injectTracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), traceIDContextKey, generateTraceID())
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuthentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAuthenticated(r) {
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		w.Header().Add("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := app.sessionManager.GetInt(r.Context(), "authenticatedUserID")
		if id == 0 {
			next.ServeHTTP(w, r)
			return
		}
		exists, err := app.users.Exists(id)
		if err != nil {
			app.serverError(w, r, err)
			return
		}
		if exists {
			ctx := context.WithValue(r.Context(), isAuthenticatedContextKey, true)
			r = r.WithContext(ctx)
		}
		next.ServeHTTP(w, r)
	})
}

func noSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   true,
	})

	return csrfHandler
}
