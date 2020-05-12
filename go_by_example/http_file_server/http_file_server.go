package main

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"path/filepath"
)

func main() {
	// Create file server handler.
	fs := http.FileServer(http.Dir("."))

	// Handle '/' route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("Content-Type", "text/html")
		fmt.Fprint(res, "<h1>Golang!</h1>")
	})

	// Handle '/static' route. Use strip prefix to remove 'static' pathname for a handler
	http.Handle("/static/", http.StripPrefix("/static", fs))

	// Handle specific file path
	http.HandleFunc("/readme", func(res http.ResponseWriter, req *http.Request) {
		name := "README.md"
		res.Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(name)))
		http.ServeFile(res, req, filepath.FromSlash("./" + name))
	})

	// Start HTTP server with a default handler
	log.Fatal(http.ListenAndServe(":9001", nil))
}
