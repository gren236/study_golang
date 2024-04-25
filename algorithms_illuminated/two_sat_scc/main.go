package main

import (
	"fmt"
	"github.com/gren236/study_golang/algorithms_illuminated/kosaraju_scc"
	"os"
	"strconv"
	"strings"
)

func parseAdjListAndBuildGraph(path string) *kosaraju_scc.Graph {
	inputRaw, _ := os.ReadFile(path) // 0
	inputStrings := strings.Split(string(inputRaw), "\n")

	size, _ := strconv.Atoi(inputStrings[0])
	g := kosaraju_scc.NewGraph(2 * size)

	for _, s := range inputStrings[1 : len(inputStrings)-1] {
		inputRow := strings.Split(s, " ")

		u, err := strconv.Atoi(inputRow[0])
		if err != nil {
			panic("could not represent string u as int")
		}

		v, err := strconv.Atoi(inputRow[1])
		if err != nil {
			panic("could not represent string v as int")
		}

		g.AddEdge(-u, v)
		g.AddEdge(-v, u)
	}

	return g
}

// twoSatSCC looks for SCCs in the implication graph, and reports if given 2SAT instance is satisfiable
func twoSatSCC(g *kosaraju_scc.Graph) bool {
	g.MarkSCCs() // Find SCC

	for _, u := range g.Vertices {
		if v, ok := g.Vertices[-u.Label]; ok {
			if u.SCC == v.SCC {
				return false
			}
		}
	}

	return true
}

func main() {
	for i := range 6 {
		path := fmt.Sprintf("./algorithms_illuminated/two_sat_scc/2sat%d.txt", i+1)

		g := parseAdjListAndBuildGraph(path)
		res := twoSatSCC(g)

		if res {
			fmt.Print("1")
		} else {
			fmt.Print("0")
		}
	}
}
