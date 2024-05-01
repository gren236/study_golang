package leetcode

// Given a binary array nums, return the maximum number of consecutive 1's in the array.
func findMaxConsecutiveOnes(nums []int) int {
	var maxResult, k int

	for _, num := range nums {
		if num != 1 {
			maxResult = max(maxResult, k)
			k = 0
			continue
		}

		k++
	}

	return max(maxResult, k)
}
