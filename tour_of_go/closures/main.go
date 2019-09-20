package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	prev := 0
	current := 1
	return func() int {
		// Result is previous value
		result := prev
		// Prepare sum with raw prev and current
		sum := prev + current
		// 'Shift' current to prev
		prev = current
		// Make current a summ of raw values
		current = sum
		return result
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
