### Worker Pool Implementation

```go 
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
``` 


### Implemenation of waitgroups we haved 

```go 
func WorkerJob(wg *sync.WaitGroup, jobs <-chan string, results chan<- string) {
	defer wg.Done()
	for job := range jobs {
		// Simulate processing the job
		results <- "Processed: " + job // Send the result back to the results channel
	}
}

func main() {
	const noOfJobs = 10
	const noOfWorkers = 3
	jobs := make(chan string, noOfJobs)
	results := make(chan string, noOfWorkers)

	var wg sync.WaitGroup

	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go WorkerJob(&wg, jobs, results)
	}

	for j := 1; j <= noOfJobs; j++ {
		jobs <- "Job " + string(j) // Send job to the channel
	}

	wg.Wait()
	close(jobs)
	for res := range results {
		fmt.Println("Result:", res)
	}
	return
}
```

### Mutex/Lock/Synchronizations:
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	1. WaitGroups and Synchronization/Semaphores concepts in the concurrency levels
	2. sync.WaitGroup is a synchronization primitive that allows you to wait for a collection of goroutines to finish executing.

	Three Steps:
	1. Add
	2. Done
	3. Wait

	without WaitGroup, you would have to use
	channels or other synchronization methods to ensure
	that the main goroutine waits for all worker goroutines to finish before exiting.
	This can lead to more complex code and potential


	Merging the results from mutliple
	go routines and then finally merging the results.

	Run Parallel DB queries and return the
	results in a single channels
*/

/*
	Locking and Semaphores apart from the Synchronization techniques
	we need to understand that well in the advanced.

	Semphores: Basicallly concurrency ko control krta hai hai kitne goroutines
	ko ek time par run hone ki permission hai. Resource ko limit bhi krta hai.
	Mulitples allowed hoota hai.

	Mutex: Ek resource ko access karne ke liye ek goroutine ko lock lena padta hai,


	Locking: Ek resource ko access karne ke liye ek goroutine ko lock lena padta hai,
	taaki dusre goroutines us resource ko access na kar sakein.



	Buffereed : Loose synchoroization : batch processing, logging pipleines etc.
	// non blocking calls.

	Unbufffereed channels : strict synchroization, find the values for the max in the current.
	Send blocks unitl the receiver. 



	Prevents Goroutine leaks: If a goroutine is waiting on a channel that no one is reading from, it will block forever.
	1. context.context for cancellations and timeouts. 
	2. Adding the timeouts for sames. 
	

	Race conditions in the goroutines: 
	1. Occurred when multiple goroutines try to access shared data concurrently..
	2. Prevents : by mutex and channeles 

	Select stmt used for non-blocking and adding the timouts for the channels

	Explain Fan-in and Fan-out concurrency patterns? 
	// 



	Profiling and Tracking : 
	pprof tool in Go is used to profile the performance of Go programs.
	

*/

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Worker", id, "is starting")
	time.Sleep(time.Second) // Simulate some work
	fmt.Println("Worker", id, "has finished")
}

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i < 3; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}
// 	wg.Wait() // Wait for all workers to finish
// 	fmt.Println("All workers have completed")
// 	return
// }

var counter int
var mu sync.Mutex

func main() {
	/*
		sem := make(chan struct{}, 2) // Semaphore with a capacity of 2 // buffered channels
		var wg sync.WaitGroup
		for i := 0; i < 5; i++ {
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				sem <- struct{}{}        // Acquire a token
				defer func() { <-sem }() // Release the token

				// Simulate work
				fmt.Println("Worker", id, "is starting")
				time.Sleep(time.Second)
				fmt.Println("Worker", id, "has finished")
			}(i)
		}
		wg.Wait()
		fmt.Println("All workers have been started")
	*/

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait() // wait for all the goroutines to finish
}

// implemented a worker pool pattern
// find the value of the max from the worker pools
// and then return the max values from the worker pool


// go run -race main.go its provide in built runtimes providers 
// 


/*
PostgresSQL allows the list paritionining in the database
1. Allow the list partitioninng in the databases.
2.

Prevention:

1. sync.Mutex → protects critical sections.
2. sync/atomic → for simple counters/flags.
3. Channels → Go’s idiomatic way:
    don’t share state,
    communicate instead.
*/


### Find the values from the implementation: 
``` 
``` 