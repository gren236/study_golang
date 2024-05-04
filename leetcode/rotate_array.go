package leetcode

// Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.
func rotate(nums []int, k int) {
	for {
		k = k % len(nums)
		if k == 0 {
			return
		}

		var i int
		j := len(nums) - k
		for j < len(nums) {
			if i == len(nums)-k {
				k = k - i
				break
			}

			nums[i], nums[j] = nums[j], nums[i]
			j++
			i++
		}

		nums = nums[i:]
	}
}
