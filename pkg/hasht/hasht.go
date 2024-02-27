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

func New[T key, U any]() *Table[T, U] {
	tbl := &Table[T, U]{}
	tbl.buckets = newBuckets[T, U](initialSize)

	return tbl
}

func (t *Table[T, U]) Insert(k T, v U) {
	if _, ok := t.Get(k); ok {
		return
	}

	// Check the table load
	if t.getLoad() >= loadThreshold {
		// Grow the table
		t.growTable(len(*t.buckets) * growCoefficient)
	}

	// Get a new bucket index for key and insert
	c := container[T, U]{k, v}
	insertContainer(*t.buckets, &c)

	t.objNum++
}

func (t *Table[T, U]) Delete(k T) {
	i := getBucketIndex(uint64(len(*t.buckets)), k.Bytes())

	deleted := (*t.buckets)[i].Delete(func(c *container[T, U]) bool {
		return c.key == k
	})
	if deleted {
		t.objNum--
	}
}

func (t *Table[T, U]) Get(k T) (res U, ok bool) {
	i := getBucketIndex(uint64(len(*t.buckets)), k.Bytes())

	cf, found := (*t.buckets)[i].Search(func(c *container[T, U]) bool {
		return c.key == k
	})
	if found {
		res, ok = cf.val, true
	}

	return
}

func (t *Table[T, U]) Len() int {
	return t.objNum
}

func (t *Table[T, U]) growTable(newSize int) {
	resized := newBuckets[T, U](newSize)

	// Go through all the current elements and rehash to new buckets
	for _, cl := range *t.buckets {
		for c, ok := cl.Next(); ok; c, ok = cl.Next() {
			insertContainer(*resized, c)
		}
	}

	t.buckets = resized
}

func (t *Table[T, U]) getLoad() float64 {
	return float64(t.objNum) / float64(len(*t.buckets))
}

func newBuckets[T key, U any](size int) *[]linkedl.Singly[*container[T, U]] {
	buckets := make([]linkedl.Singly[*container[T, U]], size)

	for i := range size {
		buckets[i] = linkedl.Singly[*container[T, U]]{}
	}

	return &buckets
}

func getBucketIndex(bSize uint64, k []byte) uint64 {
	hashSum := HashDJB2(k)

	return hashSum % bSize
}

func insertContainer[T key, U any](buckets []linkedl.Singly[*container[T, U]], c *container[T, U]) {
	i := getBucketIndex(uint64(len(buckets)), c.key.Bytes())

	buckets[i].Insert(c)
}
