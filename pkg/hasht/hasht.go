package hasht

import (
	"github.com/gren236/study_golang/pkg/linkedl"
)

const (
	initialSize     = 16
	loadThreshold   = 0.7
	growCoefficient = 2
)

type key interface {
	comparable

	Bytes() []byte
}

type container[T key, U any] struct {
	key T
	val U
}

type Table[T key, U any] struct {
	buckets *[]linkedl.Singly[*container[T, U]]
	objNum  int
}

func newBuckets[T key, U any](size int) *[]linkedl.Singly[*container[T, U]] {
	buckets := make([]linkedl.Singly[*container[T, U]], size)

	for i := range size {
		buckets[i] = linkedl.Singly[*container[T, U]]{}
	}

	return &buckets
}

func New[T key, U any]() *Table[T, U] {
	tbl := &Table[T, U]{}
	tbl.buckets = newBuckets[T, U](initialSize)

	return tbl
}

func (t *Table[T, U]) Insert(k T, v U) {
	// Check the table load
	if t.getLoad() >= loadThreshold {
		// Grow the table
		t.growTable(len(*t.buckets) * growCoefficient)
	}

	// Get a new bucket index for key and insert
	// TODO implement
}

func (t *Table[T, U]) Delete(k T) {
	// TODO implement
}

func (t *Table[T, U]) Search(k T) (res U, ok bool) {
	// TODO implement

	return
}

func (t *Table[T, U]) growTable(newSize int) {
	// TODO implement
}

func (t *Table[T, U]) getBucketIndex(k []byte) uint64 {
	hashSum := HashDJB2(k)

	return (hashSum % uint64(len(*t.buckets))) - uint64(1)
}

func (t *Table[T, U]) getLoad() float64 {
	return float64(t.objNum) / float64(len(*t.buckets))
}
