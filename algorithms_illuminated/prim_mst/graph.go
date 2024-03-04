package main

type edge struct {
	cost int
	adj  *vertex
}

type vertex struct {
	label int
	edges map[int]edge // key = label of adj, value = cost
}

func newVertex(l int) *vertex {
	return &vertex{label: l, edges: make(map[int]edge)}
}

func (v *vertex) addEdge(t *vertex, c int) {
	v.edges[t.label] = edge{
		cost: c,
		adj:  t,
	}

	t.edges[v.label] = edge{
		cost: c,
		adj:  v,
	}
}

// checkConnection checks if receiver vertex is connected to argument, and if so, returns the cost of this edge.
func (v *vertex) checkConnection(t *vertex) (int, bool) {
	if e, ok := v.edges[t.label]; ok {
		return e.cost, true
	}

	return 0, false
}

type graph struct {
	vertices map[int]*vertex
	explored map[int]*vertex
}

func newGraph() *graph {
	return &graph{
		vertices: make(map[int]*vertex),
		explored: make(map[int]*vertex),
	}
}

func (g *graph) addVertex(l int) *vertex {
	g.vertices[l] = newVertex(l)
	return g.vertices[l]
}

func (g *graph) addEdge(l1, l2, c int) {
	u, ok := g.vertices[l1]
	if !ok {
		u = g.addVertex(l1)
	}

	v, ok := g.vertices[l2]
	if !ok {
		v = g.addVertex(l2)
	}

	u.addEdge(v, c)
}

func (g *graph) resetExplored() {
	g.explored = make(map[int]*vertex)
}
