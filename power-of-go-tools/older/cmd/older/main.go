package main

import (
	"fmt"
	"os"
	"time"

	"github.com/netletic/older"
)

const Usage = `Usage: older DURATION

Lists all files older than DURATION in the tree rooted at the current directory.

Example: older 24h
(lists all files last modified more than 24 hours ago)`

func main() {
	if len(os.Args) < 2 {
		fmt.Println(Usage)
		os.Exit(0)
	}
	age, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fsys := os.DirFS("/Users/jarnotimmermans/Downloads")
	paths := older.Than(fsys, age)
	for _, p := range paths {
		fmt.Println(p)
	}
}
