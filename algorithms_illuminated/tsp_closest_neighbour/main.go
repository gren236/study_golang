package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseTestFile() *graph {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/tsp_closest_neighbour/nn.txt")
	inputStrings := strings.Split(string(inputRaw), "\r\n")

	size, _ := strconv.Atoi(inputStrings[0])

	res := newGraph(size)
	for i, inputString := range inputStrings[1 : len(inputStrings)-1] {
		row := strings.Split(inputString, " ")
		lbl, _ := strconv.Atoi(row[0])
		x, _ := strconv.ParseFloat(row[1], 64)
		y, _ := strconv.ParseFloat(row[2], 64)

		res.vertices[i] = &vertex{
			lbl: lbl,
			x:   x,
			y:   y,
		}
	}

	return res
}

// getTSPNeighbour receives an undirected graph and returns a minimal value tour
func getTSPNeighbour(g *graph) float64 {
	g.resetUnvisited()

	s := g.vertices[0]
	delete(g.unvisited, s.lbl)

	var minTourCost float64

	for len(g.unvisited) > 0 {
		if len(g.unvisited)%1000 == 0 {
			fmt.Printf("unvisited size: %d, min tour cost: %v\n", len(g.unvisited), minTourCost)
		}

		t, cost := g.findClosestUnvisited(s)

		minTourCost += cost

		s = t
		delete(g.unvisited, s.lbl)
	}

	minTourCost += s.distanceTo(g.vertices[0])

	return minTourCost
}

func main() {
	g := parseTestFile()

	res := getTSPNeighbour(g)

	fmt.Println(int(res))
}
