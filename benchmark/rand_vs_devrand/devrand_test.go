package rand_vs_devrand

import (
	"os"
	"testing"
)

func BenchmarkDevrand(b *testing.B) {
	file, err := os.OpenFile("/dev/urandom", os.O_RDONLY, 0775)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	data := make([]byte, 10)
	// Main benchmark loop (until b.N)
	for n := 0; n < b.N; n++ {
		file.Read(data)
	}
}
