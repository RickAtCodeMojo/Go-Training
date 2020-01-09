package main

import "fmt"

//Movie stores information about a movie
type Movie struct {
	Rank     int
	Title    string
	Studio   string
	Gross    float64
	Domestic float64
	Year     int
}

//Movies is the list of movies
var Movies []Movie

func main() {
	fillMovies()
	for _, m := range Movies {
		fmt.Printf("%s made %.1f%% of its revenue in the US\n", m.Title, m.PercentDomestic()*100)
	}

	fmt.Println()
	top := Movies[0]
	fmt.Printf("%s made $%.1f revenue outside the US\n", top.Title, top.Worldwide())

	pTop := &top
	fmt.Printf("%s made %.1f%% of its revenue outside the US\n", pTop.Title, pTop.PercentWorldwide()*100)

	fmt.Println()
	fmt.Println("Movies")
	fmt.Println(Movies)

}

//Worldwide - returns the world wide sales - outside the US
func (m Movie) Worldwide() float64 {
	return m.Gross - m.Domestic
}

//PercentDomestic - returns the the percentage of US sales to all sales
func (m Movie) PercentDomestic() float64 {
	return m.Domestic / m.Gross
}

//PercentWorldwide - returns the the percentage of world sales to all sales
func (m Movie) PercentWorldwide() float64 {
	return m.Worldwide() / m.Gross
}

//String - the good old string formatting method
func (m Movie) String() string {
	return fmt.Sprintf("%d) %s, %s grossed:%.2f in total and %.2f domestically in %d, \n", m.Rank, m.Title, m.Studio, m.Gross, m.Domestic, m.Year)
}

func fillMovies() {
	Movies = append(Movies, Movie{Rank: 1, Title: "Avatar", Studio: "Fox", Gross: 2788.0, Domestic: 760.5, Year: 2009})
	Movies = append(Movies, Movie{Rank: 2, Title: "Titanic", Studio: "Par.", Gross: 2187.5, Domestic: 659.4, Year: 1997})
	Movies = append(Movies, Movie{Rank: 3, Title: "Star Wars: The Force Awakens", Studio: "BV", Gross: 2068.2, Domestic: 936.7, Year: 2015})
	Movies = append(Movies, Movie{Rank: 4, Title: "Jurassic World", Studio: "Uni.", Gross: 1671.7, Domestic: 652.3, Year: 2015})
	Movies = append(Movies, Movie{Rank: 5, Title: "Marvel's The Avengers", Studio: "BV", Gross: 1518.8, Domestic: 623.4, Year: 2012})
	Movies = append(Movies, Movie{Rank: 6, Title: "Furious 7", Studio: "Uni.", Gross: 1516.0, Domestic: 353.0, Year: 2015})
	Movies = append(Movies, Movie{Rank: 7, Title: "Avengers: Age of Ultron", Studio: "BV", Gross: 1405.4, Domestic: 459.0, Year: 2015})
	Movies = append(Movies, Movie{Rank: 8, Title: "Harry Potter and the Deathly Hallows Part 2", Studio: "WB", Gross: 1341.5, Domestic: 381.0, Year: 2011})
	Movies = append(Movies, Movie{Rank: 9, Title: "Star Wars: The Last Jedi", Studio: "BV", Gross: 1332.8, Domestic: 620.2, Year: 2017})
	Movies = append(Movies, Movie{Rank: 10, Title: "Black Panther", Studio: "BV", Gross: 1331.3, Domestic: 688.0, Year: 2018})
	Movies = append(Movies, Movie{Rank: 11, Title: "Frozen", Studio: "BV", Gross: 1276.5, Domestic: 400.7, Year: 2013})
	Movies = append(Movies, Movie{Rank: 12, Title: "Beauty and the Beast (2017)", Studio: "BV", Gross: 1263.5, Domestic: 504.0, Year: 2017})
	Movies = append(Movies, Movie{Rank: 13, Title: "The Fate of the Furious", Studio: "Uni.", Gross: 1236.0, Domestic: 226.0, Year: 2017})
	Movies = append(Movies, Movie{Rank: 14, Title: "Iron Man 3", Studio: "BV", Gross: 1214.8, Domestic: 409.0, Year: 2013})
	Movies = append(Movies, Movie{Rank: 15, Title: "Minions", Studio: "Uni.", Gross: 1159.4, Domestic: 336.0, Year: 2015})

}
