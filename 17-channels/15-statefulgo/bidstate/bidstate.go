package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//In this example our ads will be owned by a single goroutine.
//This will guarantee that the data is never corrupted with concurrent access.
//In order to read or write ads, other goroutines will send messages to
//the owning goroutine and receive corresponding replies.
//These checkBid and bid structs encapsulate those requests and a way for the owning goroutine to respond.
type checkBid struct {
	key  int
	resp chan float64
}
type bid struct {
	key    int
	bidder int
	val    float64
	resp   chan bool
}

func main() {

	//As before we’ll count how many operations we perform.
	var checkBids uint64
	var placedBids uint64

	//The reads and bids channels will be used by other goroutines to issue read and write requests, respectively.
	reads := make(chan *checkBid)
	bids := make(chan *bid)

	//Here is the goroutine that owns the ads, which is a map as in the
	//previous example but now private to the goroutine.
	//This goroutine repeatedly selects on the reads and bids channels,
	//responding to requests as they arrive. A response is executed by first
	//performing the requested operation and then sending a value on the response
	//channel resp to indicate success (and the desired value in the case of reads).
	var ads = make(map[int]float64)
	var guard sync.Mutex
	go func() {
		for {
			select {
			case read := <-reads:
				read.resp <- ads[read.key]
			case bid := <-bids:
				guard.Lock()                //why do we need a mutex here??
				defer guard.Unlock()        //what would happen if we took it out
				if bid.val > ads[bid.key] { //if this is a higher bid
					ads[bid.key] = bid.val //keep it
				}
				bid.resp <- true
			}
		}
	}()

	//This starts 1000 goroutines to issue reads to the ads-owning goroutine via the reads channel.
	//Each read requires constructing a checkBid, sending it over the reads channel,
	//and the receiving the result over the provided resp channel.
	for r := 0; r < 1000; r++ {
		go func() {
			for {
				read := &checkBid{
					key:  rand.Intn(5),
					resp: make(chan float64)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&checkBids, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//We start 100 bids as well, using a similar approach.
	for w := 0; w < 100; w++ {
		go func() {
			for {
				placedBid := &bid{
					key:  rand.Intn(5),
					val:  float64(rand.Intn(100)) + float64(rand.Intn(100))/100.0,
					resp: make(chan bool)}
				bids <- placedBid
				<-placedBid.resp
				atomic.AddUint64(&placedBids, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	//Let the goroutines work for a second.
	time.Sleep(150 * time.Millisecond)
	//Finally, capture and report the op counts.
	checkBidsFinal := atomic.LoadUint64(&checkBids)
	fmt.Println("checkBids:", checkBidsFinal)
	placedBidsFinal := atomic.LoadUint64(&placedBids)
	fmt.Println("Placed Bids:", placedBidsFinal)
	fmt.Println(ads)
}
