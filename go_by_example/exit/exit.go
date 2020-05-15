package main

import (
	"fmt"
	"os"
)

func main() {
	// Defers will NOT be run when using os.Exit, so this print will never be called
	defer fmt.Println("!")

	// Exit with status 3
	os.Exit(3)
	// Note that unlike C, Go does not take an integer value returned from main as exit status
	// If we want to pass status other from zero, we should use os.Exit()
}
