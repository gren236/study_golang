package leetcode

// Write a function to find the longest common prefix string amongst an array of strings.
// If there is no common prefix, return an empty string "".
func longestCommonPrefix(strs []string) string {
	var maxSize int
	var res string

	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}

		if len(str) > maxSize {
			maxSize = len(str)
		}
	}

	for i := 0; i < maxSize; i++ {
		if i+1 > len(strs[0]) {
			return res
		}

		sample := strs[0][i]

		for _, str := range strs[1:] {
			if (i+1 > len(str)) || (str[i] != sample) {
				return res
			}
		}

		res += string(sample)
	}

	return res
}
