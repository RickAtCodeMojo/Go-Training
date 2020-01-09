package main

import (
	"fmt"
	"sort"
)

type movie struct {
	Rank      int
	Title     string
	Studio    string
	Worldwide float64
	Year      int
}

var movies map[string]movie

func main() {
	fill()

	fmt.Println("Movies ............")
	print(movies)

	// fmt.Println("Sorted Movies ............")
	// for _, k := range sortedKeys(movies) {
	// 	fmt.Println(k, movies[k].Rank)
	// }

}

func print(m map[string]movie) {
	for k, v := range m {
		fmt.Println(k, v.Rank)
	}
}

type sortedMap struct {
	m map[string]movie
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]].Rank < sm.m[sm.s[j]].Rank
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]movie) []string {
	sm := new(sortedMap)
	sm.m = m
	sm.s = make([]string, len(m))
	i := 0
	for key := range m {
		sm.s[i] = key
		i++
	}
	sort.Sort(sm)
	return sm.s
}

func fill() {

	movies = make(map[string]movie)

	movies["Avatar"] = movie{Rank: 1, Title: "Avatar", Studio: "Fox", Worldwide: 2788, Year: 2009}
	movies["Titanic"] = movie{Rank: 2, Title: "Titanic", Studio: "Par.", Worldwide: 2187.5, Year: 1997}
	movies["Star Wars: The Force Awakens"] = movie{Rank: 3, Title: "Star Wars: The Force Awakens", Studio: "BV", Worldwide: 2068.2, Year: 2015}
	movies["Jurassic World"] = movie{Rank: 4, Title: "Jurassic World", Studio: "Uni.", Worldwide: 1671.7, Year: 2015}
	movies["Marvel's The Avengers"] = movie{Rank: 5, Title: "Marvel's The Avengers", Studio: "BV", Worldwide: 1518.8, Year: 2012}
	movies["Furious 7"] = movie{Rank: 6, Title: "Furious 7", Studio: "Uni.", Worldwide: 1516, Year: 2015}
	movies["Avengers: Age of Ultron"] = movie{Rank: 7, Title: "Avengers: Age of Ultron", Studio: "BV", Worldwide: 1405.4, Year: 2015}
	movies["Harry Potter and the Deathly Hallows Part 2"] = movie{Rank: 8, Title: "Harry Potter and the Deathly Hallows Part 2", Studio: "WB", Worldwide: 1341.5, Year: 2011}
	movies["Star Wars: The Last Jedi"] = movie{Rank: 9, Title: "Star Wars: The Last Jedi", Studio: "BV", Worldwide: 1332.8, Year: 2017}
	movies["Black Panther"] = movie{Rank: 10, Title: "Black Panther", Studio: "BV", Worldwide: 1331.3, Year: 2018}
	movies["Frozen"] = movie{Rank: 11, Title: "Frozen", Studio: "BV", Worldwide: 1276.5, Year: 2013}
	movies["Beauty and the Beast (2017)"] = movie{Rank: 12, Title: "Beauty and the Beast (2017)", Studio: "BV", Worldwide: 1263.5, Year: 2017}
	movies["The Fate of the Furious"] = movie{Rank: 13, Title: "The Fate of the Furious", Studio: "Uni.", Worldwide: 1236, Year: 2017}
	movies["Iron Man 3"] = movie{Rank: 14, Title: "Iron Man 3", Studio: "BV", Worldwide: 1214.8, Year: 2013}
	movies["Minions"] = movie{Rank: 15, Title: "Minions", Studio: "Uni.", Worldwide: 1159.4, Year: 2015}
	movies["Captain America: Civil War"] = movie{Rank: 16, Title: "Captain America: Civil War", Studio: "BV", Worldwide: 1153.3, Year: 2016}
	movies["Transformers: Dark of the Moon"] = movie{Rank: 17, Title: "Transformers: Dark of the Moon", Studio: "P/DW", Worldwide: 1123.8, Year: 2011}
	movies["The Lord of the Rings: The Return of the King"] = movie{Rank: 18, Title: "The Lord of the Rings: The Return of the King", Studio: "NL", Worldwide: 1119.9, Year: 2003}
	movies["Skyfall"] = movie{Rank: 19, Title: "Skyfall", Studio: "Sony", Worldwide: 1108.6, Year: 2012}
	movies["Transformers: Age of Extinction"] = movie{Rank: 20, Title: "Transformers: Age of Extinction", Studio: "Par.", Worldwide: 1104.1, Year: 2014}
	movies["The Dark Knight Rises"] = movie{Rank: 21, Title: "The Dark Knight Rises", Studio: "WB", Worldwide: 1084.9, Year: 2012}
	movies["Toy Story 3"] = movie{Rank: 22, Title: "Toy Story 3", Studio: "BV", Worldwide: 1067, Year: 2010}
	movies["Pirates of the Caribbean: Dead Man's Chest"] = movie{Rank: 23, Title: "Pirates of the Caribbean: Dead Man's Chest", Studio: "BV", Worldwide: 1066.2, Year: 2006}
	movies["Rogue One: A Star Wars Story"] = movie{Rank: 24, Title: "Rogue One: A Star Wars Story", Studio: "BV", Worldwide: 1056.1, Year: 2016}
	movies["Pirates of the Caribbean: On Stranger Tides"] = movie{Rank: 25, Title: "Pirates of the Caribbean: On Stranger Tides", Studio: "BV", Worldwide: 1045.7, Year: 2011}
	movies["Despicable Me 3"] = movie{Rank: 26, Title: "Despicable Me 3", Studio: "Uni.", Worldwide: 1034.8, Year: 2017}
	movies["Jurassic Park"] = movie{Rank: 27, Title: "Jurassic Park", Studio: "Uni.", Worldwide: 1029.2, Year: 1993}
	movies["Finding Dory"] = movie{Rank: 28, Title: "Finding Dory", Studio: "BV", Worldwide: 1028.6, Year: 2016}
	movies["Star Wars: Episode I - The Phantom Menace"] = movie{Rank: 29, Title: "Star Wars: Episode I - The Phantom Menace", Studio: "Fox", Worldwide: 1027, Year: 1999}
	movies["Alice in Wonderland (2010)"] = movie{Rank: 30, Title: "Alice in Wonderland (2010)", Studio: "BV", Worldwide: 1025.5, Year: 2010}
	movies["Zootopia"] = movie{Rank: 31, Title: "Zootopia", Studio: "BV", Worldwide: 1023.8, Year: 2016}
	movies["The Hobbit: An Unexpected Journey"] = movie{Rank: 32, Title: "The Hobbit: An Unexpected Journey", Studio: "WB (NL)", Worldwide: 1021.1, Year: 2012}
	movies["The Dark Knight"] = movie{Rank: 33, Title: "The Dark Knight", Studio: "WB", Worldwide: 1004.6, Year: 2008}
	movies["Harry Potter and the Sorcerer's Stone"] = movie{Rank: 34, Title: "Harry Potter and the Sorcerer's Stone", Studio: "WB", Worldwide: 974.8, Year: 2001}
	movies["Despicable Me 2"] = movie{Rank: 35, Title: "Despicable Me 2", Studio: "Uni.", Worldwide: 970.8, Year: 2013}
	movies["The Lion King"] = movie{Rank: 36, Title: "The Lion King", Studio: "BV", Worldwide: 968.5, Year: 1994}
	movies["The Jungle Book (2016)"] = movie{Rank: 37, Title: "The Jungle Book (2016)", Studio: "BV", Worldwide: 966.6, Year: 2016}
	movies["Pirates of the Caribbean: At World's End"] = movie{Rank: 38, Title: "Pirates of the Caribbean: At World's End", Studio: "BV", Worldwide: 963.4, Year: 2007}
	movies["Harry Potter and the Deathly Hallows Part 1"] = movie{Rank: 39, Title: "Harry Potter and the Deathly Hallows Part 1", Studio: "WB", Worldwide: 960.3, Year: 2010}
	movies["The Hobbit: The Desolation of Smaug"] = movie{Rank: 40, Title: "The Hobbit: The Desolation of Smaug", Studio: "WB (NL)", Worldwide: 958.4, Year: 2013}
	movies["Jumanji: Welcome to the Jungle"] = movie{Rank: 41, Title: "Jumanji: Welcome to the Jungle", Studio: "Sony", Worldwide: 956.7, Year: 2017}
	movies["The Hobbit: The Battle of the Five Armies"] = movie{Rank: 42, Title: "The Hobbit: The Battle of the Five Armies", Studio: "WB (NL)", Worldwide: 956, Year: 2014}
	movies["Finding Nemo"] = movie{Rank: 43, Title: "Finding Nemo", Studio: "BV", Worldwide: 940.3, Year: 2003}
	movies["Harry Potter and the Order of the Phoenix"] = movie{Rank: 44, Title: "Harry Potter and the Order of the Phoenix", Studio: "WB", Worldwide: 939.9, Year: 2007}
	movies["Harry Potter and the Half-Blood Prince"] = movie{Rank: 45, Title: "Harry Potter and the Half-Blood Prince", Studio: "WB", Worldwide: 934.4, Year: 2009}
	movies["The Lord of the Rings: The Two Towers"] = movie{Rank: 46, Title: "The Lord of the Rings: The Two Towers", Studio: "NL", Worldwide: 926, Year: 2002}
	movies["Shrek 2"] = movie{Rank: 47, Title: "Shrek 2", Studio: "DW", Worldwide: 919.8, Year: 2004}
	movies["Harry Potter and the Goblet of Fire"] = movie{Rank: 48, Title: "Harry Potter and the Goblet of Fire", Studio: "WB", Worldwide: 896.9, Year: 2005}
	movies["Spider-Man 3"] = movie{Rank: 49, Title: "Spider-Man 3", Studio: "Sony", Worldwide: 890.9, Year: 2007}
	movies["Ice Age: Dawn of the Dinosaurs"] = movie{Rank: 50, Title: "Ice Age: Dawn of the Dinosaurs", Studio: "Fox", Worldwide: 886.7, Year: 2009}

}
