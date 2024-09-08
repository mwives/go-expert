package main

import (
	"fmt"

	"github.com/mwives/mwives-utils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
