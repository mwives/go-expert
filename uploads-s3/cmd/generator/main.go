package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i <= 1000; i++ {
		f, err := os.Create(fmt.Sprintf("tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		f.Close()
		f.WriteString("Hello, World!")
	}
}
