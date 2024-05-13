package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Ivonei"
	e float64 = 1.2
	f ID      = 666
)

func main() {
	a := "type inference"
	a = "new value"

	println(a) // Prints "new value"
	println(b)
	println(c)
	println(d)
	println(e)
	println(f)
	fmt.Printf("The type of f is %T", f)
}

func external() {
	println(a) // Prints "Hello, World!"
}
