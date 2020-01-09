package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	bids := sellAds("Old Navy", "Canadian Tire", "Ford Canada")
	elapsed := time.Since(start)
	fmt.Println(elapsed)
	for _, b := range bids {
		fmt.Printf("%s bid %.2f\n", b.Advertiser, b.Amount)
	}

}

//first approach
func sellAds(buyers ...string) (results []Bid) {
	for i := 0; i < len(buyers); i++ {
		buy := getBid(buyers[i])
		bid := buy(buyers[i])
		results = append(results, bid)
	}
	return results
}

// second approach
// func sellAds(buyers ...string) (results []Bid) {
// 	bids := make(chan Bid)
// 	for _, b := range buyers {
// 		buy := getBid(b)
// 		go func(b string) { bids <- buy(b) }(b)
// 	}
// 	for i := 0; i < len(buyers); i++ {
// 		bid := <-bids
// 		results = append(results, bid)
// 	}
// 	return results
// }

// timeout
// func sellAds(buyers ...string) (results []Bid) {
// 	bids := make(chan Bid)
// 	for _, b := range buyers {
// 		buy := getBid(b)
// 		go func(b string) { bids <- buy(b) }(b)
// 	}
// 	timeout := time.After(80 * time.Millisecond)
// 	for i := 0; i < len(buyers); i++ {
// 		select {
// 		case bid := <-bids:
// 			results = append(results, bid)
// 		case <-timeout:
// 			fmt.Println("Ran out of time for auction")

// 		}
// 	}
// 	return results
// }

var seed = rand.NewSource(time.Now().UnixNano())

//UnsignedInt generates a random unsigned int up to max value
func upTo(max int) int {
	r := rand.New(seed)
	n := r.Intn(int(max))
	return n
}
func randomBid() float64 {
	a := upTo(20)
	b := upTo(99)
	bid := float64(a) + float64(b)/100
	return bid
}

//Bid stores the result of requesting a bid
type Bid struct {
	Advertiser string
	Amount     float64
}

//Auction is a type of function that gets a bid for an adspace
type Auction func(buyer string) Bid

func getBid(advertiser string) Auction {
	return func(buyer string) Bid {
		time.Sleep(time.Duration(upTo(100)) * time.Millisecond)
		return Bid{Advertiser: advertiser, Amount: randomBid()}
	}
}
