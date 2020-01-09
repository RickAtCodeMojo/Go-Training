package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
	}()
	fmt.Println("Completed your", <-ch)
}
