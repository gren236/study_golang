package main

import "fmt"

func plus(a, b int, c float32) int {
	return a + b + int(c)
}

// You can declare return value
func mult(a, b int) (result int) {
	result = a * b
	return
}

func foo(a, b int) (i int, j int) {
	j, i = a, b
	return
}

func main() {
	fmt.Println(plus(1, 2, 4))
	fmt.Println(mult(5, 5))

	fmt.Println(foo(4, 2))
}