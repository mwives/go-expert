package main

func main() {
	// for loop
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"one", "two", "three"}
	// range loop
	for i, number := range numbers {
		println(i, number)
	}

	// while loop
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// infinite loop
	for {
		println("Infinite loop")
		break
	}
}
