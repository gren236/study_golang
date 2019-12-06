package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	// Timer returned is a struct, containing channel and runtime timer struct (parent)
	fmt.Printf("Timer returns type %T\n", timer1)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	// If you just wanted to wait, you could have used time.Sleep.
	// One reason a timer may be useful is that you can cancel the timer before it expires.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	
	// time.Sleep(2 * time.Second)
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
