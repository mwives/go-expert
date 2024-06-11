package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}

	// Write
	size, err := f.Write([]byte("Hello, World!"))
	// size, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d bytes\n", size)
	f.Close()

	// Read
	file, err := os.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(file))

	// Read stream
	loremFile, err := os.Open("lorem.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(loremFile)
	buffer := make([]byte, 10)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Print(string(buffer[:n]))
	}

	// Remove
	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
