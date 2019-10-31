package main

import "fmt"

func main() {
	if 7%2 == 0 {
        fmt.Println("7 is even")
    } else {
        fmt.Println("7 is odd")
	}
	
	if 8%4 == 0 {
        fmt.Println("8 is divisible by 4")
	}
	
	// num is in the scope of this if ONLY
	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
	}
	
	// It work for every var declared inside if
	if 8%4 == 0 {
		var foo string = "I'm inside if!"
		if true {
			fmt.Println(foo)
		}
        fmt.Println("8 is divisible by 4")
	}
}