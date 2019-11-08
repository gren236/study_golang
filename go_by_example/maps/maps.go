package main

import "fmt"

// Foo is a test struct
type Foo struct {
	Bar string
}

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	m["foo"] = 42
	fmt.Println(m)

	// Bool can be used as key type :D
	n := make(map[bool]string)
	n[true] = "foo"
	n[false] = "bar"
	fmt.Println(n)

	// Structs can be used too
	sm := make(map[Foo]bool)
	foo := Foo{Bar: "Baz"}
	sm[foo] = true
	fmt.Println("struct:", sm)
	// Change the struct value
	foo.Bar = "fuzz"
	fmt.Println("struct:", foo, "struct_map:", sm)
	// Struct value is copied inside the map key - so it won't change

	// Trying the save structure value
	bar := Foo{Bar: "Baz"}
	sm[bar] = false
	fmt.Println("struct:", sm)
	// It rewrites the key. So the struct value only has to match.

	fmt.Println("len:", len(n))
	
	// Delete a pair
	delete(n, false)
	fmt.Println(n)

	// Map key call can return 2 values - the second val indicates if this key is present in the map
	_, present := n[false]
	fmt.Println("present:", present)

	// Declare and initialize
	someMap := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(someMap)

	map2d := map[string]map[string]int{"hello": {"foo": 42}}
	fmt.Println(map2d)

	// Difference between make and declaration
	var nMap map[int]int
	// New map built with nil value. Below code cause runtime panic.
	// nMap[1] = 2
	fmt.Println(nMap)
}
