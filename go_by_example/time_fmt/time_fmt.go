package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	fmt.Printf("%T\n", p)

	// Basic example of formatting time according to RFC3339 layout
	t := time.Now()
	p(t.Format(time.RFC3339))

	// Using the same layout we can parse time to Time struct
	t1, _ := time.Parse(time.RFC3339, "2012-11-01T22:08:41+00:00")
	p(t1)

	// Datetime formats provide specific date for reference: Mon Jan 2 15:04:05 -0700 MST 2006
	// You can think of this date as: 01/02 03:04:05PM 2006 GMT-0700
	p(time.Now().Format("The time is 15:04:05 and the year is number 06"))
	p(time.Now().Format("Current month is Jan, weekday is Monday"))
	p(time.Now().Format(time.Kitchen))	// :D
	p(time.Now().Format(time.RubyDate))	// BTW this format is wrong

	// If something more precise is needed, Printf with Time functions can be used
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	// Parse will return an error explaining the problem
	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8.41PM")
	p(e)
}
