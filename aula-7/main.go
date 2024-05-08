package main

import "fmt"

func main() {
	// Initializing the map
	salaries := map[string]int{"Ivonei": 666, "João": 500, "Maria": 2000}

	// Getting the value of a key
	fmt.Println(salaries["Ivonei"])

	// Deleting a key
	delete(salaries, "João")

	// Adding a new key
	salaries["José"] = 1000

	fmt.Println(salaries)

	// Creating a new map with make
	// newSalaries := make(map[string]int)
	// newSalaries["Ivonei"] = 666

	// Iterating over the map
	for name, salary := range salaries {
		fmt.Printf("The salary of %s is %d\n", name, salary)
	}
}
