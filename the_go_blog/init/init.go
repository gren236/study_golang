package main

import "fmt"

var (
	foo = "hello!"
	bar = "World!"
)

// Init executes after all variables are initialized
func init() {
	if foo == "hello!" {
		fmt.Println(foo)
	}
	if bar == "World!" {
		fmt.Println(bar)
	}
}

// There can be multiple init files in the source file
func init() {
	fmt.Println("Another init!")
}

func main() {
	fmt.Println("Main executes after init")
}
