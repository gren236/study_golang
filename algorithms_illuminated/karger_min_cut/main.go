package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createTestGraph() *graph {
	test := newGraph()
	test.addVertex(1, nil)
	test.addVertex(2, nil)
	test.addVertex(3, nil)
	test.addVertex(4, nil)

	test.vertices[1].adj = append(test.vertices[1].adj, test.vertices[2], test.vertices[3])
	test.vertices[2].adj = append(test.vertices[2].adj, test.vertices[1], test.vertices[3], test.vertices[4])
	test.vertices[3].adj = append(test.vertices[3].adj, test.vertices[1], test.vertices[2], test.vertices[4])
	test.vertices[4].adj = append(test.vertices[4].adj, test.vertices[2], test.vertices[3])

	return test
}

func parseAdjListFile() [][]int {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/karger_min_cut/kargerMinCut.txt")
	inputStrings := strings.Split(string(inputRaw), "\r\n")
	inputStrings = inputStrings[:len(inputStrings)-1]

	res := make([][]int, len(inputStrings))
	for i, s := range inputStrings {
		inputRow := strings.Split(s, "\t")
		inputRow = inputRow[:len(inputRow)-1]

		res[i] = make([]int, len(inputRow))
		for j, v := range inputRow {
			vInt, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			res[i][j] = vInt
		}
	}

	return res
}

func buildNewGraphFromAdjList(lst [][]int) *graph {
	res := newGraph()

	for _, v := range lst {
		res.addVertex(v[0], nil)
	}

	for _, v := range lst {
		vLbl := v[0]

		for _, vAdj := range v[1:] {
			res.vertices[vLbl].adj = append(res.vertices[vLbl].adj, res.vertices[vAdj])
		}
	}

	return res
}

func main() {
	adjLst := parseAdjListFile()
	testF := buildNewGraphFromAdjList(adjLst)

	testF.contract()

	minCutEdgesTotal, err := testF.assertContractionSuccess()
	if err != nil {
		panic(err)
	}

	//trials := 211932
	trials := 40000

	for i := 1; i <= trials; i++ {
		test := buildNewGraphFromAdjList(adjLst)
		test.contract()

		minCutEdges, err := test.assertContractionSuccess()
		if err != nil {
			panic(err)
		}

		if minCutEdges < minCutEdgesTotal {
			minCutEdgesTotal = minCutEdges
		}

		if i%1000 == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("!!!!!!", minCutEdgesTotal)
}
