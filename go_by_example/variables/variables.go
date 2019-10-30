package main

import "fmt"

func main() {
	var a = "initial"
	var __a = "initial2"
	fmt.Println(a, __a)

	// Any unicode symbols can be used as var names
	var λ, 丂 string = "initial3", "initial4"
	fmt.Println(λ, 丂)

	// This produces an assignment mismatch error!
	// var λ, 丂 string = "initial3", "initial4", "test"

	// All common default var init values
	var test1 bool
	fmt.Println(test1)
	var test2 string
	fmt.Println(test2)
	var test3 uint
	fmt.Println(test3)
	var test4 float64
	fmt.Println(test4)
	var test5 complex128
	fmt.Println(test5)
	var test6 byte  // Alias for uint8
	fmt.Printf("%T : %v\n", test6, test6)
	var test7 rune  // Alias for int32
	fmt.Printf("%T : %v\n", test7, test7)
	var test8 uintptr
	fmt.Printf("%T : %v\n", test8, test8)
	var test9 *int
	fmt.Printf("%T : %v\n", test9, test9)

	// Type inferring
	var b int32
	c, d := b, b
	fmt.Printf("%T, %T, %T\n", b, c, d)

	foo := 11.0 + c
	// Will not work, "trruncated to integer error"
	// foo := 11.0 + c
	fmt.Println(foo)
}