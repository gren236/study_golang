package main

import (
	"fmt"
	"golang.org/x/text/unicode/norm"
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

	// funnyCounter()

	// Strings are immutable!
	// But string can be created from slice of bytes
	var1 := []uint8{72, 101, 108, 108, 111}
	var2 := string(var1)
	fmt.Println(var2)

	// Also can be made with slice of runes!
	varRunes := []rune{0x00C8, 0x00FD, 0x2318, 0x00FF}
	varString := string(varRunes)
	fmt.Println(varString)

	// String can be defined with backticks, but it's gonna be raw. \n won't work, but you can still
	// get to the new line if you make it so in source code!
	bs := `Hello!\nSomebody
		Once told me...`
	fmt.Println(bs)

	// Comparisons - lexicographical order
	fmt.Println('a' < 'b')
	fmt.Println('a' < 145)
	fmt.Println("hellobac" < "hellobb")
	fmt.Println("J" < "j")
	fmt.Println("привет1" < "привет2")
	//fmt.Println("HI" == []rune{0x0048, 0x0049})

	fmt.Println([]rune("Hello!⌘"))
	fmt.Println([]byte("Hello!⌘"))

	bytes := []byte(string(0x00EA))
	fmt.Println(bytes)
	fmt.Printf("%s %s\n", []byte{bytes[0]}, []byte{bytes[1]})

	// These won't work without normalization!
	fmt.Println(string([]rune{'e', 0x0301}))
	estr := "e\u0301"
	fmt.Println(estr)
	// We have to use string normalization to make it work
	fmt.Println(norm.NFC.String(estr))
}

func funnyCounter() {
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
}