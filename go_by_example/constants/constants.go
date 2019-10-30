package main

import (
	"fmt"
	// "math"
)

const foo, bar string = "bar", "baz" + "z"

func main() {
	// Const and var names from global scope can be rewritten
	var foo = 5
	fmt.Println(foo, bar)

	const test = 42 / 8
	fmt.Printf("%T : %v : %v\n", test, test, float64(test))

	// Numeric const can become any numeric data type depending on context
	var test2 float64 = test
	fmt.Printf("%T : %v\n", test2, test2)
}
