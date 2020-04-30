package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

type result struct {
	path string
	sum [md5.Size]byte
	err error
}

// First pipeline - schedules file reads and sum calculating
func sumFiles(done <-chan struct{}, root string) (<-chan result, <-chan error) {
	// For each regular file, start a goroutine that sums the file and sends
	// the result on c. Send the result of the walk on errc
	c := make(chan result)
	errc := make(chan error, 1)
	go func() {
		var wg sync.WaitGroup
		err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			wg.Add(1)
			// Goroutine for each file read
			go func() {
				data, err := ioutil.ReadFile(path)
				// Listen to 2 channels, abort if done received
				select {
				case c <- result{path, md5.Sum(data), err}:
				case <-done:
				}
				wg.Done()
			}()
			// Abort the walk if done received
			select {
			case <-done:
				return errors.New("Walk operation canceled")
			default:
				return nil
			}
		})
		// Walk has returned so all calls to wg.Add are done. Start a goroutine to close c once all the sends are done.
		go func() {
			wg.Wait()
			close(c)
		}()
		// errc is buffered, so no select here
		errc <- err
	}()
	return c, errc
}

// Receives the values from channel "c". Returns early on error, closing done with defer.
func Md5All(root string) (map[string][md5.Size]byte, error) {
	done := make(chan struct{})
	defer close(done)

	c, errc := sumFiles(done, root)

	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}
	// Check for an error from buffered channel
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

func main() {
	m, err := Md5All(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)

	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}
}
