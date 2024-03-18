package main

type edge struct {
	id   string
	cost int

	v, w vertex
}

type byCost []edge

func (c byCost) Len() int {
	return len(c)
}

func (c byCost) Less(i, j int) bool {
	return c[i].cost < c[j].cost
}

func (c byCost) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

type vertex int

type graph struct {
	edges    []edge
	vertices []vertex
}

func newGraph() *graph {
	return &graph{
		edges:    make([]edge, 0),
		vertices: make([]vertex, 0),
	}
}

func (g *graph) addEdge(e edge) {
	g.edges = append(g.edges, e)
}

func (g *graph) fillVerticesSequentially(n int) {
	g.vertices = make([]vertex, n)

	for i := range n {
		g.vertices[i] = vertex(i + 1)
	}
}
