package main

import "fmt"

func merge(a []int, b []int) []int {
	var i, j int
	res := make([]int, len(a) + len(b))

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
		}
	}

	return res
}

func mergeSort(arr []int) []int {
	ln := len(arr)

	// Basic case
	if ln <= 1 {
		return arr
	}

	// Recursive case
	a := mergeSort(arr[:ln/2])
	b := mergeSort(arr[ln/2:])

	return merge(a, b)
}

func main() {
	test := []int{5, 3, 8, 9, 1, 7, 0, 2, 6, 4, 11, 10}

	fmt.Println(mergeSort(test))
}
