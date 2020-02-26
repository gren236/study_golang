package main

import (
	"bytes"
	"fmt"
	"regexp"
)

// Regexp global variable
var REGEXP_PEACH = regexp.MustCompile("p([a-z]+)ch")

func main() {
	// Regular regexp matching
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// For other regexp tasks we should compile on optimized regexp struct
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Printf("%#v\n", r)

	// Regular match with compiled regexp
	fmt.Println(r.MatchString("peach"))

	// find the match in a string
	fmt.Println(r.FindString("peach punch"))

	// Return information for whole match and submatches
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// Same, but returns indexes
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// Also, we can add All to function to find all match entries in string. Second number is matches limit.
	fmt.Println(r.FindAllString("peach oh no, punch pinch", -1))

	// Available for all above functions
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Limit the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// If we remove String from function name, we should be able to pass bytes slice (what string actually is)
	fmt.Println(r.Match([]byte("peach")))
	fmt.Println(r.Find([]byte("peach punch"))) // In that case, find returns results as byte slice as well

	// For global variables use MustCompile, because it returns 1 value
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(REGEXP_PEACH.FindString("peach punch"))

	// We can replace substrings
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// The Func variant allows you to transform matched text with a given function.
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}