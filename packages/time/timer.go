package time

import (
	"fmt"
	"time"
)

func Timer(seconds int) {
	duration := time.Duration(seconds) * time.Second
	timer := time.NewTimer(duration)
	fmt.Println("Timer started")
	<-timer.C
	fmt.Println("Timer expired")
}
