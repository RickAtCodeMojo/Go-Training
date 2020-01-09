package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//put items into two channels concurrently
	aether := generate("Aether", "writing code")
	hortons := generate("Hortons", "writing code")

	//pop 5 items off each channel serially
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", <-aether)
		fmt.Printf("%s\n", <-hortons)
	}

}

//generate return a channel that has been populated async.
func generate(job string, item string) <-chan string {
	ch := make(chan string)
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s is %s (for Jira ticket: %d)", job, item, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return ch
}
