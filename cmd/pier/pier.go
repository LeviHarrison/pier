package main

import (
	"github.com/leviharrison/pier/parse"
	"github.com/leviharrison/pier/watch"
)

func main() {
	files := parse.All("cmd/pier")

	watch.Watch(files)
}
