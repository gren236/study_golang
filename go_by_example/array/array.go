package main

import "fmt"

func main() {
	var a [2 + 3]int
	fmt.Println(a)

	a[5 - 3] = 1
	fmt.Println(a)
	fmt.Println(len(a))

	b := [5]int{1, 2, 2 + 1, 4, 5}
	fmt.Println("dcl:", b)

	// Count elements automatically
	d := []string{"Penn", "Teller"}
	//      ^  slice  v  array
	e := [...]string{"Penn", "Teller"}
	fmt.Printf("%T and %T\n", d, e)
	
	// Two dimensional array
	var twoD [2]([3]int)
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
			// (twoD[i])[j] = i + j // Can also be this way
        }
    }
    fmt.Println("2d: ", twoD)
}