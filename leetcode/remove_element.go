package leetcode

// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place. The order of
// the elements may be changed. Then return the number of elements in nums which are not equal to val.
func removeElement(nums []int, val int) int {
	var k int
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}

	return k
}
