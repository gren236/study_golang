package main

import (
	"math"
)

type vertex struct {
	lbl  int // label
	x, y float64
}

func (v *vertex) distanceTo(u *vertex) float64 {
	return math.Sqrt(math.Pow(v.x-u.x, 2) + math.Pow(v.y-u.y, 2))
}

type graph struct {
	vertices  []*vertex
	unvisited map[int]*vertex
	distCache map[int]float64
}

func newGraph(length int) *graph {
	return &graph{
		vertices:  make([]*vertex, length),
		distCache: make(map[int]float64),
	}
}

func (g *graph) resetUnvisited() {
	g.unvisited = make(map[int]*vertex, len(g.vertices))

	for _, v := range g.vertices {
		g.unvisited[v.lbl] = v
	}
}

func (g *graph) findClosestUnvisited(s *vertex) (*vertex, float64) {
	var res *vertex
	minDist := math.Inf(1)

	for _, t := range g.unvisited {
		dist := s.distanceTo(t)

		if dist < minDist {
			minDist = dist
			res = t

			continue
		}

		if dist == minDist {
			if t.lbl < res.lbl {
				res = t
			}
		}
	}

	return res, minDist
}
