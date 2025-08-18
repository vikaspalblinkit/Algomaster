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

/*
PostgresSQL allows the list paritionining in the database
1. Allow the list partitioninng in the databases.
2.

*/
