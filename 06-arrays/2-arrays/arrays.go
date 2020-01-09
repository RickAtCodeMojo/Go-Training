package main

import (
	"fmt"
)

func main() {
	//define an array with unknown number of elements
	var movies = [...]string{
		"Avatar",
		"Titanic",
		"Star Wars: The Force Awakens",
		"Jurassic World",
		"Marvel's The Avengers",
		"Furious 7",
		"Avengers: Age of Ultron",
		"Harry Potter and the Deathly Hallows Part 2",
		"Star Wars: The Last Jedi",
		"Black Panther",
		"Frozen",
		"Beauty and the Beast (2017)",
		"The Fate of the Furious",
		"Iron Man 3",
		"Minions",
	}
	fmt.Printf("Top %d Grossing Movies:\n", len(movies))
	fmt.Println(movies)
	fmt.Println("length:", len(movies), "capacity", cap(movies))

	// Use this syntax to declare and initialize an array
	// in one line.
	flags := [7]bool{false, false, false, false, false, false, false}
	fmt.Println("flags:", flags)

	// ints := [len(flags)]int

}
