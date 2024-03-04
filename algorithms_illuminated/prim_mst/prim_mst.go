package main

import (
	"github.com/gren236/study_golang/pkg/heap"
	"math"
)

type heapVertex struct {
	from, to *vertex
	cost     int
}

func (h heapVertex) Less(t heapVertex) bool {
	return h.cost < t.cost
}

func primMST(g *graph, s *vertex) int {
	var mstCostSum int
	heapVertexIndices := make(map[int]int)

	hp := heap.NewHeap[heapVertex](func(h *heapVertex, i int) {
		heapVertexIndices[h.to.label] = i
	})

	g.resetExplored()
	g.explored[s.label] = s

	// Initialize heap
	for _, v := range g.vertices {
		if v == s {
			continue
		}

		currHeapVertex := heapVertex{to: v, cost: math.MaxInt}

		if cost, connected := s.checkConnection(v); connected {
			currHeapVertex.from = s
			currHeapVertex.cost = cost
		}

		hp.Insert(currHeapVertex)
	}

	// Main loop, go through all vertices in heap
	for hp.Len() > 0 {
		w := hp.ExtractMin()

		g.explored[w.to.label] = w.to

		mstCostSum += w.cost

		for _, e := range w.to.edges {
			if _, ok := g.explored[e.adj.label]; ok {
				continue
			}

			currHeapVertex := hp.GetByIndex(heapVertexIndices[e.adj.label])

			if e.cost < currHeapVertex.cost {
				hp.Delete(heapVertexIndices[e.adj.label])

				hp.Insert(heapVertex{
					from: w.to,
					to:   e.adj,
					cost: e.cost,
				})
			}
		}
	}

	return mstCostSum
}

func main() {
	// TODO Test!
}
