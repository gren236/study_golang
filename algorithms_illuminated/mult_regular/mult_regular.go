package main

import (
	"fmt"
	"math"
)

func reverse(a []int) {
	for i := len(a)/2-1; i >= 0; i-- {
		opp := len(a)-1-i
		a[i], a[opp] = a[opp], a[i]
	}
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func intMult(x, y, n int) int {
	resultRaw := make([][]int, n)

	var loopY int
	// Iterate over numbers to generate array of sum elements
	for i := 0; x > 0; i++ {
		var carry, currentX int
		currentX, x = x % 10, x / 10

		loopY = y
		// Iterate over all numbers of y for specific number of x
		for loopY > 0 {
			var current int
			current, loopY = loopY % 10, loopY / 10
			res := current * currentX
			res += carry

			if res > 9 {
				carry, res = res / 10, res % 10
			} else {
				carry = 0
			}
			resultRaw[i] = append(resultRaw[i], res)
		}
		if carry > 0 {
			resultRaw[i] = append(resultRaw[i], carry)
			carry = 0
		}
		reverse(resultRaw[i])
	}

	// Sum
	var result int
	for k, v := range resultRaw {
		num := sliceToInt(v)
		result += num * int(math.Pow10(k))
	}

	return result
}

func main() {
	fmt.Println(intMult(1234, 5678, 4))
}