package main

import (
	"fmt"
	"math/rand"
)

func main() {
	jobs := []string{"report", "essay", "assignment", "slide show", "performance review"}

	ch := make(chan string, 1)

	go func() {
		for p := range ch {
			fmt.Println("employee is working :", p) //guaranteed hand off
		}
	}()

	const work = 10
	for w := 0; w < work; w++ {
		ch <- jobs[rand.Intn(len(jobs))]
	}

	close(ch)
}
