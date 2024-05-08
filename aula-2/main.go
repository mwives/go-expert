package main

const a = "Hello, World!"

var (
	b bool    = true
	c int     = 10
	d string  = "Ivonei"
	e float64 = 1.2
)

func main() {
	a := "type inference"
	a = "new value"

	println(a) // Prints "new value"
	println(b)
	println(c)
	println(d)
}

func external() {
	println(a) // Prints "Hello, World!"
}
