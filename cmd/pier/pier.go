package main

import (
	"fmt"

	"github.com/leviharrison/pier/parse"
)

func main() {
	for _, i := range parse.Files("cmd/pier") {
		fmt.Println(i)
	}
}
