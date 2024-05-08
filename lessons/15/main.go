package main

func main() {
	a := 10
	var pointer *int = &a
	*pointer = 20

	b := &a
	*b = 30

	println(*b)
}
