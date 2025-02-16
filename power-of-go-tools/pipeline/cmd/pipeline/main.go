package main

import "github.com/netletic/pipeline"

func main() {
	pipeline.FromString("hello, world\n").Stdout()
}
