package main

import (
	"fmt"
	"github.com/gren236/study_golang/pkg/heap"
	"math"
	"os"
	"strconv"
	"strings"
)

type heapVertex struct {
	v          *vertex
	wFrom, wTo *vertex
	cost       int
}

func (h heapVertex) Less(t heapVertex) bool {
	return h.cost < t.cost
}

func parseAdjListFileAndBuildGraph() *graph {
	inputRaw, _ := os.ReadFile("./algorithms_illuminated/prim_mst/edges.txt")
	inputStrings := strings.Split(string(inputRaw), "\n")
	inputStrings = inputStrings[1 : len(inputStrings)-1]

	res := newGraph()
	for _, s := range inputStrings {
		inputRow := strings.Split(s, " ")

		uLabel, _ := strconv.Atoi(inputRow[0])
		vLabel, _ := strconv.Atoi(inputRow[1])
		cost, _ := strconv.Atoi(inputRow[2])

		res.addEdge(uLabel, vLabel, cost)
	}

	return res
}

func primMST(g *graph, s *vertex) int {
	var mstCostSum int
	heapVertexIndices := make(map[int]int)

	hp := heap.NewHeap[heapVertex](func(h *heapVertex, i int) {
		heapVertexIndices[h.v.label] = i
	})

	g.resetExplored()
	g.explored[s.label] = s

	// Initialize heap
	for _, v := range g.vertices {
		if v == s {
			continue
		}

		currHeapVertex := heapVertex{v: v, cost: math.MaxInt}

		if cost, connected := s.checkConnection(v); connected {
			currHeapVertex.wFrom = s
			currHeapVertex.wTo = v
			currHeapVertex.cost = cost
		}

		hp.Insert(currHeapVertex)
	}

	// Main loop, go through all vertices in heap
	for hp.Len() > 0 {
		w := hp.ExtractMin()

		g.explored[w.v.label] = w.v

		mstCostSum += w.cost

		for _, e := range w.v.edges {
			if _, ok := g.explored[e.adj.label]; ok {
				continue
			}

			currHeapVertex := hp.GetByIndex(heapVertexIndices[e.adj.label])

			if e.cost < currHeapVertex.cost {
				hp.Delete(heapVertexIndices[e.adj.label])

				hp.Insert(heapVertex{
					v:     e.adj,
					wFrom: w.v,
					wTo:   e.adj,
					cost:  e.cost,
				})
			}
		}
	}

	return mstCostSum
}

func main() {
	//gt := newGraph()
	//
	//gt.addEdge(1, 2, 1)
	//gt.addEdge(1, 3, 4)
	//gt.addEdge(1, 4, 3)
	//gt.addEdge(2, 4, 2)
	//gt.addEdge(3, 4, 5)
	//
	//fmt.Println(primMST(gt, gt.vertices[1]))

	g := parseAdjListFileAndBuildGraph()

	mstSum := primMST(g, g.vertices[1])

	fmt.Println(mstSum)
}
