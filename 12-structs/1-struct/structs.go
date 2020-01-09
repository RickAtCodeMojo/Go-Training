package main

import (
	"fmt"
)

type movie struct {
	rank     int
	title    string
	studio   string
	gross    float64
	domestic float64
	year     int
	// worldwide float64
}

func main() {
	avatar := movie{rank: 1, title: "Avatar", studio: "Fox", gross: 2788.0, domestic: 760.5, year: 2009}
	fmt.Println(avatar)

	var titanic movie
	titanic = movie{
		rank:     2,
		title:    "Titanic",
		studio:   "Par.",
		gross:    2187.5,
		domestic: 659.4,
		year:     1997,
	}

	other := movie{2, "Titanic", "Par.", 2187.5, 659.4, 1997}
	fmt.Println(other.title)

	pTitanic := &titanic //create a pointer to the struct
	fmt.Println(pTitanic)

	fmt.Println(pTitanic.gross) //you still get to use dot notation with pointers when dereferencing struct items
	pTitanic.rank += 10

	fmt.Println(titanic.rank)

}
