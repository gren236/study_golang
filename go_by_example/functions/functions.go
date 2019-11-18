package main

import "fmt"

func plus(a, b int, c float32) int {
	return a + b + int(c)
}

func main() {
	fmt.Println(plus(1, 2, 4))
}