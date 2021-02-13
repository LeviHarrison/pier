package watch

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/leviharrison/pier/parse"
)

// Watch watches the files
func Watch(files []string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Error creating watcher: %v\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	for _, file := range files {
		err := watcher.Add(file)
		if err != nil {
			fmt.Printf("Error watching file %v: %v\n", file, err)
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
				fmt.Println("modified file:", event.Name)
				files, additions = modifyList(files, parse.Partial(event.Name))

				for _, file := range additions {
					err := watcher.Add(file)
					if err != nil {
						fmt.Printf("Error watching file %v: %v\n", file, err)
						os.Exit(1)
					}
				}
			}
		}
	}
}

// modifyFiles returns a modified list of needed files and any additions
func modifyList(files, newFiles []string) ([]string, []string) {
	for _, file := range files {
		for i, new := range newFiles {
			if file == new {
				newFiles[i] = newFiles[len(newFiles)-1]
				newFiles[len(newFiles)-1] = ""
				newFiles = newFiles[:len(newFiles)-1]
			}
		}
	}

	files = append(files, newFiles...)
	return files, newFiles
}

// For some reason, Go won't let me modify the files variable and define the additions variable in the same statement, so I must define it over here
var additions []string
