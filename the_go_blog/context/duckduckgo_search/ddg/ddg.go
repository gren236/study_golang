package ddg

import (
	"context"
	"encoding/json"
	"github.com/gren236/study_golang/the_go_blog/context/duckduckgo_search/userip"
	"net/http"
)

type Result struct {
	Title, URL	string
}

type Results []Result

func httpDo(ctx context.Context, req *http.Request, f func(resp *http.Response, err error) error) error {
	// Run the HTTP request in a goroutine and pass the response to f.
	c := make(chan error, 1)
	req = req.WithContext(ctx)

	go func() { c <- f(http.DefaultClient.Do(req)) }()

	// Wait for handler to write to response, otherwise cancel if ctx.Done is closed
	select {
	case <-ctx.Done():
		<-c // Wait for f to return
		return ctx.Err()
	case err := <-c:
		return err
	}
}

func Search(ctx context.Context, query string) (Results, error) {
	// Prepare the DuckDuckGo API request
	req, err := http.NewRequest("GET", "https://api.duckduckgo.com/?q=DuckDuckGo", nil)
	if err != nil {
		return nil, err
	}
	q := req.URL.Query()
	q.Set("q", query)
	q.Set("format", "json")

	// If ctx is carrying the user IP address, forward it to the server.
	// Google APIs use the user IP to distinguish server-initiated requests from end-user requests.
	if userIP, ok := userip.FromContext(ctx); ok {
		q.Set("userip", userIP.String())
	}
	req.URL.RawQuery = q.Encode()

	var results Results
	err = httpDo(ctx, req, func(resp *http.Response, err error) error {
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// Parse the JSON search result
		var data struct {
			RelatedTopics []struct {
				Text		string
				FirstURL	string
			}
		}
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		for _, res := range data.RelatedTopics {
			results = append(results, Result{Title: res.Text, URL: res.FirstURL})
		}
		return nil
	})
	// httpDo waits for the closure we provided to return, so it's safe to
	// read results here
	return results, err
}
