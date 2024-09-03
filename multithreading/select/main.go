package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// Message struct represents a message with an ID and content
type Message struct {
	id  int64
	msg string
}

func main() {
	// Create two channels to simulate RabbitMQ and Kafka message streams
	c1 := make(chan Message)
	c2 := make(chan Message)

	var i int64 = 0

	// Simulate RabbitMQ publisher
	go func() {
		for {
			// Increment the message ID atomically
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from RabbitMQ"}
			// Send the message to channel c1
			c1 <- msg
		}
	}()

	// Simulate Kafka publisher
	go func() {
		for {
			// Increment the message ID atomically
			atomic.AddInt64(&i, 1)
			time.Sleep(time.Second * 1)
			msg := Message{i, "Hello from Kafka"}
			// Send the message to channel c2
			c2 <- msg
		}
	}()

	// Continuously listen for messages on both channels
	for {
		select {
		case msg1 := <-c1:
			// Handle message received from RabbitMQ
			fmt.Printf("Received from RabbitMQ: ID %d - %s\n", msg1.id, msg1.msg)

		case msg2 := <-c2:
			// Handle message received from Kafka
			fmt.Printf("Received from Kafka: ID %d - %s\n", msg2.id, msg2.msg)

		case <-time.After(time.Second * 4):
			// Handle timeout if no messages are received within 4 seconds
			println("timeout")

			// Default case can be used, but it would break the code in this scenario
			// default:
			// 	println("default")
		}
	}
}
