package main

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func choosePivotFirst(arr []int) int {
	return 0
}

func choosePivotLast(arr []int) int {
	return len(arr) - 1
}

type entry struct {
	k int
	v int
}

func choosePivotMedian(arr []int) int {
	m := len(arr) / 2 // middle value

	if (len(arr) % 2) == 0 {
		// even
		m--
	}

	res := []entry{
		{k: 0, v: arr[0]},
		{k: m, v: arr[m]},
		{k: len(arr) - 1, v: arr[len(arr)-1]},
	}

	for i := 0; i < 3; i++ {
		for j := i; j > 0 && res[j-1].v > res[j].v; j-- {
			res[j], res[j-1] = res[j-1], res[j]
		}
	}

	return res[1].k
}

func partition(arr []int) int {
	p := arr[0]
	i := 1

	for j := range arr {
		if arr[j] < p {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[0], arr[i-1] = arr[i-1], arr[0]

	return i - 1
}

func quickSort(arr []int, choosePivot func(arr []int) int) int {
	// Base case.
	if len(arr) <= 1 {
		return 0
	}

	i := choosePivot(arr)

	// Swap pivot with first element
	arr[0], arr[i] = arr[i], arr[0]

	comps := len(arr) - 1
	j := partition(arr) // j - new pivot

	lComps := quickSort(arr[0:j], choosePivot)
	rComps := quickSort(arr[j+1:], choosePivot)

	return comps + lComps + rComps
}

func main() {
	// Prepare input data
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/quick_sort/QuickSort.txt")
	inputStrings := strings.Split(string(inputRaw), "\r\n")

	input := make([]int, len(inputStrings))

	for i, v := range inputStrings {
		input[i], _ = strconv.Atoi(v)
	}

	// Prepare result for comparison
	inputWant := make([]int, len(input))
	copy(inputWant, input)
	sort.Ints(inputWant)

	// Run sort
	comps := quickSort(input, choosePivotMedian)

	// Output comparisons total
	fmt.Println(comps)

	// Check sorting is correct
	fmt.Println(reflect.DeepEqual(input, inputWant))
}
