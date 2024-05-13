package main

import "fmt"

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World!"

	x = "Fala fiote" // x can have value of another type assigned

	showType(x)
	showType(y)

	// Empty interface conversion
	var myVar interface{} = "Hello, 世界!"
	println(myVar.(string))

	// Converting to int
	res, ok := myVar.(int)
	fmt.Printf("res: %v, ok: %v", res, ok)

	// Go says: you need ok 😠
	res2 := myVar.(int)
	fmt.Printf("res2: %v", res2)
}

func showType(t interface{}) {
	fmt.Printf("The type of t is %T and the value is %v\n", t, t)
}
