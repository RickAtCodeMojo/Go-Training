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

func main() {
	mux := fanIn(work("Aether"), work("Hortons"))
	for i := 0; i < 10; i++ {
		fmt.Printf("%s\n", <-mux)
	}
}

func work(team string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			time.Sleep(time.Duration(upTo(1000)) * time.Millisecond)
			ch <- fmt.Sprintf("%s is working on Jira ticket %d)", team, i)
		}
	}()
	return ch
}

func fanIn(queue1, queue2 <-chan string) <-chan string {
	ch := make(chan string)
	go func() {
		for {
			ch <- <-queue1
		}
	}()
	go func() {
		for {
			ch <- <-queue2
		}
	}()
	return ch
}
