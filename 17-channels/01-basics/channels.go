package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var ch chan string
	ch = make(chan string)
	// ch := make(chan string)
	go work("commit code", ch)
	for i := 0; i < 7; i++ {
		fmt.Printf("You will %s\n", <-ch)
	}

	// jobs := generate("regression test", "commit")
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You will %s\n", <-jobs)
	// }
	// releases := generate("release", "package")
	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("You will %s\n", <-releases)
	// }
}

func work(job string, ch chan string) {
	for i := 0; ; i++ {
		ch <- fmt.Sprintf("%s. (backlog item %d)", job, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}

func generate(job string, item string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s. (%s #: %d)", job, item, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
