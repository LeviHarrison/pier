package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"strings"

	"golang.org/x/mod/modfile"
)

func main() {
	pkg := getPackage()
	imprts := getImports("./cmd/pier")

	for _, i := range imprts {
		if strings.Contains(i, pkg) {
			fmt.Println(i)
		}
	}
}

func getPackage() string {
	data, err := ioutil.ReadFile("go.mod")

	if err != nil {
		fmt.Printf("Error reading modfile: %v", err)
		os.Exit(1)
	}

	mod, err := modfile.ParseLax("go.mod", data, nil)
	if err != nil {
		fmt.Printf("Error parsing modfile: %v", err)
		os.Exit(1)
	}

	return mod.Module.Mod.Path
}

func getImports(dir string) []string {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ImportsOnly)
	if err != nil {
		fmt.Printf("Error parsing files: %v", err)
		os.Exit(1)
	}

	imports := []string{}
	for _, pkg := range pkgs {
		for _, f := range pkg.Files {
			for _, i := range f.Imports {
				imports = append(imports, i.Path.Value)
			}
		}
	}

	return imports
}
