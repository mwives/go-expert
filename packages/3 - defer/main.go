package main

func main() {
	defer println("World") // Will be executed after the main function
	println("Hello")
}
