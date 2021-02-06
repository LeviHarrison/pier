package main

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/leviharrison/pier/parse"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Could not create watcher: %v\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	files := parse.Files("cmd/pier")
}

func watch(watcher fsnotify.Watcher) {
	for {
		select {}
	}
}
