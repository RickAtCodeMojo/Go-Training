package main

import "fmt"

var drivers map[string]int

func main() {

	fill()

	// Get a value for a key with map[key].
	fmt.Println("Alonso's points: ", drivers["Fernando Alonso"])
	key := "Felipe Massa"
	fmt.Printf("Massa's points: %d\n", drivers[key])

	//use len on a map
	fmt.Println("Number of drivers:", len(drivers))
	// delete an element with the key
	delete(drivers, "Jenson Button")
	//use len on a map
	fmt.Println("Number of drivers Now:", len(drivers))

	//the optional 2nd return value indicates whether the value exists in the map
	key = "Jenson Button"
	_, ok := drivers[key]
	if ok {
		fmt.Printf("%s is indeed a driver.\n", key)
	} else {
		fmt.Printf("%s is not driving this year.\n", key)
	}

}

func print(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}

}
func fill() {
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

}
