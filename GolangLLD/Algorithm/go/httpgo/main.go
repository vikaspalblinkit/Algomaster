package main

import "sync"

type Result struct {
	URL  string
	Err  error
	Body string
}

func fetch(url string, results chan<- Result, wg *sync.WaitGroup) {

}

func main() {
	// worker pool pattern
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
}
