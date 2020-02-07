package main

import (
	"fmt"
	"sort"
)

// Slice to be sorted by length
type byLength []string

// byLength methods - implementing sort.Interface
// Len is the number of elements in the collection.
func (s byLength) Len() int {
	return len(s)
}

// Swap swaps the elements with indexes i and j.
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// Row - Sort structs by field values (e.g. id)
// Struct itself
type Row struct {
	ID int
}

// Slice of structs to be sorted
type byIds []Row

// Same
func (s byIds) Len() int {
	return len(s)
}

// Same
func (s byIds) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// If element i has ID lower than j - it comes before
func (s byIds) Less(i, j int) bool {
	return s[i].ID < s[j].ID
}

func main() {
	// Sort by length
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)

	// Sort structs by id
	structs := []Row{
		{ID: 4},
		{ID: 10},
		{ID: 1},
	}
	sort.Sort(byIds(structs))
	fmt.Println(structs)
}