package main

import (
	"fmt"
	"go-expert/math"

	"github.com/google/uuid"
)

func main() {
	s := math.Sum(1, 2)
	fmt.Println("Result:", s)

	uuid := uuid.New()
	fmt.Println("UUID:", uuid)
}
