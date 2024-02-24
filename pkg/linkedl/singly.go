package linkedl

type Node[T any] struct {
	n *Node[T]
	v T
}

// Singly represents a singly linked list by storing a pointer to its head.
type Singly[T any] struct {
	head *Node[T]
	size int
}

func NewSingly[T any](vals ...T) *Singly[T] {
	res := &Singly[T]{}

	if vals != nil {
		for _, v := range vals {
			res.Insert(v)
		}
	}

	return res
}

func (s *Singly[T]) Insert(v T) {
	s.head = &Node[T]{n: s.head, v: v}
	s.size++
}

func (s *Singly[T]) Extract() T {
	res := s.head.v

	s.head = s.head.n

	return res
}

func (s *Singly[T]) Search(eq func(T) bool) (res T, ok bool) {
	cur := s.head

	for cur != nil {
		if eq(cur.v) {
			return cur.v, true
		}

		cur = cur.n
	}

	return
}
