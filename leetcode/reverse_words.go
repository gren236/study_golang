package leetcode

import (
	"slices"
	"strings"
)

// Given an input string s, reverse the order of the words.
func reverseWords(s string) string {
	res := strings.Fields(s)
	slices.Reverse(res)

	return strings.Join(res, " ")
}
