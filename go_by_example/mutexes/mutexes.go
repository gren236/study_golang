package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	// Keep track of read/write operations
	var readOps uint64
	var writeOps uint64

	// Simulate reads
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				// Pick a random key
				key := rand.Intn(5)
				// Lock the mutex to ensure exclusive access
				mutex.Lock()
				// Read the value
				total += state[key]
				// Unlock mutex
				mutex.Unlock()
				// Increment read operations
				atomic.AddUint64(&readOps, 1)

				// Wait a bit between read operations
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Simulate writes, using the same pattern
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Let them work a bit
	time.Sleep(time.Second)

	// Final operation counters reports
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// With a final lock of state, show how it ended up (main() is a routine too!)
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()

	// TODO: How can we break this? :)
}
