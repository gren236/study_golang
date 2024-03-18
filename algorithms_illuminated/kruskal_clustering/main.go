package main

import (
	"fmt"
	"github.com/gren236/study_golang/pkg/ufind"
	"os"
	"sort"
	"strconv"
	"strings"
)

func parseAdjListFile() *graph {
	res := newGraph()

	inputRaw, _ := os.ReadFile("./algorithms_illuminated/kruskal_clustering/clustering1.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")

	nodes, _ := strconv.Atoi(inputStrings[0])
	res.fillVerticesSequentially(nodes)

	inputStrings = inputStrings[1 : len(inputStrings)-1]
	for _, s := range inputStrings {
		inputRow := strings.Split(s, " ")

		v, _ := strconv.Atoi(inputRow[0])
		w, _ := strconv.Atoi(inputRow[1])
		cost, _ := strconv.Atoi(inputRow[2])

		res.addEdge(edge{
			id:   inputRow[0] + "_" + inputRow[1],
			cost: cost,
			v:    vertex(v),
			w:    vertex(w),
		})
	}

	return res
}

func initUFind(vertices []vertex) ufind.Container[vertex] {
	res, err := ufind.New(vertices)
	if err != nil {
		panic(err)
	}

	return res
}

func main() {
	// Init
	g := parseAdjListFile()
	u := initUFind(g.vertices)
	targetK := 4
	k := len(g.vertices)
	currentSpacing := -1

	// Sort edges
	sort.Sort(byCost(g.edges))

	// Main loop
	for _, e := range g.edges {
		if targetK-1 == k {
			break
		}

		vGroup, ok := u.Find(e.v)
		if !ok {
			panic("wrong v to find")
		}

		wGroup, ok := u.Find(e.w)
		if !ok {
			panic("wrong w to find")
		}

		if vGroup != wGroup {
			currentSpacing = e.cost

			u.Union(e.v, e.w)

			k--
		}
	}

	fmt.Println(currentSpacing)
}
