package main

import "fmt"

func appendString(slice []string, data ...string) []string {
	// Check if data is too long
	oldSliceLen := len(slice)
	newSliceLen := len(slice) + len(data)
	if newSliceLen > cap(slice) {
		newSlice := make([]string, (newSliceLen + 1) * 2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[:newSliceLen]
	copy(slice[oldSliceLen:newSliceLen], data)
	return slice
}

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0], s[1], s[2] = "a", "b", "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	// Appending
	s = append(s, "d")
	fmt.Println("apd:", s)

	// Custom append function
	hello := []string{"a", "b", "c", "d"}
	fmt.Println("apd_custom1:", hello, "cap:", cap(hello))
	hello = appendString(s, "e", "f")
	fmt.Println("apd_custom2:", hello, "cap:", cap(hello))

	// Copying
	c := make([]string, len(s))
	fmt.Printf("%T\n", c[0])
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slicing
	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	fmt.Println()

	// Slice is just a reference to underlying array
	originalArray := [...]int{1, 3, 5}
	copiedArray := originalArray
	copiedArray[1] = 10
	// Array copied by value, so original is not changed
	fmt.Println(originalArray, copiedArray)

	sliceOne := originalArray[:]
	sliceTwo := originalArray[:]
	// Changing array element from the first slice only
	sliceOne[0] = 42

	// Underlying array element changed for both slices
	fmt.Println(sliceOne, sliceTwo)

	fmt.Println()

	// Declaration + initialization
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// Copy by appending vector
	b := []string{"g", "i", "t"}
	a := append([]string{}, b...)
	fmt.Println(a)

	// Reverse
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	fmt.Println(a)
}
