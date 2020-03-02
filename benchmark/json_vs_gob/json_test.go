package json_vs_gob

import (
	"encoding/json"
	"testing"
)

// TestJson structures
type TestJson struct {
	X, Y			int64
	Name, Surname	string
	Origin			*SubTestJson
}

type SubTestJson struct {
	Set	map[int]string
}

func BenchmarkJson(b *testing.B) {
	testData := TestJson{
		X:       234,
		Y:       564,
		Name:    "Roger",
		Surname: "Jackson",
		Origin:  &SubTestJson{Set: map[int]string{0: "Hello", 1: "World", 2: "foo"}},
	}

	for n := 0; n < b.N; n++ {
		dat, _ := json.Marshal(testData)
		var res TestJson
		json.Unmarshal(dat, &res)
	}
}
