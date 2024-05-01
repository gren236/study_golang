package leetcode

import "sort"

// Given a 1-indexed array of integers numbers that is already sorted in non-decreasing order, find two numbers such
// that they add up to a specific target number. Let these two numbers be numbers[index1] and numbers[index2] where
// 1 <= index1 < index2 <= numbers.length.
//
// Return the indices of the two numbers, index1 and index2, added by one as an integer array [index1, index2] of length 2.
//
// Must use only constant extra space.
func twoSum(numbers []int, target int) []int {
	var i int
	for len(numbers) > 1 {
		currTarget := target - numbers[0]

		x := sort.SearchInts(numbers[1:], currTarget)
		if x+1 < len(numbers) && numbers[x+1] == currTarget {
			return []int{i + 1, x + i + 2}
		}

		numbers = numbers[1:]
		i++
	}

	return nil
}
