package json_vs_gob

import (
	"bytes"
	"encoding/gob"
	"testing"
)

// TestJson structures
type Test struct {
	X, Y			int64
	Name, Surname	string
	Origin			*SubTest
}

type SubTest struct {
	Set	map[int]string
}

func BenchmarkGob(b *testing.B) {
	testData := Test{
		X:       234,
		Y:       564,
		Name:    "Roger",
		Surname: "Jackson",
		Origin:  &SubTest{Set: map[int]string{0: "Hello", 1: "World", 2: "foo"}},
	}

	var network bytes.Buffer		// Network connection
	enc := gob.NewEncoder(&network)	// Will write to network
	dec := gob.NewDecoder(&network)	// Will read from network

	for n := 0; n < b.N; n++ {
		enc.Encode(testData)
		var res Test
		dec.Decode(&res)
	}
}