package main

import (
	"fmt"
	"os"

	"github.com/leviharrison/pier/parse"
	"github.com/leviharrison/pier/watch"
)

const help = `Pier enables intelligent reload for Docker

Usage: pier /directory/of/your/main/package your.Dockerfile
`

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println(help)
		os.Exit(0)
	}
	if args[0] == "help" || args[0] == "Help" {
		fmt.Println(help)
		os.Exit(0)
	}

	files := parse.All(args[0])

	watch.Watch(files)
}
