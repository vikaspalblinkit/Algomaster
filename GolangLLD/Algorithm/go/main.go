package main

import (
	"fmt"
	"time"
)

// func printMessage(n int) {
// 	for i := 0; i < n; i++ {
// 		fmt.Println("Hello, World!")
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }

// write a program to with channels and selects statements

// send only channels syntax we have in the systems
func primary(ch chan<- string) {
	fmt.Println("Primary function executed")
	time.Sleep(time.Second * 5)
	ch <- "Primary task completed"
}

// write in the primary and secondary channel
// function arguments should be send only channels

func secondary(ch chan<- string) {
	fmt.Println("Secondary function executed")
	time.Sleep(time.Second * 1)
	ch <- "Secondary task completed"
}

// timeout and cancellations synchorizoantion
// A select statement in Go really shines when you need to deal with multiple possible events happening concurrently, especially when:
// where you have the multiple
func main() {

	primarychan := make(chan string)
	secondarychan := make(chan string)

	go primary(primarychan)

	select {
	case msg := <-primarychan:
		fmt.Println("Received from primary:", msg)
	case <-time.After(time.Second * 2): // Timeout after 2 seconds
		go secondary(secondarychan)
		select {
		case msg := <-secondarychan:
			fmt.Println("Received from secondary:", msg)
		case <-time.After(time.Second * 2): // Timeout after 2 seconds
			fmt.Println("both api as timeout ")
		}
	}
}

/*

	Unique event_id or uuid should be used to identify the events at
	the each steps of the process
	
	Different channels in golang concurrency to handle and
	their use cases for same as well
	func kafkaConsumer(ch chan<- string) {
		for i := 1; i <= 3; i++ {
			time.Sleep(300 * time.Millisecond)
			ch <- fmt.Sprintf("Kafka msg %d", i)
		}
	}

	func rabbitConsumer(ch chan<- string) {
		for i := 1; i <= 3; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- fmt.Sprintf("RabbitMQ msg %d", i)
		}
	}

	func main() {
		kafkaCh := make(chan string)
		rabbitCh := make(chan string)

		go kafkaConsumer(kafkaCh)
		go rabbitConsumer(rabbitCh)

		for i := 0; i < 6; i++ { // total expected messages
			select {
			case msg := <-kafkaCh:
				fmt.Println("📦 Kafka:", msg)
			case msg := <-rabbitCh:
				fmt.Println("🐇 RabbitMQ:", msg)
			}
		}
	}
*/
