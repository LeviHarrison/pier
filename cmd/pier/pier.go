package main

import (
	"github.com/leviharrison/pier/parse"
	"github.com/leviharrison/pier/watch"
)

func main() {
	files := parse.Files("cmd/pier")

	watch.Watch(files)
}
