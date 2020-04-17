package rand_vs_devrand

import (
	"crypto/rand"
	"testing"
)

func BenchmarkRand(b *testing.B) {
	data := make([]byte, 10)
	// Main benchmark loop (until b.N)
	for n := 0; n < b.N; n++ {
		rand.Read(data)
	}
}
