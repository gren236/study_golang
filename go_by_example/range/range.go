package main

import "fmt"

func main() {
	nums := []string{"foo", "bar", "baz"}
	// For range loop iterates over slices and arrays
	for index, value := range nums {
		fmt.Println(index, value)
	}

	// Indexes and values can be omitted
	for index, _ := range nums {
		fmt.Println(index)
	}
	// But it makes no sense to omit value only, it's better to use this:
	for index := range nums {
		fmt.Println(index)
	}

	fmt.Println()

	// Maps can be iterated with range
	someMap := map[string]string{"foo": "bar", "hello": "world"}
	for k, v := range someMap {
		fmt.Println(k, v)
	}

	fmt.Println()

	// Strings can be iterated with range as well
	// i - char index, c - character code itself (rune)
	for i, c := range "hello" {
		fmt.Println(i, c)
	}
}
