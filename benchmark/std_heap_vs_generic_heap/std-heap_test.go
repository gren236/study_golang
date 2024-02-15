package std_heap_vs_generic_heap

import (
	"container/heap"
	"slices"
	"testing"
)

type stdIntSlice []int

func (h stdIntSlice) Len() int           { return len(h) }
func (h stdIntSlice) Less(i, j int) bool { return h[i] < h[j] }
func (h stdIntSlice) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *stdIntSlice) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *stdIntSlice) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func BenchmarkStdHeap(b *testing.B) {
	testSlice := make(stdIntSlice, 1000)

	for i := range 1000 {
		testSlice[i] = 1000 - i
	}

	for i := 0; i < b.N; i++ {
		t := slices.Clone(testSlice)

		heap.Init(&t)
		heap.Push(&t, 1)
		heap.Pop(&t)
	}
}
