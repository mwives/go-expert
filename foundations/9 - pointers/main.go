package main

func main() {
	a := 10
	var pointer *int = &a
	*pointer = 20

	b := &a
	*b = 30

	println(*b)

	// Example 2
	num1 := 10
	num2 := 20

	sum(&num1, &num2)

	println(num1) // Prints 50
}

func sum(a, b *int) int {
	*a = 50 // Changes the value of num1 🤯
	return *a + *b
}
