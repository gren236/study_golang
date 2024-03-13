package ufind

import "errors"

type Node[T comparable] struct {
	parent T
	rank   int
}

type Container[T comparable] map[T]*Node[T]

func New[T comparable](data []T) (Container[T], error) {
	if data == nil || len(data) <= 0 {
		return nil, errors.New("empty data provided")
	}

	res := make(Container[T], len(data))

	for _, d := range data {
		res[d] = &Node[T]{
			parent: d,
			rank:   0,
		}
	}

	return res, nil
}

func (c Container[T]) Find(x T) (res T, ok bool) {
	_, fnd := c[x]
	if !fnd {
		return
	}

	return c.getRoot(x), true
}

func (c Container[T]) Union(x, y T) bool {
	rootX, fnd := c.Find(x)
	if !fnd {
		return false
	}

	rootY, fnd := c.Find(y)
	if !fnd {
		return false
	}

	// If ranks are equal - choose arbitrary as parent and up it's rank by 1
	if c[rootX].rank == c[rootY].rank {
		c[rootY].parent = rootX
		c[rootX].rank++

		return true
	}

	// If ranks are different - put one root under another with a greater rank
	if c[rootX].rank > c[rootY].rank {
		c[rootY].parent = rootX
	} else {
		c[rootX].parent = rootY
	}

	return true
}

func (c Container[T]) getRoot(x T) T {
	p := c[x].parent

	if x == p {
		return x
	}

	root := c.getRoot(c[x].parent)

	// Path compression
	c[x].parent = root

	return root
}
