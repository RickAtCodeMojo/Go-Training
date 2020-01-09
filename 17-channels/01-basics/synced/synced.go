package main

import (
	"fmt"
	"math/rand"
	"time"
)

var seed = rand.NewSource(time.Now().UnixNano())

//UnsignedInt generates a random unsigned int up to max value
func upTo(max int) int {
	r := rand.New(seed)
	n := r.Intn(int(max))
	return n
}

type message struct {
	text string
	done chan bool
}

func main() {
	mux := fanIn(work("Aether"), work("Hortons")) //crank up two concurrent processes that pump items into a channel
	for i := 0; i < 5; i++ {
		msg1 := <-mux
		fmt.Println(msg1.text)
		msg2 := <-mux
		fmt.Println(msg2.text)

		msg1.done <- true //signal that we're done
		msg2.done <- true
	}
}

//work does stuff and puts the message into a channel and returns the channel
func work(team string) <-chan message {
	ch := make(chan message)
	done := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			ch <- message{text: fmt.Sprintf("%s is working on Jira ticket %d)", team, i), done: done}
			time.Sleep(time.Duration(upTo(1000)) * time.Millisecond)
			<-done
		}
	}()
	return ch
}

//takes elements from two channels and fans them back into one
func fanIn(queue1, queue2 <-chan message) <-chan message {
	ch := make(chan message)
	go func() {
		for { //endlessly process the first queue in
			ch <- <-queue1
		}
	}()
	go func() {
		for { //endlessly process the second queue in
			ch <- <-queue2
		}
	}()
	return ch
}
