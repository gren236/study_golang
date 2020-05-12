package main

import (
	"context"
	"fmt"
	"sync"
)

type myKey int

const key myKey = 0

// Main rule is: incoming requests should create a context,
// outgoing i/o call should accept a context
func main() {
	// Create the context that can be canceled
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Use waitgroup for orchestration
	var wg sync.WaitGroup
	wg.Add(10)

	// Create ten goroutines the will derive a Context from the one created above.
	for i := 0; i < 10; i++ {
		go func(id int) {
			defer wg.Done()

			// Derive a new context for this goroutine from the main func's Context.
			ctx := context.WithValue(ctx, key, id)

			// Wait until the context is cancelled.
			<-ctx.Done()
			fmt.Println("Canceled:", id)
		}(i)
	}

	// Cancel the context and any derived contexts as well.
	cancel()
	wg.Wait()
}
