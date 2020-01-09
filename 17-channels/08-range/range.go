package main

import "fmt"

type suits int
type cards int

const (
	Ace cards = iota
	Deuce
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

func (c cards) String() string {
	switch c {
	case Ace:
		return "Ace"
	case Deuce:
		return "Deuce"
	case Three:
		return "Three"
	case Four:
		return "Four"
	case Five:
		return "Five"
	case Seven:
		return "Seven"
	case Eight:
		return "Eight"
	case Nine:
		return "Nine"
	case Ten:
		return "Ten"
	case Jack:
		return "Jack"
	case Six: //<-doesn't matter that this is out of order - bad from though
		return "Six"
	case Queen:
		return "Queen"
	case King:
		return "King"
	default:
		return ""
	}
}
func main() {
	queue := make(chan cards, int(King+1))
	for card := Ace; card <= King; card++ {
		queue <- card
	}
	close(queue) //this merely stops receiving elements on the queue
	//- items passed to the queue are still there

	for elem := range queue {
		fmt.Println(elem)
	}
}
