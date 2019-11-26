package main

import (
	"fmt"
	"time"
)

func f(from string) {
	time.Sleep(15 * time.Millisecond)
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func getChars(s string) {
	for _, c := range s {
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("%c ", c)
	}
}

func getDigits(s []int) {
	for _, d := range s {
		time.Sleep(20 * time.Millisecond)
		fmt.Printf("%d ", d)
	}
}

func main() {
	fmt.Println("Start")

	// Goroutines are concurrent, not parallel by default
	go getChars("Hello")

	go getDigits([]int{1, 2, 3, 4, 5})

	go func() {
		time.Sleep(30 * time.Millisecond)
		fmt.Println("Hello!")
	}()

	time.Sleep(200 * time.Millisecond)
	fmt.Println("Done")
}
