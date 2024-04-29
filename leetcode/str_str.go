package leetcode

// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if
// needle is not part of haystack.
func strStr(haystack string, needle string) int {
	entries := make(map[int]int) // start_index->current_index_j

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			entries[i] = 0
		}

		for s, j := range entries {
			if haystack[i] == needle[j] {
				if j == len(needle)-1 {
					return s
				}

				entries[s]++
			} else {
				delete(entries, s)
			}
		}
	}

	return -1
}
