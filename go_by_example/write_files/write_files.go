package main

import (
	"bufio"
	"fmt"
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
	// Basic file write
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/tmp/dat1", d1, 0775)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// WriteString is also available
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Call sync to write buffered changes to persistent storage
	f.Sync()

	// Bufio also provides buffered writers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %b bytes\n", n4)

	// Use flush to apply all previous operations
	w.Flush()
}
