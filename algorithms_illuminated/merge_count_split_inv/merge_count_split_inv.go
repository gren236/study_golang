package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func merge(a []int, b []int) ([]int, int) {
	var i, j, sInv int
	res := make([]int, len(a)+len(b))

	for k := 0; k < len(res); k++ {
		if i >= len(a) {
			res[k] = b[j]
			j++
			continue
		}
		if j >= len(b) {
			res[k] = a[i]
			i++
			continue
		}

		if a[i] < b[j] {
			res[k] = a[i]
			i++
		} else {
			res[k] = b[j]
			j++
			sInv += len(a) - i
		}
	}

	return res, sInv
}

func mergeSort(arr []int) ([]int, int) {
	ln := len(arr)

	// Basic case
	if ln <= 1 {
		return arr, 0
	}

	// Recursive case
	a, lInv := mergeSort(arr[:ln/2])
	b, rInv := mergeSort(arr[ln/2:])
	c, sInv := merge(a, b)

	return c, lInv + rInv + sInv
}

func main() {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/merge_count_split_inv/IntegerArray.txt")
	inputStrings := strings.Split(string(inputRaw), "\r\n")
	inputStrings = inputStrings[:len(inputStrings)-1]

	input := make([]int, len(inputStrings))

	for i, v := range inputStrings {
		input[i], _ = strconv.Atoi(v)
	}

	_, inv := mergeSort(input)
	fmt.Println(inv)
}
