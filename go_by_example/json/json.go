package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page	int
	Fruits	[]string
}

// JSON output field names can be customized with struct tags
// By the way, the field itself can be named whatever you want if tag is defined for this field
type response2 struct {
	Page	int      `json:"page"`
	Fruits	[]string `json:"fruits"`
}

// Cyclic types will send Marchal to infinite loop
type Cyclic1 struct {
	Foo	interface{}
}

type Cyclic2 struct {
	Bar interface{}
}

// Family type
type FamilyMember struct {
	Name	string
	Age		int
	Parents	[]*FamilyMember
}

func main() {
	// Encoding basic data types to JSON
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Slices and maps are encoded as intended - to arrays and objects
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Custom data types can be encoded too. Includes only exported fields
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))
	// Using MarshalIndent for more human-readable output
	res1B2, _ := json.MarshalIndent(res1D, " ", "  ")
	fmt.Println(string(res1B2))

	// See how structure tags changed JSON output field names
	res2D := response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Decoding JSON values
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Prepare a var to hold unmarshaled JSON data
	var dat map[string]interface{}

	// Pay attention to how Unmarshal returns an error only and writes data to the var passed by reference
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Here we make type assertion as every value returned by JSON is an empty interface with concrete type underneath
	num := dat["num"].(float64)
	fmt.Printf("%v - %T\n", num, num)

	// Accessing nested data requires a series of type assertions
	strs := dat["strs"].([]interface{})
	fmt.Printf("%T\n", dat["strs"])
	str1 := strs[0].(string)
	fmt.Println(str1)

	// This type-assertion hell can be avoided if we map JSON to custom type like struct
	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	// We can use NewEncoder in order to stream encoded JSON directly to os.Writers or other writers
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

	// Check cyclic structures with JSON package
	c1 := Cyclic1{}
	c2 := Cyclic2{Bar: &c1}
	c1.Foo = &c2

	// Goes to infinite loop!
	//jsonstr, _ := json.Marshal(c2)
	//fmt.Println(string(jsonstr))

	// JSON can decode pointer references in structures!
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":[{"Name":"Gomez","Age":28,"Parents":[]}]}`)
	var m FamilyMember
	json.Unmarshal(b, &m)
	fmt.Printf("%#v\n", m)
	fmt.Printf("%#v\n", m.Parents[0])
}
