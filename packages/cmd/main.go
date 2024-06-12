package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/mvives/go-expert/packages/math"
	myTime "github.com/mvives/go-expert/packages/time"
)

func main() {
	// Local Math package
	m := math.NewMath(1, 2)
	fmt.Printf("Add: %d\n", m.Add())

	// Local Time package
	myTime.Timer(2)

	// UUID package
	uuid := uuid.New().String()
	fmt.Println("UUID:", uuid)
}
