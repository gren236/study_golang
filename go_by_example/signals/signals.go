package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Go signal notification works by sending os.Signal values on a channel
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// Notify registers the given channel to receive notifications
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine execs a blocking receive for signals
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// The program will wait here until it gets the expected signal and then exits
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
