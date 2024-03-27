package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type key struct {
	i, c int
}

type item struct {
	v, s int
}

func parseTestFile() ([]item, int) {
	//inputRaw, _ := os.ReadFile("./algorithms_illuminated/knapsack_big/input_random_12_100_10.txt") // 133
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/knapsack_big/knapsack_big.txt") // 133
	inputStrings := strings.Split(string(inputRaw), "\n")

	header := strings.Split(inputStrings[0], " ")
	capacity, _ := strconv.Atoi(header[0])
	size, _ := strconv.Atoi(header[1])

	res := make([]item, size)
	for i, inputString := range inputStrings[1 : len(inputStrings)-1] {
		row := strings.Split(inputString, " ")
		v, _ := strconv.Atoi(row[0])
		s, _ := strconv.Atoi(row[1])

		res[i] = item{v, s}
	}

	return res, capacity
}

func ksCalc(cache map[key]int, items []item, i, capacity int) int {
	if i >= len(items) {
		return 0
	}

	// Check if such case is cached already
	if cVal, ok := cache[key{i, capacity}]; ok {
		return cVal
	}

	it := items[i]

	if it.s > capacity {
		return ksCalc(cache, items, i+1, capacity)
	}

	res := max(ksCalc(cache, items, i+1, capacity), ksCalc(cache, items, i+1, capacity-it.s)+it.v)

	// Cache result
	cache[key{i, capacity}] = res

	return res
}

func main() {
	items, capacity := parseTestFile()

	cache := make(map[key]int)

	res := ksCalc(cache, items, 0, capacity)

	fmt.Println(res)
}
