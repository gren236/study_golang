package main

import (
	"fmt"
	"net/http"
	"time"
)

// Context is provided by request, but we should derive it for any app async i/o calls
func hello(w http.ResponseWriter, r *http.Request) {
	// Context is available on every request
	ctx := r.Context()
	fmt.Println("Hello handler started")
	defer fmt.Println("Hello handler ended")

	// Simulate workload using sleep
	// Cancel working if Done channel received command
	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintln(w, "Hello!")
	case <-ctx.Done():
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":9000", nil)
}
