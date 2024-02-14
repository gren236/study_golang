package heap

type Comparable[T any] interface {
	Less(T) bool
}

type Container[T Comparable[T]] []T

func NewHeapFromSlice[T Comparable[T]](arr []T) *Container[T] {
	res := make(Container[T], len(arr))

	// Copy array to container
	for i, v := range arr {
		res[i] = v
	}

	// Heapify container
	for i := (len(arr) / 2) - 1; i >= 0; i-- {
		res.bubbleDown(i)
	}

	return &res
}

func (c *Container[T]) ExtractMin() T {
	res := (*c)[0]

	(*c)[0], (*c)[len(*c)-1] = (*c)[len(*c)-1], (*c)[0]
	*c = (*c)[:len(*c)-1]

	// Restore heap property
	c.bubbleDown(0)

	return res
}

func (c *Container[T]) Insert(v T) {
	*c = append(*c, v)

	// Restore heap property
	c.bubbleUp(len(*c) - 1)
}

func (c *Container[T]) bubbleDown(i int) {
	for i*2+1 <= len(*c)-1 {
		iL := i*2 + 1
		iR := i*2 + 2
		rPresent := iR <= len(*c)-1

		// Check heap property
		if !rPresent && (*c)[i].Less((*c)[iL]) {
			break
		}

		if rPresent && (*c)[i].Less((*c)[iR]) && (*c)[i].Less((*c)[iL]) {
			break
		}

		iChild := iL
		if rPresent && (*c)[iR].Less((*c)[iL]) {
			iChild = iR
		}

		(*c)[i], (*c)[iChild] = (*c)[iChild], (*c)[i]

		i = iChild
	}
}

func (c *Container[T]) bubbleUp(i int) {
	for i > 0 {
		iPrt := (i - 1) / 2

		if (*c)[iPrt].Less((*c)[i]) {
			// If parent is already less than inserted, do nothing
			break
		}

		(*c)[iPrt], (*c)[i] = (*c)[i], (*c)[iPrt]

		i = iPrt
	}
}
