package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]bool
	mux sync.Mutex
}

// Visited checks if the url has already been visited, if not, it flags it as visited.
func (c *SafeCounter) Visited(url string) bool {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	_, ok := c.v[url]
	if !ok {
		c.v[url] = true
	}
	return ok
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// Don't fetch the same URL twice.
	if c.Visited(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	// Fetch URLs in parallel.
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher, wg) // Start a new goroutine for each new URL to fetch.
	}
}

var c SafeCounter = SafeCounter{v: make(map[string]bool)}

func main() {

	var wg sync.WaitGroup
	wg.Add(1) // Add the initial call for Crawl function.

	// Start the crawling process in a separate goroutine.
	go Crawl("https://golang.org/", 4, fetcher, &wg)

	wg.Wait() // Wait for all the Crawl goroutines to finish.

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
