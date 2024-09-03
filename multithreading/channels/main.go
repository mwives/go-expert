package main

import (
	"fmt"
	"time"
)

func main() {
	// Create an unbuffered int channel
	data := make(chan int)

	// Number of worker goroutines to spawn
	numWorkers := 100
	// Number of data items to send through the channel
	numData := 1000

	// Start <numWorkers> worker goroutines
	for i := 0; i < numWorkers; i++ {
		go worker(i, data)
	}

	// Send <numData> items to the workers through the channel
	for i := 0; i < numData; i++ {
		data <- i
	}
}

// Worker function that receives data from the channel and processes it
// data is a receive-only channel
func worker(workerId int, data <-chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		// Simulate work by sleeping for 1 second
		time.Sleep(time.Second)
	}
}
