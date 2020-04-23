package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom handler without HandleFunc declaration
type HTTPHandler struct{}

func (h HTTPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Straight buffer write
	data := []byte("Foobar!\n")
	w.Write(data)
	// Using WriteString
	io.WriteString(w, "Hello Everynyan!")
	// Using fmt package
	fmt.Fprintln(w, "😅")
}

// Main thing in HTTP servers is handler concept, an object implementing http.Handler interface
// Handlers take incoming request and response with data written to responseWriter
func hello(w http.ResponseWriter, r *http.Request) {
	// Just write basic response to writer
	fmt.Fprintf(w, "Hello\n")
}

// This function respond with request headers
func headers(w http.ResponseWriter, r *http.Request) {
	for name, header := range r.Header {
		for _, v := range header {
			fmt.Fprintf(w, "%s : %s\n", name, v)
		}
	}
}

// Default page
func index(w http.ResponseWriter, r *http.Request) {
	// Set custom header
	w.Header().Set("X-Custom-Header", "Hello there!")
	// Set content type
	w.Header().Set("Content-Type", "application/json")

	// Set status header
	w.WriteHeader(http.StatusNotFound)

	fmt.Fprintln(w, `{"STATUS": 404}`)
}

func main() {
	// Getting instance of server multiplexer
	mux := http.NewServeMux()

	// Using HandleFunc we register handler functions
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)
	// Now we registering default page for "/" route. It's going to be called on every route pattern mismatch
	mux.HandleFunc("/", index)
	// If using default multiplexer, it's possible to call global http.HandleFunc

	// Finally call ListenAndServe. Second param is a specific handler, "nil" means default
	//err := http.ListenAndServe(":8090", nil)
	//err := http.ListenAndServe(":8090", HTTPHandler{})
	err := http.ListenAndServe(":8090", mux)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
