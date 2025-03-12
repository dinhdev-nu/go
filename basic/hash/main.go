package main

import (
	"fmt"
	"time"
)

type message struct {
	Id      int
	Content string
}

func main() {

	messageChannel := make(chan message) // creat a channel

	orders := []message{
		{Id: 1, Content: "Hello"},
		{Id: 2, Content: "World"},
		{Id: 3, Content: "From"},
		{Id: 4, Content: "Go"},
	}

	// Sender goroutine
	go func() {
		defer close(messageChannel)
		for _, order := range orders {
			fmt.Println("Sending:::  ", order.Id, " Content ", order.Content)
			messageChannel <- order
			time.Sleep(time.Second * 1)
		}
	}()

	// Receiver goroutine
	go func() {
		for order := range messageChannel {
			fmt.Println("Message:::  ", order.Id, " Content ", order.Content)
		}
	}() 

	// Wait for receiver to finish processing

	time.Sleep(time.Second * 5)
	fmt.Println("Done")
}
