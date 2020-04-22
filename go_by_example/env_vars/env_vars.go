package main

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Using Env vars is simple - os package is helpful
	err := os.Setenv("FOO", "1")
	check(err)
	// Check vars
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))
	fmt.Println()

	// Also, we can list all env vars available
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0], ":", pair[1])
	}
	// Try setting env var before calling:
	// BAR=42 go run env_vars.go
}