package leetcode

import (
	"slices"
	"strings"
)

// Given a string s, reverse the order of characters in each word within a sentence while still preserving whitespace
// and initial word order.
func reverseLetters(s string) string {
	res := strings.Fields(s)

	for i, str := range res {
		rstr := []rune(str)

		slices.Reverse(rstr)

		res[i] = string(rstr)
	}

	return strings.Join(res, " ")
}
