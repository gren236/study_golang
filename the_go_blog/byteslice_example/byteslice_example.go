package main

import (
	"fmt"
	bs "github.com/gren236/study_golang/pkg/byteslice"
)

func main() {
	slice := bs.ByteSlice("Hello")
	slice.Append([]byte(", World!"))
	fmt.Printf("%s\n", slice)

	fmt.Fprintf(&slice, "Foobar")
	fmt.Printf("%s\n", slice)
}
