// _Slices_ are a key data type in Go, giving a more
// powerful interface to sequences than arrays.

package main

import (
	"fmt"
	"sort"
)

//sorting slices of structs

//Team Struct tracks teams
type team struct {
	name     string
	games    uint
	wins     uint
	losses   uint
	otlosses uint
}

func main() {

	bids := []struct {
		bidder string
		amount float64
		height int
		width  int
	}{ //we can declare and initalize a slice of structs
		{bidder: "Old Navy", amount: 13.99, height: 100, width: 50},
		{bidder: "Orvis", amount: 14.99, height: 100, width: 50},
		{bidder: "Microsoft", amount: 11.99, height: 25, width: 500},
		{bidder: "Apple", amount: 12.59, height: 25, width: 500}, //<-- note final comma
	}
	fmt.Println("Bids:", bids)

	var teams []team
	teams = append(teams, team{name: "Montréal", games: 68, wins: 25, losses: 31, otlosses: 12})
	teams = append(teams, team{name: "Ottawa", games: 67, wins: 23, losses: 33, otlosses: 11})
	teams = append(teams, team{name: "Winnipeg", games: 68, wins: 41, losses: 18, otlosses: 9})
	fmt.Println("Teams:", teams)

	// Sorting a slice of structs - requires defining  function
	sort.Slice(bids, func(i, j int) bool {
		return bids[i].bidder < bids[j].bidder
	})
	fmt.Println("Bids sorted by bidder name:", bids)

	//of course you can have the function predefined instead of in place
	byAmount := func(i, j int) bool {
		return bids[i].amount < bids[j].amount
	}
	sort.Slice(bids, byAmount) //which looks cleaner - though you don't see what it does
	fmt.Println("Bids sorted by amount:", bids)
}
