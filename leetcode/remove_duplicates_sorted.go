package leetcode

// Given an integer array nums sorted in non-decreasing order, remove the duplicates in-place such that each unique
// element appears only once. The relative order of the elements should be kept the same. Then return the number of
// unique elements in nums.
func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	var j int
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[j] {
			j++
			nums[j] = nums[i]
		}
	}

	return j + 1
}
