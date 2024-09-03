package main

import (
	"fmt"
	"sync"
	"time"
)

// Main function spawns multiple goroutines to run tasks concurrently
func main() {
	// Initialize a WaitGroup to synchronize the completion of tasks
	wg := sync.WaitGroup{}
	// Add 25 to the WaitGroup counter (total work to be done)
	wg.Add(25)

	// Start a goroutine to run task A
	go task("A", &wg)
	// Start a goroutine to run task B
	go task("B", &wg)
	// Start an anonymous goroutine to run another task
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("%d: Task %s is running\n", i, "anonymous")
			time.Sleep(1 * time.Second)
			// Signal that one iteration of the anonymous task is done
			wg.Done()
		}
	}()

	// Allow time for all goroutines to complete (should use wait groups instead)
	// time.Sleep(time.Second * 10)

	// Block the main thread until all tasks have completed
	wg.Wait()
}

// Simulate a task that runs for 10 iterations, signaling completion to the WaitGroup
func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		// Signal that one iteration of the task is done
		wg.Done()
	}
}
