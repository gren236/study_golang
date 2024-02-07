package main

import "slices"

type vertex struct {
	label   int       // label of the vertex
	in, out []*vertex // incoming and outgoing edges

	explored bool // if this node was explored by DFS
	scc      int  // number of SCC which vertex belongs to
}

func newVertex(l int) *vertex {
	return &vertex{
		label:    l,
		in:       make([]*vertex, 0),
		out:      make([]*vertex, 0),
		explored: false,
		scc:      -1,
	}
}

type graph struct {
	vertices map[int]*vertex

	topoOrder []int // topological order -> label
	curLabel  int
	sccNum    int
	reversed  bool
}

func newGraph() *graph {
	return &graph{
		vertices:  make(map[int]*vertex),
		topoOrder: make([]int, 0),
		curLabel:  0,
		sccNum:    0,
		reversed:  false,
	}
}

func (g *graph) addVertex(label int) *vertex {
	// Check if vertex exists, otherwise create
	u, ok := g.vertices[label]
	if !ok {
		u = newVertex(label)
		g.vertices[label] = u
	}

	return u
}

func (g *graph) addEdge(from, to int) {
	// u -> v
	u := g.addVertex(from)
	v := g.addVertex(to)

	// Set v as outgoing of u
	u.out = append(u.out, v)

	// Set u as incoming of v
	v.in = append(v.in, u)
}

func (g *graph) resetExplored() {
	for _, u := range g.vertices {
		u.explored = false
	}
}

func (g *graph) dfsTopo(s *vertex) {
	s.explored = true

	sAdj := s.out

	if g.reversed {
		sAdj = s.in
	}

	for _, v := range sAdj {
		if !v.explored {
			g.dfsTopo(v)
		}
	}

	g.topoOrder[g.curLabel] = s.label
	g.curLabel--
}

func (g *graph) topoSort() {
	g.resetExplored()

	g.topoOrder = make([]int, len(g.vertices))
	g.curLabel = len(g.vertices) - 1

	for _, v := range g.vertices {
		if !v.explored {
			g.dfsTopo(v)
		}
	}
}

func (g *graph) dfsSCC(s *vertex) {
	s.explored = true

	s.scc = g.sccNum

	for _, v := range s.out {
		if !v.explored {
			g.dfsSCC(v)
		}
	}
}

func (g *graph) markSCCs() {
	g.reversed = true
	g.topoSort()
	g.reversed = false

	g.resetExplored()

	g.sccNum = 0

	for _, l := range g.topoOrder {
		v := g.vertices[l]

		if !v.explored {
			g.sccNum++

			g.dfsSCC(v)
		}
	}
}

// getSccSizeRating returns sizes of all found SCCs in desc order
func (g *graph) getSccSizeRating() []int {
	res := make([]int, g.sccNum)

	for _, v := range g.vertices {
		res[v.scc-1]++
	}

	slices.Sort(res)
	slices.Reverse(res)

	return res
}
