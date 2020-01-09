package main

import "fmt"

var drivers map[string]int

func main() {

	//use make to create a map
	drivers = make(map[string]int)
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
	fmt.Println("drivers:", drivers)

}
