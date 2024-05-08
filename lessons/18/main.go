package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	x = "Fala fiote" // x can have value of another type assigned

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("The type of t is %T and the value is %v\n", t, t)
}
