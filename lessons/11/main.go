package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Active bool
}

// Structs
func main() {
	ivonei := Person{
		Name:   "Ivonei",
		Age:    20,
		Active: true,
	}

	ivonei.Age = 21

	fmt.Println(ivonei)
}
