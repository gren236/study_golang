package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Use strconv package to parse numbers
	// ParseFloat can take precision bits size as an argument
	// If 32 bit size is passed, then float64 is returned, but it can be easily
	// converted to float32 without number changed.
	f, _ := strconv.ParseFloat("1.2345", 32)
	fmt.Printf("%T - %v\n", f, f)
	// Only number is allowed to be inside the string
	// Returns an error on wrong input
	f1, e := strconv.ParseFloat("1.2345sdasndas", 32)
	fmt.Printf("%T - %v\n", f1, f1)
	fmt.Println(e)
	// Special symbols can be used in strings - +Inf, -Inf, NaN (Non case-sensitive)
	f2, _ := strconv.ParseFloat("+Inf", 64)
	fmt.Printf("%T - %#v\n", f2, f2)
	f3, _ := strconv.ParseFloat("NaN", 64)
	fmt.Printf("%T - %#v\n", f3, f3)

	// ParseInt also receives bits. Base can be 0, 2,... to 36. 0 base means it's declared in string, for example:
	// 0b for 2 base.
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Printf("%T - %v\n", i, i)
	i1, _ := strconv.ParseInt("0b110011010001", 0, 32)
	fmt.Printf("%T - %v\n", i1, i1)

	// Atoi is a convenience function for basic base-10 int parsing.
	k, _ := strconv.Atoi("135")
	fmt.Printf("%T - %v\n", k, k)
}
