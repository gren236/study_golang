package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// A basic streaming entity, implementing Reader
type MyStringData struct {
	str string
	readIndex int // default 0
}

func NewStringData(str string) *MyStringData {
	return &MyStringData{str: str}
}

func (msd *MyStringData) Rewind() {
	msd.readIndex = 0
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

// A basic custom writer
type CustomStorage struct {
	data []byte
}

func NewCustomStorage(data []byte) *CustomStorage {
	return &CustomStorage{data: data}
}

func (cs *CustomStorage) Write(p []byte) (n int, err error) {
	// Check if 10 bytes were written
	if len(cs.data) == 10 {
		return 0, io.EOF
	}

	// Get remaining capacity
	remainingCap := 10 - len(cs.data)

	// Get length of data to write
	writeLength := len(p)
	if writeLength >= remainingCap {
		writeLength = remainingCap
		err = io.EOF
	}

	// Append data
	cs.data = append(cs.data, p[:writeLength]...)

	// Set number of bytes written and return
	n = writeLength
	return
}

func main() {
	// Reading from custom reader
	customString := NewStringData("Something!")

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
	strR := strings.NewReader("Hello!")
	strD := make([]byte, 5)

	for {
		_, err := strR.Read(strD)

		fmt.Println(string(strD))

		if err != nil {
			fmt.Println("EOF")
			break
		}
	}
	
	// Check writer work
	cs := NewCustomStorage([]byte{})
	n, err := cs.Write([]byte("hello!"))
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v bytes written: %s\n", n, cs.data)

	// More useful function to write strings is io.WriteString
	cs1 := NewCustomStorage([]byte{})
	n1, err := io.WriteString(cs1, "Foo Bar")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v bytes written: %s\n", n1, cs1.data)

	// Standard io streams are: os.Stdin, os.Stdout, os.Stderr
	// Basically, fmt package is just a WriteString wrapper around standard streams.
	fmt.Fprintln(os.Stderr, "Some error")
	// We can write directly to Stdout
	os.Stdout.WriteString("Hi mark!\n")

	// Values from io.Reader can be copied to io.Writer
	customString.Rewind()
	cs2 := NewCustomStorage([]byte{})
	n3, err := io.Copy(cs2, customString)
	fmt.Printf("%v bytes written: %s\n", n3, cs2.data)

	// We can communicate between readers and writers using io.Pipe
	src, dst := io.Pipe()

	// Launch goroutine that writes data to Pipe writer
	go func() {
		dst.Write([]byte("ABCDEF"))
		dst.Write([]byte("GHIJKL"))
		dst.Close()
	}()

	// Read from src will block main goroutine until something appears for reading
	pipeData := make([]byte, 6)
	for {
		n, err := src.Read(pipeData)
		fmt.Printf("%v bytes read: %s\n", n, pipeData[:n])

		if err == io.EOF {
			break
		}
	}

	// Buffer example
	buf := bytes.NewBufferString("Hello world!")
	strReader := strings.NewReader(" Doing ok?")
	from, err := buf.ReadFrom(strReader)
	fmt.Println("Bytes written:", from)
	buf.WriteTo(os.Stdout)
	fmt.Println()
	// Buffer can be used to accumulate data before sending to consumer
	// Provided by bufio package, which is basically a Writer wrapper
	// buffer can be flushed with all it's contents to underlying io.Writer passed to buffer
	buf1 := bufio.NewWriter(os.Stdout)
	buf1.WriteString("Hello there!\n")
	buf1.Flush()
}
