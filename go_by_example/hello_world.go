package main

import (
	"fmt"

	"github.com/study_golang/pkg/string"
)

// `go run` - run as interpreter
// `go build` - compile to file
func main() {
	fmt.Println(string.Reverse("Hello World!"))
}
