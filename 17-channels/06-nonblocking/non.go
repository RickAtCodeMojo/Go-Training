package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool)

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default: //<-if no message is received default will happen immediately
		fmt.Println("no message received")
	}

	msg := "hello world"
	select {
	case messages <- msg:
		fmt.Println("sent message:", msg)
	default: //<-- even though we have a msg, since it is not immediately sent control falls to default
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default: //finally control falls to this default and no activity prints
		fmt.Println("no activity")
	}
}
