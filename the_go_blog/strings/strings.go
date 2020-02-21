package main

import (
	"fmt"
	"unicode/utf8"
)

// Arbitrary set of chars
const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

func main() {
	fmt.Println(sample)

	// Let's try analysing bytes individually
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}

	fmt.Println()

	// Shorter way of outputting bytes slice
	fmt.Printf("%x\n", sample)
	// Also, we can use "space" flag to divide bytes
	fmt.Printf("% x\n", sample)

	// Also, we can escape any non-printable byte in output
	fmt.Printf("%q\n", sample)
	// Escape all signs
	fmt.Printf("%+q\n", sample)

	// Change string to slice of bytes
	sampleBytes := []byte(sample)
	// Also, we can escape any non-printable byte in output
	fmt.Printf("%q\n", sampleBytes)
	// Escape all signs
	fmt.Printf("%+q\n", sampleBytes)
	// Output is the same!

	// Loop over the string using the %q format on each byte
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q", sample[i])
	}

	fmt.Println()
	stringLiterals()

	fmt.Println()
	runes()

	fmt.Println()
	forRange()
}

func forRange() {
	// There is difference between for and for range over strings
	// for iterates over every byte in string. For range - converts characters to runes (UTF code points)
	const nihongo = "日本語\x8c\xe2\x98"
	for index, runeValue := range nihongo {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}

	fmt.Println()

	// We can do the same with unicode/utf8 package
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func runes() {
	// Runes are the same as code-points in UTF-8
	sampleRune := '⌘' // Same as 0x2318
	fmt.Println(sampleRune) // Output 8984 - decimal of 0x2318!
	// So, we can use any bytes to represent strings (which are just byte slices), even for UTF-8 notation in hex
	// But also there is rune that stores UTF-8 "code point" as decimal.
}

func stringLiterals() {
	// Backtick operator - for raw string, without escape sequences
	const placeOfInterest = `⌘`

	// Just a plain string output
	fmt.Printf("Plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	// ASCII-only quoted string output
	fmt.Printf("Quoted string: ")
	fmt.Printf("%+q", placeOfInterest)
	fmt.Printf("\n")

	// Individual bytes in HEX
	fmt.Printf("Hex bytes: ")
	for i := 0; i < len(placeOfInterest); i++ {
		fmt.Printf("%x ", placeOfInterest[i])
	}
	fmt.Printf("\n")

	// Concluding - string literals always contain UTF-8 chars, but other strings may contain everything
}