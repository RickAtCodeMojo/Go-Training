package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)
	//FAN OUT
	jobs := []string{"report", "essay", "assignment", "slide show", "performance review"}
	for _, job := range jobs {
		go func(job string) {
			fmt.Println("Working on your", job)
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch <- job
		}(job)
	}
	for i := range jobs {
		fmt.Println(i, "Completed your", <-ch)
	}
}
