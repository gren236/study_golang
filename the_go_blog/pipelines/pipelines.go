package main

import (
	"fmt"
	"sync"
)

// Let's look at a pipeline consisting of 3 stages

// First stage - receive numbers and send them to a channel (making an emitter)
func gen(nums ...int) <-chan int {
	out := make(chan int)

	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()

	return out
}

// Second stage - receiving numbers through channel and returning a channel with numbers' squares
func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n*n
		}
		close(out)
	}()

	return out
}

// merge merges received channels to a single one
func merge(cs... <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(ch <-chan int) {
		for n := range ch {
			out <- n
		}
		wg.Done()
	}

	// Add wait group according to number of received inbound channels (Nice tip!)
	wg.Add(len(cs))
	for _, c := range cs {
		output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Main function is the final stage - receiving values from second stage and printing them
func main() {
	c := gen(2, 3, 4, 5)
	out := sq(c)

	for n := range out {
		fmt.Println(n)
	}
	fmt.Println()

	// We can run pipeline as we want. Even call stages multiple times:
	for n := range sq(sq(gen(2, 3, 4))) {
		fmt.Println(n)
	}
	fmt.Println()

	// Fan-in usage - distribute multiple inbound channels between the second stage to parallelize work
	in := gen(2, 3, 4, 5)

	c1 := sq(in)
	c2 := sq(in)

	for n := range merge(c1, c2) {
		fmt.Println(n)
	}
}
