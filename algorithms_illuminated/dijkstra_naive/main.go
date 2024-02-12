package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseAdjListFileAndBuildGraph() *graph {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/dijkstra_naive/dijkstraData.txt")
	inputStrings := strings.Split(string(inputRaw), "\r\n")
	inputStrings = inputStrings[:len(inputStrings)-1]

	res := newGraph()
	for _, s := range inputStrings {
		inputRow := strings.Split(s, "\t")
		inputRow = inputRow[:len(inputRow)-1]

		uLabel, _ := strconv.Atoi(inputRow[0])
		for _, tplRaw := range inputRow[1:] {
			tpl := strings.Split(tplRaw, ",")
			vLabel, _ := strconv.Atoi(tpl[0])
			w, _ := strconv.Atoi(tpl[1])

			res.addEdge(uLabel, vLabel, w)
		}
	}

	return res
}

func main() {
	g := parseAdjListFileAndBuildGraph()

	g.setMinDistances(g.vertices[1])

	fmt.Printf(
		"%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
		g.vertices[7].minPath,
		g.vertices[37].minPath,
		g.vertices[59].minPath,
		g.vertices[82].minPath,
		g.vertices[99].minPath,
		g.vertices[115].minPath,
		g.vertices[133].minPath,
		g.vertices[165].minPath,
		g.vertices[188].minPath,
		g.vertices[197].minPath,
	)
}
