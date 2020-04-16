package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// TempFile creates and opens for r/w a temporary file in the default OS location
	f, err := ioutil.TempFile("", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name())

	// OS will likely clean up mp files after some time, but it's better to do it explicitly
	defer os.Remove(f.Name())

	// Write some data to file
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	// Check written info
	data := make([]byte, 4)
	_, err = f.ReadAt(data, 0)
	fmt.Println("Tmp file contains:", data)

	// Also temp dir can be created
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)

	// Cleanup dir too
	defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}