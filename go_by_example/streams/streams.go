package main

import (
	"fmt"
	"io"
	"strings"
)

// A basic streaming entity, implementing Reader
type MyStringData struct {
	str string
	readIndex int // default 0
}

func New(str string) *MyStringData {
	return &MyStringData{str: str}
}

func (MyStringData *MyStringData) Read(p []byte) (n int, err error) {
	strData := MyStringData.str
	// Check if index is greater than string length
	if MyStringData.readIndex >= len(strData) {
		return 0, io.EOF
	}

	// Get next readable limit
	nextLimit := MyStringData.readIndex + len(p)

	if nextLimit >= len(strData) {
		nextLimit = len(strData)
		err = io.EOF
	}

	// Read data from string
	nextBytes := []byte(strData[MyStringData.readIndex:nextLimit])
	n = len(nextBytes)

	// Copy to received buffer
	//copy(p, nextBytes) // use with caution, because overlapping may occur!
	for i := range p {
		if i >= len(nextBytes) {
			p[i] = 0
			continue
		}
		p[i] = nextBytes[i]
	}

	// Increment ReadIndex
	MyStringData.readIndex = nextLimit

	return
}

func main() {
	// Reading from custom reader
	customString := New("Something is going on here!")

	data := make([]byte, 5)
	for {
		_, err := customString.Read(data)

		fmt.Println(string(data))

		if err != nil {
			fmt.Println("EOF")
			break
		}
	}

	// But strings package provides more convenient reader
	strR := strings.NewReader("Hello there mate!")
	strD := make([]byte, 5)

	for {
		_, err := strR.Read(strD)

		fmt.Println(string(strD))

		if err != nil {
			fmt.Println("EOF")
			break
		}
	}
}
