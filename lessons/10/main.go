package main

import "fmt"

// Closures
func main() {
	total := func() int { return sum(123, 243, 147, 153) }() // Anonymous function

	fmt.Println(total)
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
