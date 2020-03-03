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
	fmt.Printf("%#v\n", now)

	// Build time struct from scratch passing year, month, day, etc.
	// Returns Time, which stores datetime as seconds and nanoseconds bitwise.
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
	// Diff stores interval as a number of nanoseconds between dates, so max interval is 290 years.
	diff := now.Sub(then)
	p(diff)

	// Parse duration with format
	// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h"
	dur, _ := time.ParseDuration("1h5m50µs")
	p(dur)

	// Manual duration init
	var cdur time.Duration = 1 << 63 - 1
	p(cdur)

	// Get interval in various units
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	// Time object does not have Milliseconds method, but Duration has
	p(diff.Milliseconds())
	p(diff.Nanoseconds())

	// Using Add we can advance in time, or go back by passing negative value
	p(then.Add(diff))
	p(then.Add(-diff))

	// Async functions
	aft := time.After(time.Second * 3)
	p(<-aft)
}