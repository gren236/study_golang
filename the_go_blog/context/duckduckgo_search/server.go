// I had to change Google search to DuckDuckGo as Google's not providing
// an open search API anymore :(
package main

import (
	"context"
	"github.com/gren236/study_golang/the_go_blog/context/duckduckgo_search/ddg"
	"github.com/gren236/study_golang/the_go_blog/context/duckduckgo_search/userip"
	"html/template"
	"log"
	"net/http"
	"time"
)

func handleSearch(w http.ResponseWriter, r *http.Request) {
	// ctx is the Context for this handler. Calling cancel closes the
	// ctx.Done channel, which is the cancellation signal for requests
	// started by this handler
	var (
		ctx		context.Context
		cancel	context.CancelFunc
	)
	timeout, err := time.ParseDuration(r.FormValue("timeout"))
	if err == nil {
		// The request has a timeout, so create a context ths is
		// canceled automatically when the timeout expires
		ctx, cancel = context.WithTimeout(context.Background(), timeout)
	} else {
		ctx, cancel = context.WithCancel(context.Background())
	}
	defer cancel()

	// Check the search query
	query := r.FormValue("q")
	if query == "" {
		http.Error(w, "no query", http.StatusBadRequest)
	}

	// Store the user IP in ctx for use by code in other packages
	userIP, err := userip.FromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx = userip.NewContext(ctx, userIP)

	// Run the DuckDuckGo search and print the results
	start := time.Now()
	results, err := ddg.Search(ctx, query)
	elapsed := time.Since(start)

	if err := resultsTemplate.Execute(w, struct {
		Results          ddg.Results
		Timeout, Elapsed time.Duration
	}{
		Results: results,
		Timeout: timeout,
		Elapsed: elapsed,
	}); err != nil {
		log.Print(err)
		return
	}
}

func main() {
	http.HandleFunc("/search", handleSearch)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

var resultsTemplate = template.Must(template.New("results").Parse(`
<html>
<head/>
<body>
  <ol>
  {{range .Results}}
    <li>{{.Title}} - <a href="{{.URL}}">{{.URL}}</a></li>
  {{end}}
  </ol>
  <p>{{len .Results}} results in {{.Elapsed}}; timeout {{.Timeout}}</p>
</body>
</html>
`))
