package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// Always receiving worker
	go func() {
        for {
			// "more" var will get true/false if channel open/closed
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
	}()
	
	for j := 1; j <= 3; j++ {
        jobs <- j
		fmt.Println("sent job", j)
		time.Sleep(50 * time.Millisecond)
    }
    close(jobs)
	fmt.Println("sent all jobs")
	
	// Sync
	<-done

	// We can still read from closed BUFFERED channels
	ch := make(chan string, 4)

	ch <- "foo"
	ch <- "bar"

	close(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}