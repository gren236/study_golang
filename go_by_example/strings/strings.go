package main

import (
	"fmt"
	"time"
)

var testString = "Hello!⌘"

func main() {
	// len() treats string as a slice of bytes, so it returns a number of bytes that form a string, not code points
	fmt.Println(len(testString))

	for i := 0; i < len(testString); i++ {
		fmt.Printf("% x", testString[i])
	}
	fmt.Println()

	// RUne is just a UTF-8 code point value in decimal (can be assigned as hex)
	var testRune rune
	testRune = 0x2629
	fmt.Printf("%T - %#U\n", testRune, testRune)

	// Funny utf-8 counter
	ticker := time.NewTicker(5 * time.Millisecond)
	var r rune = 0x0061
	for {
		select {
		case <-ticker.C:
			fmt.Printf("\r%#U", r)
			r++
		}
	}

	// WIP
}