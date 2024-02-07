package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseAdjListAndBuildGraph() *graph {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/graph_kosaraju_scc/SCC.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")

	g := newGraph()

	for _, s := range inputStrings {
		inputRow := strings.Split(s, " ")

		u, err := strconv.Atoi(inputRow[0])
		if err != nil {
			panic("could not represent string u as int")
		}

		v, err := strconv.Atoi(inputRow[1])
		if err != nil {
			panic("could not represent string v as int")
		}

		g.addEdge(u, v)
	}

	return g
}

func main() {
	g := parseAdjListAndBuildGraph()

	g.markSCCs()

	fmt.Println(g.getSccSizeRating()[:5])
}
