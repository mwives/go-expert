package main

import "fmt"

func main() {
	// Original slice
	s := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Printf("Original Slice: len=%d cap=%d %v\n", len(s), cap(s), s)

	// Slice with length 0
	fmt.Printf("Slice with length 0: len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])

	// Slice with length 4
	fmt.Printf("Slice with length 4: len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])

	// Slice from index 2 to end
	fmt.Printf("Slice from index 2 to end: len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	// Append 110 to the slice
	s = append(s, 110)
	fmt.Printf("Original Slice: len=%d cap=%d %v\n", len(s), cap(s), s)
}
