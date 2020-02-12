package main

import (
	"os"
	"fmt"
)

func accessElement(a []int, index int) int {
	if len(a) > index {
		return a[index]
	} else {
		// Panic receives empty interface, which means value of any type!
		panic("out of bounds condition")
	}
}

func foo() {
	fmt.Println("foo is called")
	panic("something went wrong")
	fmt.Println("foo is done")
}

func main() {
	// panic("a problem")

	// We should "panic" when we don't know how to handle errors
	_, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
	}
	
	// Imitating runtime panic with out-of-bounds error
	// a := []int{2, 3, 6}
	// Accessing fourth element, which does not exist
	// fmt.Println(accessElement(a, 4))

	// Let's try calling something after panic is thrown
	fmt.Println("main is called")
	foo()
	fmt.Println("main is done")
}