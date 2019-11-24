package main

import (
	"fmt"
	// "errors"
	"github.com/pkg/errors"
)

// Adding context to error
type argError struct {
	Arg  int
	Prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.Arg, e.Prob)
}

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "can't work with it"}
	}

	return arg + 3, nil
}

func main() {
	fmt.Println(f1(42))

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	// Use type assertion to get underlying custom error instance
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.Arg)
		fmt.Println(ae.Prob)
	}

	// Errors with stack trace
	originalError := errors.New("I'm an error!")

	// Add some context and stack trace
	err := errors.Wrap(originalError, "Context Info")

	// Print with stack trace
	fmt.Printf("with stack trace => %+v \n\n", err)
}
