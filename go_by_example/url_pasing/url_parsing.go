package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	s := "mysql://user:pass@host.com:5432/path?k=v&foo=bar&foo=42#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing parsed scheme
	fmt.Println(u.Scheme)

	// User info
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	// Second param is passwordSet flag - true if password is given
	fmt.Println(u.User.Password())

	// Host contains both host and port if present
	fmt.Println(u.Host)
	fmt.Println(net.SplitHostPort(u.Host))
	// Special Hostname and Port functions are also present
	fmt.Println(u.Hostname())
	fmt.Println(u.Port())

	fmt.Println(u.Path)
	fmt.Println(u.Fragment)
	fmt.Println(u.RequestURI())

	// Get raw query string
	fmt.Println(u.RawQuery)
	// Get parsed query string (map)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Printf("%T - %v\n", m["foo"][1], m["foo"][1])
}
