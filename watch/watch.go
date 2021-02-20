package watch

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/leviharrison/pier"
	"github.com/leviharrison/pier/image"
	"github.com/leviharrison/pier/parse"
	"github.com/radovskyb/watcher"
)

var additions []string

// Watch watches the files
func Watch(targets pier.Targets) {
	w := watcher.New()
	defer w.Close()

	go func() {
		for {
			select {
			case event := <-w.Event:
				if event.Op == watcher.Write {
					targets := findTargets(event.Path, targets)

					for _, target := range targets {
						target.Files, additions = modifyList(target.Files, parse.Partial(event.Path))
						for _, file := range additions {
							err := w.Add(file)
							if err != nil {
								fmt.Printf("Error watching file %v: %v\n", file, err)
								os.Exit(1)
							}
						}

						image.Build(target)
					}
				}
			case err := <-w.Error:
				fmt.Printf("Error with watcher: %v\n", err)
				os.Exit(1)
			case <-w.Closed:
				fmt.Println("Watcher closed")
			}
		}
	}()

	for _, target := range targets {
		for _, file := range target.Files {
			err := w.Add(file)
			if err != nil {
				fmt.Printf("Error watching file %v: %v\n", file, err)
				os.Exit(1)
			}
		}
	}

	err := w.Start(time.Millisecond * 100)
	if err != nil {
		fmt.Printf("Error with watcher: %v\n", err)
		os.Exit(1)
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
out:
	for _, target := range targets {
		for _, file := range target.Files {
			abs, err := filepath.Abs(found)
			if err != nil {
				fmt.Printf("Error finding absolute path for %v: %v\n", file, err)
				os.Exit(1)
			}
			if found == abs {
				for _, t := range result {
					if target == t {
						continue out
					}
				}
				result = append(result, target)
			}
		}
	}

	return result
}
