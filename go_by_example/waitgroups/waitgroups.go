package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)

		results <- j * 2
	}

	// Done working
	fmt.Printf("Worker %d done\n", id)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup // Create waitgroup

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3 worker routines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Increment counter
		go worker(&wg, i, jobs, results)
	}

	// Send 5 jobs and close channel
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait() // Blocks here until counter is 0 (all jobs are done)

	for a := 1; a <= 5; a++ {
		fmt.Println("Result", a, "is", <-results)
	}
}