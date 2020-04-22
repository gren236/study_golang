package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func main() {
	// http.Get is a shortcut for creating Client object and calling Get method
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		panic(err)
	}

	fmt.Println("Response status:", resp.Status)

	// Print first 5 lines of body
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	// MUST close Reader every time!
	resp.Body.Close()
	fmt.Println()

	// Testing JSON responses
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	fmt.Println(response.Header)

	data, _ := ioutil.ReadAll(response.Body)
	var jsonData map[string]interface{}
	json.Unmarshal(data, &jsonData)
	fmt.Println(jsonData)

	// Download image using http client
	resp, err = http.Get("https://images.pexels.com/photos/34153/pexels-photo.jpg")
	if err != nil {
		panic(err)
	}

	// Copy received bytes to file without dumping it all to memory
	if resp.Header.Get("Content-Type") == "image/jpeg" {
		file, err := os.OpenFile("image.jpg", os.O_CREATE | os.O_WRONLY, 0775)
		defer file.Close()
		if err != nil {
			panic(err)
		}

		// We could just use io.Copy - it's for studying purposes :)
		// io.Copy(file, resp.Body)
		buf := make([]byte, 32)
		for {
			n, err := resp.Body.Read(buf)

			_, ferr := file.Write(buf[:n])
			if ferr != nil {
				panic(ferr)
			}

			if err != nil {
				break
			}
		}
	}

	resp.Body.Close()
	fmt.Println()

	// Send basic POST request
	// We can also use PostForm to "post" data of "application/x-www-form-urlencoded"
	// Make struct and turn it into JSON string
	jdata, err := json.Marshal(
		struct {
		title, body string
		userId int
	} {"test", "Lorem ipsum", 7})
	reqData := bytes.NewReader(jdata)

	res, err := http.Post(
		"https://jsonplaceholder.typicode.com/posts",
		"application/json; charset=UTF-8",
		reqData,
	)
	if err != nil {
		panic(err)
	}

	data, _ = ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n", data)

	// Creating custom request object
	// Parse request URL
	reqUrl, _ := url.Parse("https://jsonplaceholder.typicode.com/posts")

	// Add request body
	reqBody := ioutil.NopCloser(strings.NewReader(`{"title":"testing","body":"Lorem Ipsum","userId":43}`))

	// Build request object
	req := &http.Request{
		Method: "POST",
		URL:    reqUrl,
		Header: map[string][]string{"Content-Type": {"application/json; charset=UTF-8"}},
		Body:   reqBody,
	}

	// Send built request using default client
	rsp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	bdata, err := ioutil.ReadAll(rsp.Body)
	rsp.Body.Close()

	fmt.Println("Status:", rsp.Status)
	fmt.Println("Response:", string(bdata))
	// Alternatively, we could use NewRequest() to build request object faster
	fmt.Println()

	// Create custom http client
	// NOTE: Default http.Client timeout value is 0! So we should always use custom timeouts!
	client := &http.Client{
		Timeout: 2 * time.Millisecond,
	}

	// Send sample GET request using this client
	clresp, err := client.Get("https://jsonplaceholder.typicode.com/posts/1")

	// Check for response error
	if err != nil {
		// Get underlying urlError struct pointer
		urlErr := err.(*url.Error)
		// Check if error occurred due to timeout
		if urlErr.Timeout() {
			fmt.Println("Timeout error!")
		}
		// Just panic if it's something else
		panic(err)
	} else {
		fmt.Println("Status:", clresp.Status)
	}
}
