package main

import (
	"errors"
	"math/rand"
	"slices"
)

type vertex struct {
	label int
	adj   []*vertex
}

func (v *vertex) moveEdges(dst *vertex) {
	for _, nbr := range v.adj {
		for j, u := range nbr.adj {
			if u == v {
				nbr.adj[j] = dst
				dst.adj = append(dst.adj, nbr)
			}
		}
	}
}

func (v *vertex) clearSelfLoops() {
	resAdj := make([]*vertex, 0)

	for _, u := range v.adj {
		if u != v {
			resAdj = append(resAdj, u)
		}
	}

	v.adj = resAdj
}

type graph struct {
	nextLabel int
	labels    []int
	vertices  map[int]*vertex
}

func newGraph() *graph {
	return &graph{
		nextLabel: 1,
		labels:    make([]int, 0),
		vertices:  make(map[int]*vertex),
	}
}

func (g *graph) addVertex(l int, adj []*vertex) *vertex {
	if adj == nil {
		adj = make([]*vertex, 0)
	}

	g.labels = append(g.labels, l)
	g.nextLabel = l + 1
	g.vertices[l] = &vertex{label: l, adj: adj}

	return g.vertices[l]
}

func (g *graph) deleteVertex(u *vertex) {
	uIndex := slices.Index(g.labels, u.label)
	g.labels[uIndex], g.labels[len(g.labels)-1] = g.labels[len(g.labels)-1], g.labels[uIndex]
	g.labels = g.labels[:len(g.labels)-1]

	delete(g.vertices, u.label)
}

func (g *graph) getRandomEdge() (*vertex, *vertex) {
	u := g.vertices[g.labels[rand.Intn(len(g.labels))]]
	v := u.adj[rand.Intn(len(u.adj))]

	return u, v
}

func (g *graph) merge(u, v *vertex) {
	// create a new vertex (super-node)
	dstLabel := g.nextLabel
	dst := g.addVertex(dstLabel, nil)

	// move edges from u
	u.moveEdges(dst)

	// move edges from v
	v.moveEdges(dst)

	// delete self loops from new node
	dst.clearSelfLoops()

	// delete u and v
	g.deleteVertex(u)
	g.deleteVertex(v)
}

func (g *graph) contract() {
	for len(g.vertices) > 2 {
		u, v := g.getRandomEdge()

		g.merge(u, v)
	}
}

func (g *graph) assertContractionSuccess() (int, error) {
	if len(g.labels) != 2 || len(g.vertices) != 2 {
		return 0, errors.New("wrong number of vertices")
	}

	u := g.vertices[g.labels[0]]
	v := g.vertices[g.labels[1]]

	if len(u.adj) != len(v.adj) {
		return 0, errors.New("wrong number of edges")
	}

	for _, nbr := range u.adj {
		if nbr != v {
			return 0, errors.New("u has wrong adj vertices")
		}
	}

	for _, nbr := range v.adj {
		if nbr != u {
			return 0, errors.New("v has wrong adj vertices")
		}
	}

	return len(u.adj), nil
}
