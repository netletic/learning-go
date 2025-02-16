package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/netletic/adventure"
)

func chapterView(w http.ResponseWriter, r *http.Request, adventure adventure.Adventure) {
	chapterQueryParam := r.URL.Query().Get("chapter")
	if chapterQueryParam == "" {
		chapterQueryParam = "intro"
	}
	chapter, ok := adventure[chapterQueryParam]
	if !ok {
		http.Error(w, "Chapter not found", http.StatusNotFound)
		return
	}

	ts, err := template.ParseFiles("./ui/html/pages/chapter.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, chapter)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}
