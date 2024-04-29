package leetcode

const (
	one  = uint8(49)
	zero = uint8(48)
)

// Given two binary strings a and b, return their sum as a binary string.
func addBinary(a string, b string) string {
	var res string
	var carry bool

	for i := 0; i < max(len(a), len(b)) || carry; i++ {
		aCurr := zero
		if i+1 <= len(a) {
			aCurr = a[len(a)-1-i]
		}

		bCurr := zero
		if i+1 <= len(b) {
			bCurr = b[len(b)-1-i]
		}

		// 1+0 OR 0+1
		if aCurr != bCurr {
			if carry {
				res = "0" + res
			} else {
				res = "1" + res
			}
		}

		// 1+1
		if aCurr == one && bCurr == one {
			if carry {
				res = "1" + res
			} else {
				res = "0" + res
				carry = true
			}
		}

		// 0+0
		if aCurr == zero && bCurr == zero {
			if carry {
				res = "1" + res
				carry = false
			} else {
				res = "0" + res
			}
		}
	}

	return res
}
