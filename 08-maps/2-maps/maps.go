package main

import "fmt"

func main() {

	//use make to create a map
	drivers := make(map[string]int)
	drivers["Lewis Hamilton"] = 363
	drivers["Sebastian Vettel"] = 317
	drivers["Valtteri Bottas"] = 305
	drivers["Kimi Räikkönen"] = 205
	drivers["Daniel Ricciardo"] = 200
	drivers["Max Verstappen	"] = 168
	drivers["Sergio Perez"] = 100
	drivers["Esteban Ocon"] = 87
	drivers["Carlos Sainz"] = 54
	drivers["Nico Hulkenberg"] = 43
	drivers["Felipe Massa"] = 43
	drivers["Lance Stroll"] = 40
	drivers["Romain Grosjean"] = 28
	drivers["Kevin Magnussen"] = 19
	drivers["Fernando Alonso"] = 17
	drivers["Stoffel Vandoorne"] = 13
	drivers["Jolyon Palmer"] = 8
	drivers["Pascal Wehrlein"] = 5
	drivers["Daniil Kvyat"] = 5
	drivers["Jenson Button"] = 0
	drivers["Marcus Ericsson"] = 0
	drivers["Pierre Gasly"] = 0
	drivers["Antonio Giovinazzi"] = 0
	drivers["Brendon Hartley"] = 0

	//print an entire map
	fmt.Println("Drivers ............")
	//print keys and values
	print(drivers) // the same line with this syntax.

	fmt.Println("\nTeams ............")
	teams := map[string]int{"Mercedes": 1, "Ferrari": 2, "Red Bull": 3, "Force India": 4, "Renault": 5, "Williams": 6, "Haas": 7}
	print(teams)
}

func print(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}
}
