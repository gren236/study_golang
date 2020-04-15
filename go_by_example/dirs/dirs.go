package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Helper to create new empty file
func createEmptyFile(name string) error {
	return ioutil.WriteFile(name, []byte(""), 0644)
}

func main() {
	err := os.Mkdir("subdir", 0775)
	check(err)

	// For temporary directories it's better to remove them at the end with defer
	defer os.RemoveAll("subdir")

	createEmptyFile("subdir/file1")

	// Create recursive directories
	err = os.MkdirAll("subdir/parent/child", 0775)
	check(err)

	// Add some test files
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// ReadDir can be used to get directory contents
	c, err := ioutil.ReadDir("subdir/parent")
	check(err)
	// Iterate over files returned
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Mode(), "	", entry.Name())
	}

	// Change current working directory (like "cd" command)
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Get and list current category
	c, err = ioutil.ReadDir(".")
	check(err)
	fmt.Println("Listing current dir (subdir/parent/child)")
	for _, entry := range c {
		fmt.Println(" ", entry.Mode(), "	", entry.Name())
	}

	// Get back to original directory
	err = os.Chdir(filepath.Join("..", "..", ".."))
	check(err)

	// We can walk a directory recursively!
	fmt.Println("Visiting subdir")
	err = filepath.Walk("subdir", visit)
}

func visit(p string, info os.FileInfo, err error) error {
	if err != nil {
		return nil
	}

	// Count of slashes in path means levels of recursive output
	fmt.Println(strings.Repeat("   ", strings.Count(p, "/")), info.Mode(), "   ", info.Name())

	return nil
}
