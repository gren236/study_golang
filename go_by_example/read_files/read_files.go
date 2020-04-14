package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// Checking for errors every time is a good idea
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	fmt.Printf("%T - %s", dat, dat)

	// To have more control over read process, call os.Open
	f, err := os.Open("/tmp/dat")
	check(err)
	fmt.Printf("%#v\n", f)

	// Read bytes from file using buffer
	b1 := make([]byte, 5)
	// Read writes bytes to buffer, returns how many bytes were read
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, b1[:n1])

	// You can use Seek to search for a certain position and read from there
	o2, err := f.Seek(3, 0)
	check(err)
	fmt.Printf("%T - %v\n", o2, o2)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%s\n", b2[:n2])
	// All of that can be replaced with ReadAt method call (reading with offset)

	// The io package provides helpful functions for file reading
	// ex. ReadAtLeast()
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, b3)

	// To read from file start again use f.Seek(0, 0)
	_, err = f.Seek(0, 0)
	check(err)

	// More complex and useful buffer can be used with bufio package
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", b4)

	f.Close()

	// Working with directories
	f1, err := os.Open("/home/gren236")
	check(err)
	fmt.Printf("%T - %#v\n", f1, f1)
	files, err := f1.Readdir(5)
	check(err)
	fmt.Printf("%#v\n", files)
	for _, v := range files {
		fmt.Printf("%s\n", v.Name())
	}
}
