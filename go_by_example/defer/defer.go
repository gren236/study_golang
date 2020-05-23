package main

import (
	"os"
	"fmt"
)

func createFile(p string) *os.File {
	fmt.Println("Creating file...")
	f, err := os.Create(p)
	// If error is thrown - panic!
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing to file...")
	fmt.Fprintln(f, "some dummy data")
}

func closeFile(f *os.File) {
	fmt.Println("Closing file...")
	err := f.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v/n", err)
		os.Exit(1)
	}
}

// Simple example functions
func defFooStart() {
	fmt.Println("defFooStart() executed")
}

func defFooEnd() {
	fmt.Println("defFooEnd() executed")
}

func defMain() {
	fmt.Println("defFooMain() executed")
}

func foo() {
	fmt.Println("foo() executed")

	defer defFooStart() // Defer defFooStart call

	panic("panic from foo()")

	defer defFooEnd() // Defer defFooEnd call

	fmt.Println("foo() done")
}

func main() {
	// Suppose we want to write file...
	f := createFile("/tmp/defer.txt")
	// Defer will be executed at the end of enclosing function
	defer closeFile(f)
	writeFile(f)

	// Simplier example with foo
	fmt.Println("main() started")

	defer defMain() // Defer defFooMain() call

	foo() // Call foo function

	fmt.Println("main() done")
	// Conclusion: deferred functions are called LIFO - following the call stack from inner to outer functions

	// Panic can also be caught by recover() function (See recover.go)
}
