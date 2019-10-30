package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("go" + "lang")
	fmt.Printf("Type: %T\n", 14 + 13.5)

	fmt.Println(7/3)
	fmt.Println(7/3.0)

	// fmt.Println(5/0) - compiler error!

	fmt.Println(+7/-3)
	fmt.Println(7%3)

	// Mod operator cannot be used with float type!
	// fmt.Println(7%3.0)
	// But we can use it with "math" package
	fmt.Println(math.Mod(7.5, 3.02))

	// Precendence according to real math
	fmt.Println(3 - 5 + 6)

	// Boolean operators work with Bool Data Type only!
	fmt.Println(false || true)
	// fmt.Println(true || 0) - not working
}