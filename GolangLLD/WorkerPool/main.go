package main

import (
	"fmt"
	"sync"
)

/*
	Worker = GoRoutine which is fixed no of goroutines
	JobsPOOL = You have no of jobs which you want to process

*/

// type Job struct {
// 	ID      int
// 	Message string
// }

// func worker(id int, jobs <-chan Job, results chan<- string) {
// 	// Worker functions that process jobs from the jobs channels
// 	for job := range jobs {
// 		fmt.Printf("Worker %d started job %d\n", id, job.ID)
// 		time.Sleep(time.Second) // Simulate time taken to process the job
// 		fmt.Printf("Worker %d finished job %d\n", id, job.ID)
// 		results <- fmt.Sprintf("Worker %d processed job %d: %s", id, job.ID, job.Message)
// 	}
// }

// func main() {
// 	fmt.Println("Worker Pool Example Welcome to Golang!")
// 	const noOfWorkers = 3
// 	const noOfJobs = 10

// 	jobs := make(chan Job, noOfJobs)
// 	results := make(chan string, noOfJobs)

// 	// iterave over the no of workers
// 	for i := 0; i < noOfWorkers; i++ {
// 		go worker(i, jobs, results) // start a worker goroutine
// 	}

// 	for j := 1; j <= noOfJobs; j++ {
// 		jobs <- Job{ID: j, Message: fmt.Sprintf("do something %d", j)}
// 	}
// 	close(jobs) // Close the jobs channel to signal no more jobs will be sent
// 	// Collect the results
// 	for i := 0; i < noOfJobs; i++ {
// 		fmt.Println("Result:", <-results)
// 	}
// 	return
// }

/*
	How does the go.sum and go.mod works ?

	func unbufferedExample() {
    ch := make(chan string)

    go func() {
        fmt.Println("Sending...")
        ch <- "hello" // This blocks until someone receives
        fmt.Println("Sent!")
    }()

    time.Sleep(1 * time.Second)
    msg := <-ch // This receives the message
    fmt.Println("Received:", msg)

	Synchronous : Untill the message is received blocked
	No Internal storage for the same.


	Buffered Channel Example:
	1. created a buffered for a parituclar sized and then send the message
	 untill the size of bufffer is full

	2. Direction channels:
		- send onlys : ch chan<- int
		- received onlys : ch <-chan int
}
*/

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
