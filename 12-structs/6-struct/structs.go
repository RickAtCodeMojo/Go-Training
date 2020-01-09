//keys and values to maps can be structs
package main

import "fmt"

type slot struct {
	position int
	team     string
}
type driver struct {
	points  int
	name    string
	country string
}

func main() {
	drivers := make(map[slot]driver)
	//note that there can be no ambiguity when you use literal structs
	drivers[slot{position: 1, team: "MERCEDES"}] = driver{points: 363, name: "Lewis Hamilton", country: "GBR"}
	drivers[slot{position: 1, team: "FERRARI"}] = driver{points: 317, name: "Sebastian Vettel", country: "GER"}
	drivers[slot{position: 2, team: "MERCEDES"}] = driver{points: 305, name: "Valtteri Bottas", country: "FIN"}
	drivers[slot{position: 2, team: "FERRARI"}] = driver{points: 205, name: "Kimi Räikkönen", country: "FIN"}
	drivers[slot{position: 1, team: "RED BULL RACING"}] = driver{points: 200, name: "Daniel Ricciardo", country: "AUS"}
	drivers[slot{position: 2, team: "RED BULL RACING"}] = driver{points: 168, name: "Max Verstappen", country: "NED"}
	drivers[slot{position: 1, team: "FORCE INDIA MERCEDES"}] = driver{points: 100, name: "Sergio Perez", country: "MEX"}
	drivers[slot{position: 1, team: "FORCE INDIA MERCEDES"}] = driver{points: 87, name: "Esteban Ocon", country: "FRA"}
	drivers[slot{position: 1, team: "RENAULT"}] = driver{points: 54, name: "Carlos Sainz", country: "ESP"}
	drivers[slot{position: 2, team: "RENAULT"}] = driver{points: 43, name: "Nico Hulkenberg", country: "GER"}
	drivers[slot{position: 1, team: "WILLIAMS MERCEDES"}] = driver{points: 43, name: "Felipe Massa", country: "BRA"}
	drivers[slot{position: 2, team: "WILLIAMS MERCEDES"}] = driver{points: 40, name: "Lance Stroll", country: "CAN"}
	drivers[slot{position: 1, team: "HAAS FERRARI"}] = driver{points: 28, name: "Romain Grosjean", country: "FRA"}
	drivers[slot{position: 2, team: "HAAS FERRARI"}] = driver{points: 19, name: "Kevin Magnussen	", country: "DEN"}
	drivers[slot{position: 1, team: "MCLAREN HONDA"}] = driver{points: 17, name: "Fernando Alonso", country: "ESP"}
	drivers[slot{position: 2, team: "MCLAREN HONDA"}] = driver{points: 13, name: "Stoffel Vandoorne", country: "BEL"}
	drivers[slot{position: 3, team: "RENAULT"}] = driver{points: 8, name: "Jolyon Palmer", country: "GBR"}
	drivers[slot{position: 1, team: "SAUBER FERRARI"}] = driver{points: 5, name: "Pascal Wehrlein", country: "GER"}
	drivers[slot{position: 1, team: "TORO ROSSO"}] = driver{points: 5, name: "Daniil Kvyat", country: "RUS"}
	drivers[slot{position: 2, team: "SAUBER FERRARI"}] = driver{points: 0, name: "Marcus Ericsson	", country: "SWE"}
	drivers[slot{position: 2, team: "TORO ROSSO"}] = driver{points: 0, name: "Pierre Gasly", country: "FRA"}
	drivers[slot{position: 1, team: "SAUBER FERRARI"}] = driver{points: 0, name: "Antonio Giovinazzi", country: "ITA"}
	drivers[slot{position: 3, team: "TORO ROSSO"}] = driver{points: 0, name: "Brendon Hartley", country: "NZL"}
	fmt.Println(drivers[slot{position: 2, team: "RENAULT"}])
	delete(drivers, slot{position: 3, team: "RENAULT"})
	delete(drivers, slot{position: 3, team: "TORO ROSSO"})
	fmt.Println(driversFrom(drivers, "FRA"))
	//Stupid Gopher tricks
	fmt.Println("Stupid Gopher tricks")
	m := map[cons]string{}
	c := cons{}
	for _, s := range []string{"life?", "with my", "I doing", "What am"} {
		c = cons{s, c}
	}
	m[c] = "No idea."
	fmt.Println(c, m[c])
}
func driversFrom(drivers map[slot]driver, country string) []string {
	var names []string
	for _, v := range drivers {
		if v.country == country {
			names = append(names, v.name)
		}
	}
	return names
}

//Stupid Gopher Tricks..
type cons struct {
	car string
	cdr interface{}
}

func (c cons) String() string {
	if c.cdr == nil || c.cdr == (cons{}) {
		return c.car
	}
	return fmt.Sprintf("%v %v", c.car, c.cdr)
}
