package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"net/url"
	"os"
	"sync"
	"time"
)

const numWorkers = 10

type page struct {
	url string
	body io.ReadCloser
	err error
}

type result struct {
	url string
	hash []byte
}

// 3 stages: get all links, follow them, output every followed page's content hashed
// parseLinks is emitting links for downstream stages
func parseLinks(done <-chan struct{}, rootUrl string) (<-chan *url.URL, <-chan error) {
	res := make(chan *url.URL)
	errc := make(chan error)

	go func() {
		defer close(res)
		client := &http.Client{
			Timeout: 5 * time.Second,
		}
		resp, err := client.Get(rootUrl)
		if err != nil {
			errc <- err
			return
		}
		defer resp.Body.Close()

		base, err := url.Parse(rootUrl)
		if err != nil {
			errc <- err
			return
		}

		z := html.NewTokenizer(resp.Body)
		for {
			if z.Next() == html.ErrorToken {
				// Returning io.EOF indicates success.
				break
			}
			if tagn, isAttr := z.TagName(); string(tagn) == "a" && isAttr {
				key, val, _ := z.TagAttr()
				if string(key) == "href" {
					u, _ := url.Parse(string(val))
					select {
					case res <- base.ResolveReference(u):
					case <-done:
						errc <- errors.New("link parsing canceled")
						return
					}
				}
			}
		}
	}()

	return res, errc
}

// followUrls is a worker that follows links and receives body readers
func followUrls(done <-chan struct{}, links <-chan *url.URL, pages chan<- *page) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	for link := range links {
		page := page{
			url: link.String(),
		}
		resp, err := client.Get(link.String())
		if err != nil {
			page.err = err
			continue
		}
		page.body = resp.Body
		select {
		case pages <- &page:
		case <-done:
			return
		}
	}
}

// hashPages get the SHA256 hash of a followed pages
func hashPages(done chan struct{}, rootUrl string) <-chan result {
	// Get followed links bodies
	links, _ := parseLinks(done, rootUrl)
	pages := make(chan *page)
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			followUrls(done, links, pages)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(pages)
	}()

	// Process bodies
	var wgh sync.WaitGroup
	res := make(chan result)
	wgh.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for page := range pages {
				if page.err != nil {
					continue
				}

				h := sha256.New()
				if _, err := io.Copy(h, page.body); err != nil {
					continue
				}
				page.body.Close()
				select {
				case res <- result{page.url, h.Sum(nil)}:
				case <-done:
					return
				}
			}
			wgh.Done()
		}()
	}
	go func() {
		wgh.Wait()
		close(res)
	}()
	return res
}

func main() {
	urlArg := os.Args[1]
	done := make(chan struct{})
	pages := hashPages(done, urlArg)

	for page := range pages {
		fmt.Printf("%s : %x\n", page.url, page.hash)
	}
}
