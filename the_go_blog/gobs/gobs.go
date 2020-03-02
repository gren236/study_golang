package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// Gob works with structs only
type P struct {
	X, Y, Z	int
	Name	string
}

type Q struct {
	X, Y	*int32
	Name	string
}

func main() {
	// Initialize encoder and decoder. Normally, these would be different processes or machines
	var network bytes.Buffer		// Network connection
	enc := gob.NewEncoder(&network)	// Will write to network
	dec := gob.NewDecoder(&network)	// Will read from network

	// Encode the value (send)
	err := enc.Encode(P{3, 4, 5, "Pythagoras"})
	// Check errors
	if err != nil {
		panic(err)
	}

	fmt.Println(network)

	// Decode (receive) the value
	var q Q
	err = dec.Decode(&q)
	// Check errors
	if err != nil {
		panic(err)
	}

	fmt.Printf("%q: {%d:%d}\n", q.Name, *q.X, *q.Y)

	// Non-struct values are encoded as 1 field structs and decoded back to plain type

	err = enc.Encode([]string{"Hello", "World"})
	if err != nil {
		fmt.Println("Encoding error:", err)
	}
	var res []string
	err = dec.Decode(&res)
	if err != nil {
		fmt.Println("Decoding error:", err)
	}

	fmt.Printf("%#v\n", res)

	// Signed int encoding example
	var i int = -13
	fmt.Printf("i - %b\n", i)
	var u uint
	if i < 0 {
		// Cast signed int to unsigned
		u2 := uint(i)
		fmt.Printf("u2 - %b : %v\n", u2, u2)
		// Inverse all bits
		uxor := ^u2
		fmt.Printf("uxor - %b\n", uxor)
		// Shift bits to left to get a pare bit for sign
		ushift := uxor << 1
		fmt.Printf("ushift - %b\n", ushift)
		// xxxxx | 00001 always returns xxxx1
		u = ushift | 1 // complement i, bit 0 is 1
	} else {
		u = (uint(i) << 1) // do not complement i, bit 0 is 0
	}

	fmt.Printf("u - %b\n", u)
}
