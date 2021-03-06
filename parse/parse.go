package parse

import (
	"fmt"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"

	"github.com/leviharrison/pier"
	"golang.org/x/mod/modfile"
)

// All gets all the files needed starting from the main package
func All(targets pier.Targets) {
	mod := getMod()

	for _, target := range targets {
		files := getFiles(target.MainDir, mod, "main")
		target.Files = removeDuplicates(files)
	}
}

// Partial get all the files needed starting from a certain file
func Partial(path string) []string {
	mod := getMod()

	imports, files := getImportsFile(path)

	for _, i := range imports {
		// Remove the parenthesis
		i = i[1 : len(i)-1]

		// Check if the first part of the import statement contains the current package
		if len(i) > len(mod) && i[:len(mod)] == mod {
			files = append(files, getFiles("."+i[len(mod):], mod, getPkg(i))...)
		}
	}

	return removeDuplicates(files)
}

func removeDuplicates(files []string) []string {
	m := make(map[string]bool)

	for _, file := range files {
		if _, found := m[file]; !found {
			m[file] = true
		}
	}

	result := []string{}
	for file := range m {
		result = append(result, file)
	}

	return result
}

// dir is the directory, mod is the module name, pkg is the package of the files we're looking for
func getFiles(dir, mod, pkg string) []string {
	imports, files := getImports(dir, pkg)

	for _, i := range imports {
		// Remove the parenthesis
		i = i[1 : len(i)-1]

		// Check if the first part of the import statement contains the current package
		if len(i) > len(mod) && i[:len(mod)] == mod {
			files = append(files, getFiles("."+i[len(mod):], mod, getPkg(i))...)
		}
	}

	return files
}

// Get the name of the package from the import path
func getPkg(path string) string {
	// Making sure that we don't go out of range
	if len(path) <= 2 {
		fmt.Printf("Invalid package name: %v", path)
		os.Exit(1)
	}

	for i := len(path) - 1; i >= 0; i-- {
		if string(path[i]) == "/" {
			return path[i+1:]
		}
	}

	fmt.Printf("No package found for import: %v\n", path)
	os.Exit(1)
	return ""
}

func getMod() string {
	data, err := ioutil.ReadFile("go.mod")

	if err != nil {
		fmt.Printf("Error reading modfile: %v\n", err)
		os.Exit(1)
	}

	mod, err := modfile.ParseLax("go.mod", data, nil)
	if err != nil {
		fmt.Printf("Error parsing modfile: %v\n", err)
		os.Exit(1)
	}

	return mod.Module.Mod.Path
}

// Get the imports of all the files in a directory
func getImports(dir, p string) ([]string, []string) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, nil, parser.ImportsOnly)
	if err != nil {
		fmt.Printf("Error parsing files: %v\n", err)
		os.Exit(1)
	}

	imports := []string{}
	files := []string{}
	for _, pkg := range pkgs {
		for name, file := range pkg.Files {
			// Check if the file is from the package we're looking for
			if file.Name.Name == p {
				files = append(files, name)
				for _, i := range file.Imports {
					imports = append(imports, i.Path.Value)
				}
			}
		}
	}

	return imports, files
}

// Get the imports of a file
func getImportsFile(path string) ([]string, []string) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
	if err != nil {
		fmt.Printf("Error parsing files: %v\n", err)
		os.Exit(1)
	}

	imports := []string{}
	files := []string{}
	for _, i := range file.Imports {
		imports = append(imports, i.Path.Value)
	}

	return imports, files
}
