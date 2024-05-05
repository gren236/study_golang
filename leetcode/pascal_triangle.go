package leetcode

import "math/big"

// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}

	res := make([]int, rowIndex+1)
	res[0], res[rowIndex] = 1, 1

	for i := 1; i < rowIndex; i++ {
		res[i] = int(big.NewInt(0).Binomial(int64(rowIndex), int64(i)).Int64())
	}

	return res
}
