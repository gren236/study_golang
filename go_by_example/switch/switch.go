package main

import (
	"fmt"
	"time"
)

type Test []byte

func (t Test) String() string {
	return string(t)
}

func main() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
	}

	// Same as if - else
	t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
	}
	
	// Type switch - works with interfaces ONLY!
	whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
	whatAmI("hey")

	// Fallthrough - should be the last thing in the case
	v := 42
	switch v {
	case 100:
		fmt.Println(100)
		fallthrough
	case 42:
		fmt.Println(42)
		fallthrough
	case 1:
		fmt.Println(1)
		fallthrough
	default:
		fmt.Println("default")
	}

	// Type switch can also check if other interface is implemented!
	var testBytes interface{}
	testBytes = Test("Hello!")
	switch testBytes.(type) {
	case []byte:
		fmt.Println("It's a byte slice!")
	case string:
		fmt.Println("It's string!")
	case fmt.Stringer:
		fmt.Println(testBytes)
	}

	// This is the same as following:
	if _, ok := testBytes.([]byte); ok {
		fmt.Println("It's a byte slice!")
	} else if _, ok := testBytes.(string); ok {
		fmt.Println("It's string!")
	} else if _, ok := testBytes.(fmt.Stringer); ok {
		fmt.Println(testBytes)
	}

	fmt.Printf("%T\n", testBytes.(Test))
	fmt.Printf("%T\n", testBytes.(fmt.Stringer))
}
