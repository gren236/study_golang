package linkedl

type Node[T any] struct {
	n *Node[T]
	v T
}

// Singly represents a singly linked list by storing a pointer to its head.
type Singly[T any] struct {
	head   *Node[T]
	cursor *Node[T]
	size   int
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

func (s *Singly[T]) Len() int {
	return s.size
}

func (s *Singly[T]) Extract() T {
	res := s.head.v

	s.head = s.head.n
	s.size--

	return res
}

func (s *Singly[T]) Delete(eq func(T) bool) (ok bool) {
	var prev *Node[T]
	cur := s.head

	for cur != nil {
		if eq(cur.v) {
			// Rewire pointers to exclude this node
			if prev != nil {
				prev.n = cur.n
			} else {
				s.head = cur.n
			}

			s.size--

			return true
		}

		prev = cur
		cur = cur.n
	}

	return
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

// Next returns subsequent element from the Node stored in cursor. If cursor is nil, then returns head. Moves the cursor.
func (s *Singly[T]) Next() (T, bool) {
	res := s.head

	if s.cursor != nil {
		if s.cursor.n == nil {
			return *new(T), false
		}

		res = s.cursor.n
	}

	s.cursor = res

	return res.v, true
}

func (s *Singly[T]) Reset() {
	s.cursor = nil
}
