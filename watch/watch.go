package watch

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/leviharrison/pier"
	"github.com/leviharrison/pier/parse"
)

var additions []string

// Watch watches the files
func Watch(targets pier.Targets) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Printf("Error creating watcher: %v\n", err)
		os.Exit(1)
	}
	defer watcher.Close()

	for _, target := range targets {
		for _, file := range target.Files {
			err := watcher.Add(file)
			if err != nil {
				fmt.Printf("Error watching file %v: %v\n", file, err)
				os.Exit(1)
			}
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

				targets := findTargets(event.Name, targets)
				for _, target := range targets {
					target.Files, additions = modifyList(target.Files, parse.Partial(event.Name))

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

// findTargets finds what targets depend on a given file
func findTargets(found string, targets pier.Targets) pier.Targets {
	result := []*pier.Target{}
	for _, target := range targets {
		for _, file := range target.Files {
			if found == file {
				result = append(result, target)
				continue
			}
		}
	}

	return result
}
