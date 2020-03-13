package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Return random integer: 1 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100), "\n")
	fmt.Println()

	// Return random float: 0.0 <= f < 1.0
	fmt.Println(rand.Float64())
	// This can be used to generate other ranges, for example 5.0 <= f < 10.0
	fmt.Print((rand.Float64() * 5) + 5, ",")
	fmt.Println((rand.Float64() * 5) + 5)

	// By default produced numbers are deterministic, the same number produced every time with the same seed.
	// So it's not safe to use for secret random numbers (use crypto/rand)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	fmt.Print(r1.Intn(100), ",")
	fmt.Println(r1.Intn(100))

	// If same seed is used, the output would be the same
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Println(r2.Intn(100))
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Println(r3.Intn(100))
}
