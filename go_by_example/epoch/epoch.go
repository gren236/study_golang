package main

import (
	"fmt"
	"time"
)

func main() {
	// This is a special set of functions of time package to work with Unix epoch
	// Get elapsed time since UNIX epoch
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// Milliseconds should be calculated manually
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// Convert seconds or nanoseconds to Time
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
