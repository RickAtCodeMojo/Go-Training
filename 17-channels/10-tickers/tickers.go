package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var firstPart sync.WaitGroup
	firstPart.Add(1)
	timeChan := time.NewTimer(time.Second).C

	tickChan := time.NewTicker(time.Millisecond * 400).C

	doneChan := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		doneChan <- true
	}()

	go func() {
		for {
			select {
			case <-timeChan:
				fmt.Println("Timer expired")
			case <-tickChan:
				fmt.Println("Ticker ticked")
			case <-doneChan:
				fmt.Println("Done")
				firstPart.Done()
				return
			}
		}
	}()
	firstPart.Wait()

	ticker := time.NewTicker(500 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()
	//Tickers can be stopped like timers.
	//Once a ticker is stopped it won’t receive any more values on its channel. We’ll stop ours after 1600ms.

	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	fmt.Println("Ticker stopped")

}
