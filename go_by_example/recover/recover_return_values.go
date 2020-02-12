package main

import "fmt"

// access an element of a slice by index
func accessElement(a []int, index int) int {
	// Inline call of deferred function with recovering code
	// IIFE - immidiately invoked function expression
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered but nothing to do here!")
		}
	}()

	return a[index]
}

// access an element of a slice by index (with fallback)
func accessElementFallback(a []int, index int) (v int) {
	// Using named return values, we can override returned element from recovering expression
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Program recovered with fallback provided!")
			v = a[len(a)-1] // Set returned value to the last element from slice
		}
	}()

	v = a[index]

	return
}

func main() {
	a := []int{1, 2, 3}

	// Access 3rd element
	fmt.Println(accessElement(a, 2))

	// Access 4th element (out of bound)
	fmt.Println(accessElement(a, 3))
	// This call returned 0, as it's the default value for int return type of function

	// Recover with fallback value:
	fmt.Println(accessElementFallback(a, 4))
}
