package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	// Get current time
	now := time.Now()
	p(now)

	// Build time struct from scratch passing year, month, day, etc.
	then := time.Date(
		2009, 3, 14, 23, 50, 13, 651387237, time.UTC,
		)
	p(then)

	// Get exact component of a time
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	// Weekday is also possible to get
	p(then.Weekday())

	// Comparing times
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// Also, we can get Duration struct representing interval between two datetimes
	diff := now.Sub(then)
	p(diff)

	// Get interval in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Using Add we can advance in time, or go back by passing negative value
	p(then.Add(diff))
	p(then.Add(-diff))
}