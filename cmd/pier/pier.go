package main

import (
	"fmt"
	"os"

	"github.com/leviharrison/pier"
	"github.com/leviharrison/pier/parse"
	"github.com/leviharrison/pier/watch"
)

var targets pier.Targets

const help = `Pier enables intelligent reload for Docker

Usage: pier directory/of/your/main/package your.Dockerfile build/context`

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		fmt.Println(help)
		os.Exit(0)
	}
	if args[0] == "help" || args[0] == "Help" {
		fmt.Println(help)
		os.Exit(0)
	}

	register(args)

	parse.All(targets)
	watch.Watch(targets)
}

func register(args []string) {
	target := &pier.Target{MainDir: args[0], Dockerfile: args[1], Context: args[2]}

	targets = append(targets, target)
	fmt.Printf("Registered new target %v with main located in %v and the context %v\n", target.MainDir, target.Dockerfile, target.Context)
}
