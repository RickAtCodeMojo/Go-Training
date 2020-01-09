package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	duration := 50 * time.Millisecond //set a limit to how long this task can take to complete

	ctx, cancel := context.WithTimeout(context.Background(), duration) //return a context and a cancellation function
	defer cancel()                                                     //defer calling cancel until the end

	ch := make(chan string, 1) //using context relies on a channel with a buffer of 1

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) //this won't get done in time
		ch <- "paper"                                                //put data in the channel
	}()

	select {
	case p := <-ch: //read the channel
		fmt.Println("work complete", p) //say the work is done

	case <-ctx.Done(): //if the context is recieved first
		fmt.Println("moving on") //say that we are complete
	}
	//...cancel()
}
