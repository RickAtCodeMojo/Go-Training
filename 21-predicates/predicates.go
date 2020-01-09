package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//Slot is the driver's position on the team
type Slot struct {
	Position int
	Team     string
}

//Driver contains information about the driver
type Driver struct {
	Points  int
	Name    string
	Country string
}

//Position returns the first Slot Position if the Driver's Name matches "d"
func Position(values map[Slot]Driver, d string) int {

	for i, v := range values {
		if v.Name == d {
			return i.Position
		}
	}
	return -1
}

// Includes returns `true` if the target Driver is in the map
func Includes(values map[Slot]Driver, d string) bool {
	return Position(values, d) >= 0
}

// Has returns TRUE if a member of the map satisfies the predieate "f"
func Has(values map[Slot]Driver, f func(Driver) bool) bool {
	for _, v := range values {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns TRUE if all of the Drivers in the map
// satisfy the predicate "f".
func All(values map[Slot]Driver, f func(Driver) bool) bool {
	for _, v := range values {
		if !f(v) {
			return false
		}
	}
	return true
}

//Filter returns a slice of Drivers that satisfy a condition
//specified by the function passed in
func Filter(values map[Slot]Driver, f func(Driver) bool) []string {
	var Names []string
	for _, v := range values {
		if f(v) {
			Names = append(Names, v.Name)
		}
	}
	return Names
}

// Convert returns a new map containing the results of applying
// the function "f" to each Driver in the original  map.
func Convert(values map[Slot]Driver, f func(Driver) Driver) map[Slot]Driver {
	converted := make(map[Slot]Driver, len(values))
	for i, v := range values {
		converted[i] = f(v)
	}
	return converted
}

//Country returns the Country that matches a given code
func Country(values []locale, code string, f func(locale, string) bool) string {
	for _, v := range values {
		if f(v, code) {
			return v.Country
		}
	}
	return ""
}

//Drivers is the list of all drivers competing in formula one in 2017
var Drivers map[Slot]Driver
var locales []locale

func main() {
	Drivers = make(map[Slot]Driver)
	fill()

	fmt.Println("Lewis Hamilton's Position", Position(Drivers, "Lewis Hamilton"))

	fmt.Println("Valtteri Bottas is on a team:", Includes(Drivers, "Valtteri Bottas"))

	fmt.Println("There are drivers without points:", Has(Drivers, func(v Driver) bool {
		return v.Points == 0
	}))

	fmt.Println("There are drivers from Brazil:", Has(Drivers, func(v Driver) bool {
		return v.Country == "BRA"
	}))

	fmt.Println("All drivers have some points:", All(Drivers, func(v Driver) bool {
		return v.Points > 0
	}))

	fmt.Println("Teams and Drivers with full country name:\n", Convert(Drivers, func(v Driver) Driver {
		d := v
		d.Country = d.FullCountryName(locales, v.Country)
		return d
	}))

	fmt.Println("Only the drivers from France:", Filter(Drivers, func(v Driver) bool {
		return v.Country == "FRA"
	}))

}

func (s Slot) String() string {
	return fmt.Sprintf("%d %s", s.Position, strings.Title(strings.ToLower(s.Team)))
}
func (d Driver) String() string {
	return fmt.Sprintf(d.Name+" from "+d.Country+", has %d Points\n", d.Points)
}

//FullCountryName returns the full Country Name based on an ISO code
func (d Driver) FullCountryName(locales []locale, acronym string) string {
	return Country(locales, acronym, func(l locale, code string) bool {
		return l.ThreeLetter == code
	})
}

func fill() {
	locales = getCountries("countries.json")
	//note that there can be no ambiguity when you use literal structs
	Drivers[Slot{Position: 1, Team: "MERCEDES"}] = Driver{Points: 363, Name: "Lewis Hamilton", Country: "GBR"}
	Drivers[Slot{Position: 1, Team: "FERRARI"}] = Driver{Points: 317, Name: "Sebastian Vettel", Country: "DEU"}
	Drivers[Slot{Position: 2, Team: "MERCEDES"}] = Driver{Points: 305, Name: "Valtteri Bottas", Country: "FIN"}
	Drivers[Slot{Position: 2, Team: "FERRARI"}] = Driver{Points: 205, Name: "Kimi Räikkönen", Country: "FIN"}
	Drivers[Slot{Position: 1, Team: "RED BULL RACING"}] = Driver{Points: 200, Name: "Daniel Ricciardo", Country: "AUS"}
	Drivers[Slot{Position: 2, Team: "RED BULL RACING"}] = Driver{Points: 168, Name: "Max Verstappen", Country: "NLD"}
	Drivers[Slot{Position: 1, Team: "FORCE INDIA MERCEDES"}] = Driver{Points: 100, Name: "Sergio Perez", Country: "MEX"}
	Drivers[Slot{Position: 1, Team: "FORCE INDIA MERCEDES"}] = Driver{Points: 87, Name: "Esteban Ocon", Country: "FRA"}
	Drivers[Slot{Position: 1, Team: "RENAULT"}] = Driver{Points: 54, Name: "Carlos Sainz", Country: "ESP"}
	Drivers[Slot{Position: 2, Team: "RENAULT"}] = Driver{Points: 43, Name: "Nico Hulkenberg", Country: "DEU"}
	Drivers[Slot{Position: 1, Team: "WILLIAMS MERCEDES"}] = Driver{Points: 43, Name: "Felipe Massa", Country: "BRA"}
	Drivers[Slot{Position: 2, Team: "WILLIAMS MERCEDES"}] = Driver{Points: 40, Name: "Lance Stroll", Country: "CAN"}
	Drivers[Slot{Position: 1, Team: "HAAS FERRARI"}] = Driver{Points: 28, Name: "Romain Grosjean", Country: "FRA"}
	Drivers[Slot{Position: 2, Team: "HAAS FERRARI"}] = Driver{Points: 19, Name: "Kevin Magnussen", Country: "DNK"}
	Drivers[Slot{Position: 1, Team: "MCLAREN HONDA"}] = Driver{Points: 17, Name: "Fernando Alonso", Country: "ESP"}
	Drivers[Slot{Position: 2, Team: "MCLAREN HONDA"}] = Driver{Points: 13, Name: "Stoffel Vandoorne", Country: "BEL"}
	Drivers[Slot{Position: 3, Team: "RENAULT"}] = Driver{Points: 8, Name: "Jolyon Palmer", Country: "GBR"}
	Drivers[Slot{Position: 1, Team: "SAUBER FERRARI"}] = Driver{Points: 5, Name: "Pascal Wehrlein", Country: "DEU"}
	Drivers[Slot{Position: 1, Team: "TORO ROSSO"}] = Driver{Points: 5, Name: "Daniil Kvyat", Country: "RUS"}
	Drivers[Slot{Position: 2, Team: "SAUBER FERRARI"}] = Driver{Points: 0, Name: "Marcus Ericsson", Country: "SWE"}
	Drivers[Slot{Position: 2, Team: "TORO ROSSO"}] = Driver{Points: 0, Name: "Pierre Gasly", Country: "FRA"}
	Drivers[Slot{Position: 1, Team: "SAUBER FERRARI"}] = Driver{Points: 0, Name: "Antonio Giovinazzi", Country: "ITA"}
	Drivers[Slot{Position: 3, Team: "TORO ROSSO"}] = Driver{Points: 0, Name: "Brendon Hartley", Country: "NZL"}

}

type locale struct {
	Country     string `json:"Country"`
	TwoLetter   string `json:"Alpha-2"`
	ThreeLetter string `json:"Alpha-3"`
	Code        string `json:"UN-M49-Numeric"`
}

func getCountries(filePath string) []locale {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var locales []locale
	json.Unmarshal(raw, &locales)
	return locales
}
