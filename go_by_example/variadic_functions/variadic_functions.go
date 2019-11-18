package main

import "fmt"

func sum(nums ...int) {
	fmt.Printf("Recieved args are packed to a %T : %v\n", nums, nums)
}

// Variadics should be a final parameter
func foo(a string, b ...string) {
	fmt.Println(a, b)
}

// func bar(a, b int) {
func bar(a ...int) {
	fmt.Println(a)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	foo("hello", "world", "foo")

	nums := []int{1, 15}
	// Receiving function must have a variadic argument to get packed args
	// Because it's basically the same slice passed, no new slices created inside a function
	// So it won't work with regular args as expected. Only variadics
	bar(nums...)
}
