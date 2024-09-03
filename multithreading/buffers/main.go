package main

func main() {
	// Create a buffered channel that can hold 2 string values
	ch := make(chan string, 2)

	// Send "Hello" and "World" into the channel
	ch <- "Hello"
	ch <- "World"

	// Receive and print the values from the channel
	println(<-ch) // Prints "Hello"
	println(<-ch) // Prints "World"
}
