package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// generate a certificate struct
	cert, _ := tls.LoadX509KeyPair("localhost.crt", "localhost.key")
	
	// Create a custom server with TLS config
	s := &http.Server{
		Addr:    ":9000",
		Handler: nil,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	// Handle '/' route
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World!")
	})

	// run server on port 9000
	log.Fatal(s.ListenAndServeTLS("", ""))
}
