package main

import (
	"fmt"
	"github.com/gren236/study_golang/algorithms_illuminated/kosaraju_scc"
	"os"
	"strconv"
	"strings"
)

func parseAdjListAndBuildGraph() *kosaraju_scc.Graph {
	//inputRaw, _ := os.ReadFile("./algorithms_illuminated/kosaraju_scc/SCC.txt")
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/kosaraju_scc/input_mostlyCycles_1_8.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")

	g := kosaraju_scc.NewGraph(0)

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

		g.AddEdge(u, v)
	}

	return g
}

func main() {
	g := parseAdjListAndBuildGraph()

	g.MarkSCCs()

	fmt.Println(g.GetSccSizeRating())
}
