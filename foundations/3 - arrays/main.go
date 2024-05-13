package main

import "fmt"

const a = "Hello, World!"

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v := range myArray {
		fmt.Printf("The value of index %d is %d\n", i, v)
	}
}
