package kosaraju_scc

import "slices"

type vertex struct {
	Label   int       // Label of the vertex
	in, out []*vertex // incoming and outgoing edges

	explored bool // if this node was explored by DFS
	SCC      int  // number of SCC which vertex belongs to
}

func newVertex(l int) *vertex {
	return &vertex{
		Label:    l,
		in:       make([]*vertex, 0),
		out:      make([]*vertex, 0),
		explored: false,
		SCC:      -1,
	}
}

type Graph struct {
	Vertices map[int]*vertex

	topoOrder []int // topological order -> label
	curLabel  int
	sccNum    int
	reversed  bool
}

func NewGraph(size int) *Graph {
	return &Graph{
		Vertices:  make(map[int]*vertex, size),
		topoOrder: make([]int, size),
		curLabel:  0,
		sccNum:    0,
		reversed:  false,
	}
}

func (g *Graph) addVertex(label int) *vertex {
	// Check if vertex exists, otherwise create
	u, ok := g.Vertices[label]
	if !ok {
		u = newVertex(label)
		g.Vertices[label] = u
	}

	return u
}

func (g *Graph) AddEdge(from, to int) {
	// u -> v
	u := g.addVertex(from)
	v := g.addVertex(to)

	// Set v as outgoing of u
	u.out = append(u.out, v)

	// Set u as incoming of v
	v.in = append(v.in, u)
}

func (g *Graph) resetExplored() {
	for _, u := range g.Vertices {
		u.explored = false
	}
}

func (g *Graph) dfsTopo(s *vertex) {
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

	g.topoOrder[g.curLabel] = s.Label
	g.curLabel--
}

func (g *Graph) topoSort() {
	g.resetExplored()

	g.topoOrder = make([]int, len(g.Vertices))
	g.curLabel = len(g.Vertices) - 1

	for _, v := range g.Vertices {
		if !v.explored {
			g.dfsTopo(v)
		}
	}
}

func (g *Graph) dfsSCC(s *vertex) {
	s.explored = true

	s.SCC = g.sccNum // mark the SCC for specific vertex

	for _, v := range s.out {
		if !v.explored {
			g.dfsSCC(v)
		}
	}
}

func (g *Graph) MarkSCCs() {
	g.reversed = true
	g.topoSort()
	g.reversed = false

	g.resetExplored()

	g.sccNum = 0

	for _, l := range g.topoOrder {
		v := g.Vertices[l]

		if !v.explored {
			g.sccNum++

			g.dfsSCC(v)
		}
	}
}

// GetSccSizeRating returns sizes of all found SCCs in desc order
func (g *Graph) GetSccSizeRating() []int {
	res := make([]int, g.sccNum)

	for _, v := range g.Vertices {
		res[v.SCC-1]++
	}

	slices.Sort(res)
	slices.Reverse(res)

	return res
}
