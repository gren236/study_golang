package heap

type Comparable[T any] interface {
	Less(T) bool
}

type Container[T Comparable[T]] struct {
	storage       []T
	onIndexUpdate func(*T, int)
}

func NewHeap[T Comparable[T]](updateIndex func(*T, int)) *Container[T] {
	return &Container[T]{
		storage:       make([]T, 0),
		onIndexUpdate: updateIndex,
	}
}

func NewHeapFromSlice[T Comparable[T]](arr []T) *Container[T] {
	res := Container[T]{
		storage: make([]T, len(arr)),
	}

	// Copy array to container
	for i, v := range arr {
		res.storage[i] = v
	}

	// Heapify container
	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		res.bubbleDown(i)
	}

	return &res
}

func (c *Container[T]) ExtractMin() T {
	res := c.storage[0]

	c.Delete(0)

	// Restore heap property
	c.bubbleDown(0)

	return res
}

func (c *Container[T]) Delete(i int) {
	if i < 0 || i >= len(c.storage) {
		return
	}

	c.storage[i], c.storage[len(c.storage)-1] = c.storage[len(c.storage)-1], c.storage[i]
	c.storage = c.storage[:len(c.storage)-1]

	if c.Len() == i {
		return
	}

	if c.onIndexUpdate != nil {
		c.onIndexUpdate(&c.storage[i], i)
	}

	// Restore heap property
	if !c.bubbleDown(i) {
		c.bubbleUp(i)
	}

	return
}

func (c *Container[T]) Insert(v T) {
	c.storage = append(c.storage, v)

	if c.onIndexUpdate != nil {
		c.onIndexUpdate(&c.storage[len(c.storage)-1], len(c.storage)-1)
	}

	// Restore heap property
	c.bubbleUp(len(c.storage) - 1)
}

func (c *Container[T]) GetByIndex(i int) (res T) {
	if i < 0 || i >= len(c.storage) {
		return
	}

	res = c.storage[i]

	return
}

func (c *Container[T]) PeekMin() T {
	return c.storage[0]
}

func (c *Container[T]) Len() int {
	return len(c.storage)
}

func (c *Container[T]) CheckMinCorrect() bool {
	if c.Len() <= 0 {
		return true
	}

	minHeap := c.storage[0]

	minCurr := minHeap
	for _, v := range c.storage {
		if v.Less(minCurr) {
			minCurr = v
		}
	}

	if minCurr.Less(minHeap) || minHeap.Less(minCurr) {
		return false
	}

	return true
}

func (c *Container[T]) GetMinLinear() T {
	minCurr := c.storage[0]
	for _, v := range c.storage {
		if v.Less(minCurr) {
			minCurr = v
		}
	}

	return minCurr
}

func (c *Container[T]) bubbleDown(i0 int) bool {
	updateHook := c.onIndexUpdate != nil

	i := i0
	for i*2+1 <= len(c.storage)-1 {
		iL := i*2 + 1
		iR := i*2 + 2
		rPresent := iR <= len(c.storage)-1

		// Check heap property
		if !rPresent && c.storage[i].Less(c.storage[iL]) {
			break
		}

		if rPresent && c.storage[i].Less(c.storage[iR]) && c.storage[i].Less(c.storage[iL]) {
			break
		}

		iChild := iL
		if rPresent && c.storage[iR].Less(c.storage[iL]) {
			iChild = iR
		}

		c.storage[i], c.storage[iChild] = c.storage[iChild], c.storage[i]
		if updateHook {
			c.onIndexUpdate(&c.storage[i], i)
			c.onIndexUpdate(&c.storage[iChild], iChild)
		}

		i = iChild
	}

	return i > i0
}

func (c *Container[T]) bubbleUp(i int) {
	updateHook := c.onIndexUpdate != nil

	for i > 0 {
		iPrt := (i - 1) / 2

		if c.storage[iPrt].Less(c.storage[i]) {
			// If parent is already less than inserted, do nothing
			break
		}

		c.storage[iPrt], c.storage[i] = c.storage[i], c.storage[iPrt]
		if updateHook {
			c.onIndexUpdate(&c.storage[i], i)
			c.onIndexUpdate(&c.storage[iPrt], iPrt)
		}

		i = iPrt
	}
}
