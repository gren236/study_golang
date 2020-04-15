package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Scanning stdin input to buffer
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Text() returns the current token (next line from the input)
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// Check for errors during scan. EOF is expected, not an error
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}