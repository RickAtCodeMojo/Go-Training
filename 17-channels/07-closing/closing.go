package main

import "fmt"

func main() {

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, ok := <-jobs //read the job channel
			if ok {         //if we recieved something on the channel
				fmt.Println("recieved job", j) //print it
			} else {
				fmt.Println("all jobs received") //otherwise we've run out of stuff
				done <- true                     //send the signal that we're done
				return                           //leave the func
			}
		}
	}()

	for j := 1; j <= 3; j++ { //send 3 ints to the job channel
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done //block until done received
}
