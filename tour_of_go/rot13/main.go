package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func Rot13Encrypt(b byte) byte {
	result := b
	// Uppercase symbols
	if b >= 65 && b <= 90 {
		result = b + 13
		if result > 90 {
			result = result - 26
		}
	}
	// Regular symbols
	if b >= 97 && b <= 122 {
		result = b + 13
		if result > 122 {
			result = result - 26
		}
	}
	return result
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	_, err := reader.r.Read(bytes)
	
	for i, v := range bytes {
		bytes[i] = Rot13Encrypt(v)
	}
	
	return len(bytes), err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
