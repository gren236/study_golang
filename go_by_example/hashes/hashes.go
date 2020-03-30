package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "Sha1 this string!"

	// Start a new hash instance
	h := sha1.New()

	// Write bytes of string to hasher
	h.Write([]byte(s))

	// Receive the hash as a byte slice
	bs := h.Sum(nil)

	// SHA1 values are usually printed as hex (ex. Git)
	fmt.Println(s)
	fmt.Println(bs)
	fmt.Printf("%x\n", bs)

	// Similar approach can be used for other hashing functions available: MD5, SHA256, etc.
	sm := "ZA WARUDO"
	hm := md5.New()
	hm.Write([]byte(sm))
	bsm := hm.Sum(nil)

	fmt.Println(bsm)
	fmt.Printf("%x\n", bsm)
}
