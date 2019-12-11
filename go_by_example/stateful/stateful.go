package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// Read operation request
type readOp struct {
	key int
	resp chan int
}

// Write operation request
type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	// Count operations
	var readOps uint64
	var writeOps uint64
	
	// Channels to perform reads and writes
	reads := make(chan readOp)
	writes := make(chan writeOp)
	
	// This goroutine is a single responsible owner of state
	go func() {
		var state = make(map[int]int)
		// Repeatedly selecting read and write channels
		for {
            select {
			case read := <-reads:
				// Get read struct and send to response channel
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
	}()

	// Create 100 read goroutines
	for r := 0; r < 100; r++ {
        go func() {
            for {
				// Create read request
                read := readOp{
                    key:  rand.Intn(5),
					resp: make(chan int),
				}
				// Send constructed read to reads channel
				reads <- read
				// Get the response
				<-read.resp
				// Add new read operation to counter
				atomic.AddUint64(&readOps, 1)
				// Add some time to operation
                time.Sleep(time.Millisecond)
            }
        }()
	}
	
	// Create 10 write goroutines
	for w := 0; w < 10; w++ {
        go func() {
            for {
				// Similar approach here
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
					resp: make(chan bool),
				}
                writes <- write
                <-write.resp
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
	}
	
	// Let it work for a second
	time.Sleep(time.Second)

	// Capture and report op counts
	readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)
}