package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func Md5All(root string) (map[string][md5.Size][]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// TODO
	})
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
