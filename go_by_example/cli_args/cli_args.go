package main

import (
	"fmt"
	"os"
)

func main() {
	// Get raw args
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// Get individual args
	arg := os.Args[2]

	fmt.Printf("%T: %v\n", argsWithProg, argsWithProg)
	fmt.Printf("%T: %v\n", argsWithoutProg, argsWithoutProg)
	fmt.Printf("%T: %v\n", arg, arg)
}
