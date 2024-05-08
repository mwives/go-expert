package main

import "fmt"

func main() {
	fmt.Println(sum(123, 243, 147, 153))
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}
