package leetcode

import "slices"

// Given an integer array nums of 2n integers, group these integers into n pairs (a1, b1), (a2, b2), ..., (an, bn) such
// that the sum of min(ai, bi) for all i is maximized. Return the maximized sum.
func arrayPairSum(nums []int) int {
	slices.Sort(nums)

	var res int
	for i := 0; i < len(nums); i += 2 {
		res += min(nums[i], nums[i+1])
	}

	return res
}
