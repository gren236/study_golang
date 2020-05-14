package main

import (
	"fmt"
	"net/http"
	"sync"
)

func createServer(name string, port int) *http.Server {
	// Create ServerMux
	mux := http.NewServeMux()

	// Create a default route handler
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello: " + name)
	})
	
	// Create new server
	server := http.Server{
		Addr:    fmt.Sprintf(":%v", port),
		Handler: mux,
	}

	// return new server
	return &server
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	// Create a default route handler
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello: " + request.Host)
	})

	// Spawn a goroutine to listen 9000
	go func() {
		server := createServer("ONE", 9000)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	// Spawn a goroutine to listen 9001
	go func() {
		server := createServer("TWO", 9001)
		fmt.Println(server.ListenAndServe())
		wg.Done()
	}()

	wg.Wait()
}
