package main

import "fmt"

func vals() (int,int) {
	return 4, 2
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Not going to work - compile error
	// fmt.Printf("Two values - %v and %v", vals())
	
	_, c := vals()
    fmt.Println(c)
}