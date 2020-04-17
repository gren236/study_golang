package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

// Checking for errors every time is a good idea
func check(e error) {
	if e != nil {
		panic(e)
	}
}

// As stdlib DOES NOT have any functions to copy files, we can write one ourselves
// copyFs copies src file/directory to dst directory, preserving structure and permissions
func copyFs(src, dst string) error {
	srcStat, err := os.Lstat(src)
	if err != nil {
		return err
	}

	// Create all needed directories
	err = os.MkdirAll(dst, 0775)
	if err != nil {
		return err
	}

	// Check if src is directory
	if srcStat.IsDir() {
		// Get contents of directory
		dirContents, err := ioutil.ReadDir(src)
		if err != nil {
			return err
		}

		// For every entry call copyFs
		var wg sync.WaitGroup
		for _, v := range dirContents {
			path := filepath.Join(src, v.Name())
			dstPath := filepath.Join(dst, srcStat.Name())

			wg.Add(1)
			go func() {
				wg.Done()
				copyFs(path, dstPath)
			}()
		}
		wg.Wait()
	} else {
		// Just a file, copy it's contents
		srcF, err := os.OpenFile(src, os.O_RDONLY, srcStat.Mode())
		if err != nil {
			return err
		}
		defer srcF.Close()
		dstFilePath := filepath.Join(dst, srcStat.Name())
		dstF, err := os.OpenFile(dstFilePath, os.O_CREATE | os.O_WRONLY, 0775)
		if err != nil {
			return err
		}
		defer dstF.Close()

		io.Copy(dstF, srcF)
	}
	return nil
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
	// It's optional to call this func as every change is written to file in real time
	f.Sync()

	// Bufio also provides buffered writers
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %b bytes\n", n4)

	// Use flush to apply all previous operations
	w.Flush()

	// Fprintf can be used to write formatted string to file
	file, err := os.OpenFile("/tmp/testdata", os.O_CREATE | os.O_WRONLY, 0777)
	check(err)

	defer file.Close()

	fmt.Fprintf(file, "File fd is: %v\n", file.Fd())
	file.Sync()

	// We can also rename files!
	err = os.Rename("/tmp/dat2", "/tmp/dat42")
	check(err)

	// Check custom copyFs function
	err = copyFs("/home/gren236/Desktop/imp", "/home/gren236/Desktop/not_useful")
	check(err)
}
