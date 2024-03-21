package main

type Node struct {
	l, r *Node

	weight, maxRank, minRank int
}

func (n *Node) Less(v *Node) bool {
	return n.weight < v.weight
}

func NewNode(w int) *Node {
	return &Node{weight: w}
}

func (n *Node) Merge(v *Node) *Node {
	return &Node{
		l:       n,
		r:       v,
		weight:  n.weight + v.weight,
		maxRank: max(n.maxRank, v.maxRank) + 1,
		minRank: min(n.minRank, v.minRank) + 1,
	}
}
