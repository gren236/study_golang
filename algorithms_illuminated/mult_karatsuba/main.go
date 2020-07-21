package main

import (
	"fmt"
	"math"
)

func recIntMult(x, y, n int) int {
	if n == 1 {
		return x * y
	}

	a, b := x / int(math.Pow10(n / 2)), x % int(math.Pow10(n / 2))
	c, d := y / int(math.Pow10(n / 2)), y % int(math.Pow10(n / 2))

	ac := recIntMult(a, c, n/2)
	bd := recIntMult(b, d, n/2)

	pq := recIntMult(a + b, c + d, n/2)
	adbc := pq - ac - bd

	return int(math.Pow10(n)) * ac + int(math.Pow10(n/2)) * adbc + bd
}

func main() {
	fmt.Println(recIntMult(1234, 5678, 4))
}
