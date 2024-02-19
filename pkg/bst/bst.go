package bst

type Comparable[T any] interface {
	Less(T) bool
	Equals(T) bool
}

type Node[T Comparable[T]] struct {
	val T

	p, l, r *Node[T] // parent, left child, right child
}

// Insert adds given value to the tree, returning node containing value. If value is present in the tree - just returns it.
func (n *Node[T]) Insert(k T) *Node[T] {
	cur := n

	for cur != nil {
		v, l, r := cur.val, cur.l, cur.r

		if k.Equals(v) {
			return cur
		}

		if k.Less(v) {
			if l == nil {
				cur.l = &Node[T]{val: k, p: cur}

				return cur.l
			}

			cur = l
		} else {
			if r == nil {
				cur.r = &Node[T]{val: k, p: cur}

				return cur.r
			}

			cur = r
		}
	}

	return cur
}

func (n *Node[T]) Search(k T) (res *Node[T], ok bool) {
	cur := n

	for cur != nil {
		v, l, r := cur.val, cur.l, cur.r

		if k.Equals(v) {
			res = cur
			ok = true
		}

		if k.Less(v) {
			cur = l
		} else {
			cur = r
		}
	}

	return
}

// Delete removes node with given value from the tree and returns a new root
func (n *Node[T]) Delete(k T) (*Node[T], bool) {
	dn, ok := n.Search(k)
	if !ok {
		return n, false
	}

	root := n

	// Has no children
	if dn.l == nil && dn.r == nil {
		dn.deleteWithoutChildren()
	}

	// Has 1 child
	if (dn.l == nil) != (dn.r == nil) {
		dn.deleteWithOneChild()
	}

	// Has multiple children
	if dn.l != nil && dn.r != nil {
		pred, _ := dn.Pred()

		// swap dn and pred
		if dn.p != nil {
			if dn.p.l == dn {
				dn.p.l = pred
			} else {
				dn.p.r = pred
			}
		} else {
			root = pred
		}

		if pred.p.l == pred {
			pred.p.l = dn
		} else {
			pred.p.r = dn
		}

		pred.p, dn.p = dn.p, pred.p
		pred.l, dn.l = dn.l, pred.l
		pred.r, dn.r = dn.r, pred.r

		pred.l.p = pred
		pred.r.p = pred

		// delete dn on new place
		if dn.l == nil && dn.r == nil {
			dn.deleteWithoutChildren()
		} else {
			dn.deleteWithOneChild()
		}
	}

	return root, true
}

func (n *Node[T]) deleteWithoutChildren() {
	if n.p.l == n {
		n.p.l = nil
	} else {
		n.p.r = nil
	}

	n.p = nil
}

func (n *Node[T]) deleteWithOneChild() {
	child := n.l
	if n.l == nil {
		child = n.r
	}

	if n.p != nil {
		if n.p.l == n {
			n.p.l = child
		} else {
			n.p.r = child
		}
	}

	child.p = n.p
}

func (n *Node[T]) Min() (res *Node[T]) {
	last := n

	for last != nil {
		res, last = last, last.l
	}

	return
}

func (n *Node[T]) Max() (res *Node[T]) {
	last := n

	for last != nil {
		res, last = last, last.r
	}

	return
}

func (n *Node[T]) Pred() (res *Node[T], ok bool) {
	if n.l != nil {
		return n.l.Max(), true
	}

	cur := n.p
	for cur != nil {
		if cur.p != nil && cur.p.r == cur {
			return cur.p, true
		}

		cur = cur.p
	}

	return
}

func (n *Node[T]) Succ() (res *Node[T], ok bool) {
	if n.r != nil {
		return n.r.Min(), true
	}

	cur := n.p
	for cur != nil {
		if cur.p != nil && cur.p.l == cur {
			return cur.p, true
		}

		cur = cur.p
	}

	return
}
