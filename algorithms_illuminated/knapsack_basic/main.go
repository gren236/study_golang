package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type item struct {
	v, s int
}

func parseTestFile() ([]item, int) {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/knapsack_basic/knapsack1.txt")
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

func main() {
	items, c := parseTestFile()

	a := make([][]int, len(items)+1)
	for i := range a {
		a[i] = make([]int, c+1)
	}

	//for i := range c {
	//	a[0][i] = 0
	//}

	for i := 1; i <= len(items); i++ {
		for j := 0; j <= c; j++ {
			s_i := items[i-1].s
			v_i := items[i-1].v

			if s_i > j {
				a[i][j] = a[i-1][j]
			} else {
				a[i][j] = max(a[i-1][j], a[i-1][j-s_i]+v_i)
			}
		}
	}

	fmt.Println(a[len(a)-1][c])
}
