package leetcode

// Given an integer array nums, move all 0's to the end of it while maintaining the relative order of the non-zero elements.
func moveZeroes(nums []int) {
	var j int
	for i := 1; i < len(nums); i++ {
		if nums[j] == 0 && nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
		}

		if nums[j] != 0 {
			j++
		}
	}
}
