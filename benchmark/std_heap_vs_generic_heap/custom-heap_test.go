package std_heap_vs_generic_heap

import (
	"github.com/gren236/study_golang/pkg/heap"
	"slices"
	"testing"
)

type customInt int

func (ci customInt) Less(t customInt) bool {
	return ci < t
}

func BenchmarkCustomHeap(b *testing.B) {
	testSlice := make([]customInt, 1000)

	for i := range 1000 {
		testSlice[i] = customInt(1000 - i)
	}

	for i := 0; i < b.N; i++ {
		t := slices.Clone(testSlice)

		tc := heap.NewHeapFromSlice(t)
		tc.Insert(1)
		tc.ExtractMin()
	}
}
