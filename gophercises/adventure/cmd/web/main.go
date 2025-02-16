package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/netletic/adventure"
)

func main() {
	file := flag.String("file", "crawler.json", "JSON file that contains the adventure story.")
	fmt.Println(*file)
	adventure, err := adventure.FromJSON(*file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not open %s: %s", *file, err)
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		chapterView(w, r, adventure)
	})

	log.Print("starting server on :4000\n")

	err = http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
