package linkedl

type Node[T any] struct {
	n *Node[T]
	v T
}

type Single[T any] *Node[T]

func NewSingle[T any](vals ...T) *Single[T] {
	return (*Single[T])()
}
