package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println()

	// Multiple assignment inside loops
	for j, k := 7, 8; j <= 9; j, k = j+1, k+1 {
		fmt.Println(j, k)
	}

	fmt.Println()

	// While loop analogue
	i = 0
	for {
		if i >= 2 {
			break
		}

		i++
		fmt.Println(i)
	}

	fmt.Println()

	var n = 25
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}

		fmt.Println(n)
	}
	// For has it's own scope
	fmt.Println(n)

	fmt.Println()

	var p = 0
	// Last expression will evaluate even if comparison is not true anymore (the last time)
	for p := &p; *p <= 5; *p++ {
		fmt.Println(*p)
	}
	// For has it's own scope
	fmt.Println(p)
}
