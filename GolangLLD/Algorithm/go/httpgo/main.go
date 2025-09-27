package main

import "sync"

type Result struct {
	URL  string
	Err  error
	Body string
}

func fetch(url string, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	// Simulate fetching the URL
	results <- Result{URL: url, Err: nil, Body: "Response body"}  
}

func main() {
	// worker pool pattern 
	// Sender: chan<-  and Receiver channel: <-chan 
	urls := []string{
		"https://example.com/api1",
		"https://example.com/api2",
		"https://example.com/api3",
	}

	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, results, &wg)
	}

	wg.Wait()
	close(results) 

	for res := range results {
		if res.Err != nil {
			// Handle error
			continue
		}
		// Process the response body
		println("Fetched:", res.URL, "Body:", res.Body)
	}
}

