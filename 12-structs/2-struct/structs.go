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
}

func main() {
	var movies []movie
	movies = append(movies, movie{rank: 1, title: "Avatar", studio: "Fox", gross: 2788.0, domestic: 760.5, year: 2009})
	movies = append(movies, movie{rank: 2, title: "Titanic", studio: "Par.", gross: 2187.5, domestic: 659.4, year: 1997})
	movies = append(movies, movie{rank: 3, title: "Star Wars: The Force Awakens", studio: "BV", gross: 2068.2, domestic: 936.7, year: 2015})
	movies = append(movies, movie{rank: 4, title: "Jurassic World", studio: "Uni.", gross: 1671.7, domestic: 652.3, year: 2015})
	movies = append(movies, movie{rank: 5, title: "Marvel's The Avengers", studio: "BV", gross: 1518.8, domestic: 623.4, year: 2012})
	movies = append(movies, movie{rank: 6, title: "Furious 7", studio: "Uni.", gross: 1516.0, domestic: 353.0, year: 2015})
	movies = append(movies, movie{rank: 7, title: "Avengers: Age of Ultron", studio: "BV", gross: 1405.4, domestic: 459.0, year: 2015})
	movies = append(movies, movie{rank: 8, title: "Harry Potter and the Deathly Hallows Part 2", studio: "WB", gross: 1341.5, domestic: 381.0, year: 2011})
	movies = append(movies, movie{rank: 9, title: "Star Wars: The Last Jedi", studio: "BV", gross: 1332.8, domestic: 620.2, year: 2017})
	movies = append(movies, movie{rank: 10, title: "Black Panther", studio: "BV", gross: 1331.3, domestic: 688.0, year: 2018})
	movies = append(movies, movie{rank: 11, title: "Frozen", studio: "BV", gross: 1276.5, domestic: 400.7, year: 2013})
	movies = append(movies, movie{rank: 12, title: "Beauty and the Beast (2017)", studio: "BV", gross: 1263.5, domestic: 504.0, year: 2017})
	movies = append(movies, movie{rank: 13, title: "The Fate of the Furious", studio: "Uni.", gross: 1236.0, domestic: 226.0, year: 2017})
	movies = append(movies, movie{rank: 14, title: "Iron Man 3", studio: "BV", gross: 1214.8, domestic: 409.0, year: 2013})
	movies = append(movies, movie{rank: 15, title: "Minions", studio: "Uni.", gross: 1159.4, domestic: 336.0, year: 2015})
	fmt.Println(movies)
	fmt.Println(movies[9].title)

}
