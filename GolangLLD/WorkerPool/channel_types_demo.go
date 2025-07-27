package main

import (
	"fmt"
	"time"
)

func unbufferedDemo() {
	fmt.Println("\n=== Unbuffered Channel Demo ===")
	ch := make(chan string)

	go func() {
		fmt.Println("Goroutine: About to send...")
		ch <- "Hello from unbuffered"
		fmt.Println("Goroutine: Sent! (this prints after receiver reads)")
	}()

	time.Sleep(1 * time.Second) // Simulate some work
	fmt.Println("Main: About to receive...")
	msg := <-ch
	fmt.Println("Main: Received:", msg)
}

func bufferedDemo() {
	fmt.Println("\n=== Buffered Channel Demo ===")
	ch := make(chan string, 2) // Buffer size 2

	go func() {
		fmt.Println("Goroutine: About to send first message...")
		ch <- "First message"
		fmt.Println("Goroutine: Sent first! (this prints immediately)")

		fmt.Println("Goroutine: About to send second message...")
		ch <- "Second message"
		fmt.Println("Goroutine: Sent second! (this prints immediately)")

		fmt.Println("Goroutine: About to send third message...")
		ch <- "Third message" // This will block because buffer is full
		fmt.Println("Goroutine: Sent third! (this prints after receiver reads)")
	}()

	time.Sleep(2 * time.Second) // Let goroutine send first two messages
	fmt.Println("Main: About to receive...")
	
	msg1 := <-ch // This unblocks the third send
	fmt.Println("Main: Received:", msg1)
	
	msg2 := <-ch
	fmt.Println("Main: Received:", msg2)
	
	msg3 := <-ch
	fmt.Println("Main: Received:", msg3)
}

func directionalDemo() {
	fmt.Println("\n=== Directional Channels Demo ===")
	
	// Function that only sends
	sender := func(ch chan<- int) {
		for i := 1; i <= 3; i++ {
			ch <- i
			fmt.Printf("Sent: %d\n", i)
		}
		close(ch)
	}
	
	// Function that only receives
	receiver := func(ch <-chan int) {
		for num := range ch {
			fmt.Printf("Received: %d\n", num)
		}
	}
	
	ch := make(chan int, 3)
	go sender(ch)
	receiver(ch)
}

func nilChannelDemo() {
	fmt.Println("\n=== Nil Channel Demo (Select Statement) ===")
	
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	var nilCh chan string // This is nil
	
	ch1 <- "from ch1"
	ch2 <- "from ch2"
	
	for i := 0; i < 3; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("Received:", msg)
			ch1 = nilCh // Disable this case by setting to nil
		case msg := <-ch2:
			fmt.Println("Received:", msg)
			ch2 = nilCh // Disable this case by setting to nil
		default:
			fmt.Println("All channels are nil or empty")
		}
	}
}

func runChannelDemo() {
	fmt.Println("Channel Types Demonstration")
	
	unbufferedDemo()
	bufferedDemo()
	directionalDemo()
	nilChannelDemo()
}

// To run this demo, rename this function to main() 
// or call runChannelDemo() from another main function
