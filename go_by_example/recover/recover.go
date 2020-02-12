package main

import "fmt"

// Recovering won't work because recover() is called from deferred function
// that is not in the scope of panicked function (lower level of call stack)
func defBar() {
	fmt.Println("defBar() started")

	if r := recover(); r != nil {
		fmt.Println("Program is panicking with value:", r)
	}
}

func defFoo() {
	fmt.Println("defFoo() started")

	defer defBar() // defer call

	fmt.Println("defFoo() done")
}

func normMain() {
	fmt.Println("normMain() started")

	defer defFoo() //defer defFoo call

	panic("HELP!")

	fmt.Println("normMain() done")
}

func defMain() {
	fmt.Println("defMain() started")

	fmt.Println("defMain() done")
}

func main() {
	fmt.Println("main() started")

	defer defMain() // defer call

	normMain() // normal call

	fmt.Println("main() stopped")
}