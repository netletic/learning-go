package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// this is a handler -> handlers execute app logic and write http resp. headers & bodies
func home(w http.ResponseWriter, r *http.Request) {
	_ = r
	w.Write([]byte("hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("display a specific snippet with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("display a form for creating a new snippet..."))
}

func main() {
	fmt.Println("hello world!")

	// this is a servemux -> servemuxes store a mapping between URL routing patterns
	// and the corresponding handler for that routing pattern
	mux := http.NewServeMux()

	// servemux treats pattern ending in a "/" as a catch-all, it'll match /anything/foo/bar
	// i.e. /**
	// changing it to "/{$} enforces an exact match and stops the wildcard behavior"
	mux.HandleFunc("GET /{$}", home) // restrict this route to exact matches on "/" only

	// because /snippet/view doesn't end with a slash, it's an exact match.
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)

	// additional notes about servemux features:
	// request patterns are automatically sanitized, /foo/bar/..//baz will be
	// automatically be sent a 301 Permanent Redirect to /foo/baz instead

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	// errors returned by ListenAndServe() are always non-nil
	log.Fatal(err)
}
