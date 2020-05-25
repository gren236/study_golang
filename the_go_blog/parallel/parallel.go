package main

import (
	"fmt"
	"runtime"
)

// Using GOMAXPROCS value so user can rewrite the env variable
var numCPU = runtime.GOMAXPROCS(0)

type Vector []float64

func (v Vector) Op(e float64) float64 {
	return e * e
}

// DoSome applies the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
	for ; i < n; i++ {
		v[i] += u.Op(v[i])
	}
	c <- 1 // signal that this piece is done
}

func (v Vector) DoAll(u Vector) {
	c := make(chan int, numCPU)
	// n = 6
	// 0 0 1
	// 1 1 3
	// 2 3 4
	// 3 4 6
	for i := 0; i < numCPU; i++ {
		go v.DoSome(i * len(v) / numCPU, (i + 1) * len(v) / numCPU, u, c)
	}
	// Drain the channel
	for i := 0; i < numCPU; i++ {
		<-c
	}
	// All done
}

func main() {
	v := Vector{1, 2, 3, 4, 5, 6}
	v.DoAll(Vector{})
	fmt.Println(v)
}