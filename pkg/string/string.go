/*
This is a simple package for string operations
There is only one operation now, so... we will work on that :)
Functions can be called like that:

	string.Reverse("Hello!")

Simple as that.
*/
package string

// Reverse takes a string and returns this string with its characters in reversed order.
func Reverse(s string) string {
	b := []rune(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
