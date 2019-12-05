package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("processed!")

	done <- true
}

// Channel for sending messages only can be defined...
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// For receiving only as well
func pong(pings <-chan string, pongs chan<- string) {
	// msg := <-pings
	// pongs <- msg
	// ^ the same as v
	pongs <- <- pings
}

func main() {
	// Unbuffered - send value only if there is at least 1 receiver, otherwise - deadlock
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

	// Buffered - accept a limited number of values without a corresponding receiver
	ch := make(chan string, 2)

	ch <- "foo"
	ch <- "bar"
	// In this case, no deadlock without a receiver. But you can still have deadlock in case of buffer overflow!
	// ch <- "baz"

	fmt.Println(<-ch)
	fmt.Println(<-ch)

	// Sync channels using goroutines - push true when goroutine is done working
	done := make(chan bool, 1) // Buffering is not really important - it's just a safe switch for not getting deadlock
	go worker(done)

	<-done

	// Directions
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)

	// Empty interfaces allow to write differently typed values to 1 channel -_-
	ch1 := make(chan interface{}, 2)
	ch1 <- "hello"
	ch1 <- 1

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	
	// Length and capacity
	channel := make(chan string, 3)

	channel <- "foo"

	fmt.Printf("Length (number of elements) is %v, and capacity (buffer size) is %v\n", len(channel), cap(channel))

	// Creating uni-directional channels
	roc := make(<-chan int)
	soc := make(chan<- int)

	fmt.Printf("Data type of roc is %T\n", roc)
	fmt.Printf("Data type of soc is %T\n", soc)

	// Channels can be used as data type - we can make channel that stores channels
	cc := make(chan chan string)

	go func(cc chan chan string) {
		// Make a new channel and store it to channels
		c := make(chan string)
		cc <- c
	}(cc)

	// Get a stored channel from channels
	c := <-cc

	go func(c chan string) {
		fmt.Println("Hello", <-c, "!")
	}(c)

	c <- "John"
}
