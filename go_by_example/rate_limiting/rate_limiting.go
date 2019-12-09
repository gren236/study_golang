package main

import (
	"fmt"
	"time"
)

func main() {
	// Basic rate-limiting
	requests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        requests <- i
    }
	close(requests)
	
	// Limiter limits or rate with 1 request per 200 milliseconds
	limiter := time.Tick(200 * time.Millisecond)

	for req := range requests {
        <-limiter // Here loop blocks to wait for another tick
        fmt.Println("request", req, time.Now())
	}

	// With bursts of requests (still limited)
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
	}
	
	// Add ticks to bursty limiter
	go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
	}()
	
	// Make 5 more requests
	burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
	close(burstyRequests)
	
    for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}