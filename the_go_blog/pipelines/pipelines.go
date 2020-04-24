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
func sq(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()

	return out
}

// merge merges received channels to a single one
func merge(done <-chan struct{}, cs... <-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	// Start an output goroutine for each input channel in cs.  output
	// copies values from c to out until c or done is closed, then calls
	// wg.Done
	output := func(ch <-chan int) {
		defer wg.Done()
		for n := range ch {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}

	// Add wait group according to number of received inbound channels (Nice tip!)
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

// Main function is the final stage - receiving values from second stage and printing them
//
// Here are the guidelines for pipeline construction:
// * stages close their outbound channels when all the send operations are done.
// * stages keep receiving values from inbound channels until those channels are closed or the senders are unblocked.
// Pipelines unblock senders either by ensuring there's enough buffer for all the values that are sent or by explicitly
// signalling senders when the receiver may abandon the channel.
func main() {
	// Basic usage
	//c := gen(2, 3, 4, 5)
	//out := sq(c)
	//
	//for n := range out {
	//	fmt.Println(n)
	//}
	//fmt.Println()

	// We can run pipeline as we want. Even call stages multiple times:
	//for n := range sq(sq(gen(2, 3, 4))) {
	//	fmt.Println(n)
	//}
	//fmt.Println()

	// Fan-in usage - distribute multiple inbound channels between the second stage to parallelize work
	in := gen(2, 3, 4, 5)

	// Set up a done channel that's shared by the whole pipeline,
	// and close that channel when this pipeline exits, as a signal
	// for all the goroutines we started to exit.
	done := make(chan struct{})
	defer close(done)

	c1 := sq(done, in)
	c2 := sq(done, in)

	for n := range merge(done, c1, c2) {
		fmt.Println(n)
	}
	// Done will be closed by a deferred call
}
