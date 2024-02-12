package main

import "maps"

type edge struct {
	weight int
	adj    *vertex
}

type vertex struct {
	label int
	edges []edge

	minPath int
}

func newVertex(l int) *vertex {
	return &vertex{label: l, edges: make([]edge, 0), minPath: 1000000}
}

func (v *vertex) addEdge(t *vertex, w int) {
	v.edges = append(v.edges, edge{
		weight: w,
		adj:    t,
	})

	t.edges = append(t.edges, edge{
		weight: w,
		adj:    v,
	})
}

type graph struct {
	vertices   map[int]*vertex
	unexplored map[int]*vertex
}

func newGraph() *graph {
	return &graph{
		vertices:   make(map[int]*vertex),
		unexplored: make(map[int]*vertex),
	}
}

func (g *graph) addVertex(l int) *vertex {
	g.vertices[l] = newVertex(l)
	return g.vertices[l]
}

func (g *graph) addEdge(l1, l2, w int) {
	u, ok := g.vertices[l1]
	if !ok {
		u = g.addVertex(l1)
	}

	v, ok := g.vertices[l2]
	if !ok {
		v = g.addVertex(l2)
	}

	u.addEdge(v, w)
}

func (g *graph) searchNextMinScore() *vertex {
	var resW *vertex
	minScore := 1000000

	for _, w := range g.unexplored {
		for _, e := range w.edges {
			v := e.adj

			// Check if target not preset in unexplored
			_, ok := g.unexplored[v.label]
			if !ok {
				// Not present in unexplored, so explored
				if curScore := v.minPath + e.weight; curScore < minScore {
					resW = w
					minScore = curScore
				}
			}
		}
	}

	// Set minPath of w to minScore (if found)
	if resW != nil {
		resW.minPath = minScore
	}

	return resW
}

func (g *graph) setMinDistances(s *vertex) {
	maps.Copy(g.unexplored, g.vertices)
	delete(g.unexplored, s.label)

	s.minPath = 0

	for w := g.searchNextMinScore(); w != nil; w = g.searchNextMinScore() {
		delete(g.unexplored, w.label)
	}
}
