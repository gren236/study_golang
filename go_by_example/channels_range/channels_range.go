package main

import (
	"fmt"
	"time"
)

func outputSquare(c <-chan int) {
	// Will loop over channel until it closes
	for elem := range c {
		fmt.Println(elem * elem)
	}
}

func main() {
	queue := make(chan int, 3)

	go outputSquare(queue)

	queue <- 1
	fmt.Println("sent")
	queue <- 2
	fmt.Println("sent")
	
	time.Sleep(time.Second)
}