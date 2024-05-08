package main

func sum(a, b *int) int {
	*a = 50 // Changes the value of num1 🤯
	return *a + *b
}

func main() {
	num1 := 10
	num2 := 20

	sum(&num1, &num2)

	println(num1) // Prints 50
}
