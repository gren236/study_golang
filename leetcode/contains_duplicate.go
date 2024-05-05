package leetcode

// Given an integer array nums, return true if any value appears at least twice in the array, and return false if every
// element is distinct.
func containsDuplicate(nums []int) bool {
	cnt := make(map[int]bool, len(nums))

	for _, n := range nums {
		if cnt[n] {
			return true
		}

		cnt[n] = true
	}

	return false
}
