package main

import (
	"fmt"
	"sort" // A sort package!
)

func main() {
	// Sorting strings
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// String sorting is per-byte, NOT natural. So 601 comes before 63
	strs2 := []string{"12", "63", "601"}
	sort.Strings(strs2)
	fmt.Println("Numeric strings:", strs2)

	// Sorting ints
	ints := []int{5, 3, 6}
	sort.Ints(ints)
	fmt.Println(ints)

	// Checking if ints are sorted
	s := sort.IntsAreSorted(ints)
	fmt.Println(s)
}
