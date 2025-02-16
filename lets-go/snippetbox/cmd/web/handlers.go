package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Server", "Web")

	files := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.logger.Error(
			err.Error(),
			slog.String("method", r.Method),
			slog.String("uri", r.URL.RequestURI()),
			slog.Int("response_code", http.StatusInternalServerError),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.logger.Error(
			err.Error(),
			slog.String("method", r.Method),
			slog.String("uri", r.URL.RequestURI()),
			slog.Int("response_code", http.StatusInternalServerError),
		)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a form for creating a new snippet..."))
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("save a new snippet..."))
}
