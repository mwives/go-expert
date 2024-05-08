package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Person struct {
	Name   string
	Age    int
	Active bool
	Address
}

// Structs
func main() {
	ivonei := Person{
		Name:   "Ivonei",
		Age:    20,
		Active: true,
	}

	ivonei.Age = 21
	ivonei.Address.City = "Hollyudi"

	fmt.Println(ivonei)
}
