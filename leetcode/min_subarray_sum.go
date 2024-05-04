package leetcode

// Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray whose
// sum is greater than or equal to target. If there is no such subarray, return 0 instead.
func minSubArrayLen(target int, nums []int) int {
	var bestSubSize, subSize, sum, j int

	for i := 0; i < len(nums); i++ {
		subSize++
		sum += nums[i]

		if sum >= target {
			for (sum - nums[j]) >= target {
				sum -= nums[j]
				j++
				subSize--
			}

			if bestSubSize == 0 {
				bestSubSize = subSize
			} else {
				bestSubSize = min(bestSubSize, subSize)
			}
		}
	}

	return bestSubSize
}
