package main

import (
	"fmt"
	"github.com/gren236/study_golang/pkg/string"
)

// `go run` - run as interpreter
// `go build` - compile to file
// Curly brace should stay on the same line
func main() {
	fmt.Println(string.Reverse("Hello World!"));
	// Yes, golang has semicolons!
	fmt.Print("Hi, "); fmt.Print("Mark!\n")
}
