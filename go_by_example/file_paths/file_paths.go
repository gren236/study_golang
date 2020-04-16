package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Parse and construct portable file paths
func main() {
	// Joins should be used to construct system-independent file paths
	p := filepath.Join("home", "gren236", ".bashrc")
	fmt.Println(p)

	// Join also analyses file paths and removes redundant things
	fmt.Println(filepath.Join("home", "dir1", "..", "dir1"))
	fmt.Println(filepath.Join("home", "/gren236/"))

	// Dir() returns the dir path of a file, Base() returns the filename only
	fmt.Println(filepath.Dir(p))
	fmt.Println(filepath.Base(p))

	// Check if path is absolute
	fmt.Println(filepath.IsAbs("/tmp/dir"))
	fmt.Println(filepath.IsAbs("tmp/dir"))

	// Get file extension
	filename := "config.json"
	ext := filepath.Ext(filename)
	fmt.Println(ext)

	// Get only file name, without extension
	fmt.Println(strings.TrimSuffix(filename, ext))

	// Rel() find the relative path between base and target
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

	// Using Glob for filepath patterns
	files, err := filepath.Glob("/home/gren236/Pictures/*.jpg")

	for _, file := range files {
		f, _ := os.Stat(file)

		fmt.Println(f.Mode(), "	", f.Name())
	}
}
