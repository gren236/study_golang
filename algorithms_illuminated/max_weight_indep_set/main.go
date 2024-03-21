package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vertex struct {
	id, weight int
}

func parseListFile() []vertex {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/max_weight_indep_set/mwis.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")

	length, _ := strconv.Atoi(inputStrings[0])
	res := make([]vertex, length)

	inputStrings = inputStrings[1 : len(inputStrings)-1]
	for i, s := range inputStrings {
		inputRow := strings.Split(s, " ")

		w, _ := strconv.Atoi(inputRow[0])

		res[i] = vertex{
			id:     i + 1,
			weight: w,
		}
	}

	return res
}

func getMaxWeightIndependentSet(in []vertex) []int {
	res := make([]int, len(in)+1)
	res[0] = 0
	res[1] = in[0].weight

	for i := 2; i <= len(in); i++ {
		res[i] = max(res[i-1], res[i-2]+in[i-1].weight)
	}

	return res
}

func reconstructVertices(a []int, vList []vertex) map[int]int {
	res := make(map[int]int) // id -> weight

	i := len(vList)

	for i >= 2 {
		v := vList[i-1]

		if a[i-1] >= a[i-2]+v.weight {
			i--
		} else {
			res[v.id] = v.weight
			i -= 2
		}
	}
	if i == 1 {
		res[vList[0].id] = vList[0].weight
	}

	return res
}

func main() {
	vList := parseListFile()
	//vList := []vertex{
	//	{1, 3},
	//	{2, 2},
	//	{3, 1},
	//	{4, 6},
	//	{5, 4},
	//	{6, 5},
	//}

	sols := getMaxWeightIndependentSet(vList)

	vIncl := reconstructVertices(sols, vList)

	testVertices := []int{1, 2, 3, 4, 17, 117, 517, 997}
	resString := ""

	for _, testVertex := range testVertices {
		if _, ok := vIncl[testVertex]; ok {
			resString += "1"
		} else {
			resString += "0"
		}
	}

	fmt.Println(resString)
}
