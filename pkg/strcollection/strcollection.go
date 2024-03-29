package strcollection

// Index Returns the first index of target string or -1 if no match found
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}

	return -1
}

// Include return true if the target string t is in slice
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns true if one of the strings in the slice satisfies the predicate f
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}

	return false
}

// All returns true if all the strings in the slice satisfy the predicate f
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}

	return true
}

// Filter returns a new slice containing all strings in the slice that satisfy the predicate f
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)

	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}

	return vsf
}

// Map returns a new slice containing the result of applying f to each string
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))

	for i, v := range vs {
		vsm[i] = f(v)
	}

	return vsm
}
