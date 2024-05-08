package main

import "fmt"

func main() {
	var myVar interface{} = "Hello, 世界!"
	println(myVar.(string))

	// Converting to int
	res, ok := myVar.(int)
	fmt.Printf("res: %v, ok: %v", res, ok)

	// Go says: you need ok 😠
	res2 := myVar.(int)
	fmt.Printf("res2: %v", res2)
}
