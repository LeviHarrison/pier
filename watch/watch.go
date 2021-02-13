package watch

import (
	"fmt"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
)

// Watch watches the files
func Watch(files []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Could not create watcher: %v\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	for _, file := range files {
		err = watcher.Add(file)
		if err != nil {
			fmt.Printf("Could not watch file: %v\n", err)
			os.Exit(1)
		}
	}
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Op&fsnotify.Write == fsnotify.Write {
				log.Println("modified file:", event.Name)
			}
		}
	}
}
