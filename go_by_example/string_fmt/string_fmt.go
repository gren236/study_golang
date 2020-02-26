package main

import (
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

type Foo uint64

type Bar uint64

func (b Bar) String() string {
	return fmt.Sprintf("the value of Bar is: %d\n", uint64(b))
}

func main() {
	p := Point{1, 2}

	// Print struct without any formatting or newline
	fmt.Print(p); fmt.Print("\n")

	// Put S before any function to get formatted string, not write it to stdout
	pstr := fmt.Sprint(p)
	fmt.Println(pstr)

	// Print instance of Point struct
	fmt.Printf("%v\n", p)

	// Add + to also output struct field names
	fmt.Printf("%+v\n", p)
	fmt.Printf("%+v\n", &p)

	// Use # instead of + to get value as a Go source code snippet
	fmt.Printf("%#v\n", p)

	// Use %T to print a type
	fmt.Printf("%T\n", p)

	// Lowercase t is for booleans
	fmt.Printf("%t\n", true)

	// There are lots of decimal representations. Use d for standard base-10
	fmt.Printf("%d\n", 123)
	// Print a binary representation
	fmt.Printf("%b\n", 14)
	// Character, corresponding to a given integer
	fmt.Printf("%c\n", 33)
	// Unicode format (use # to output original character next to code)
	fmt.Printf("%#U\n", '⌘', '⌘')
	// Hex encoding
	fmt.Printf("%x\n", 456)
	// Add # to show as in code, with leading "0x"
	fmt.Printf("%#x\n", 456)

	// Several format options for floats. f is for basic output
	fmt.Printf("%f\n", 78.9)
	// e and E - slightly different scientific notations
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	// %g or %G can be used to output float as-is, but use scientific notation if it's too large
	fmt.Printf("%G\n", 123400000.023427384928374823)

	// Basic string printing
	fmt.Printf("%s\n", "\"string\"")
	// Use q to double-quote strings as in source code
	fmt.Printf("%q\n", "\"string\"")
	// x also can be used to output string as hex bytes
	fmt.Printf("%x\n", "hex this")

	// Use %p to print a pointer
	fmt.Printf("%p\n", &p)

	// Control equal width of outputting with numbers after % sign
	fmt.Printf("|%6d|%6d|\n", 12, 345)
	// Control width with decimal precision for floats
	fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// Use - before width number to make it left-justify
	fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// width can be applied to strings to - use it for nice looking tables!
	fmt.Printf("|%6s|%6s|\n", "foo", "b")
	// - for left-justify, as with numbers
	fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

	// Format string without printing it anywhere
	s := fmt.Sprintf("a %s\n", "string")
	fmt.Println(s)

	// You can choose where to print output with Fprintf
	fmt.Fprintf(os.Stderr, "as %s\n", "error")

	// To use % character raw, you can escape it with another preceding %
	fmt.Printf("%d%%\n", 15)

	// Implement String method for custom type to use it with custom value for default %v verb
	var ftype Foo = 345
	fmt.Printf("%v\n", ftype)
	var btype Bar = 42
	fmt.Printf("%v\n", btype)
}