package main

import (
	"fmt"
	"sync"
)

func main() {
	const cap = 5
	ch := make(chan string, 5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for p := range ch {
			fmt.Println("employee has received :", p)
		}
	}()

	work := 15
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager recieved work", w)
		default:
			fmt.Println("the work was dropped", w)
		}
	}
	wg.Wait()
	close(ch)
}
