package main

import (
	json "encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

//Name struct - holds a name
type Name struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

//Driver struct - describes a driver
type Driver struct {
	Name
	Country     string            `json:"country"` //when encoded the tag will be country not Country
	Constructor string            `json:"constructor"`
	Score       int               `json:"score"`
	Results     map[string]Result `json:"results"`
}

//Result struct describes a race result
type Result struct {
	Starts       time.Time
	Race         string
	Standing     string
	Points       int
	GridPosition int
}

//Race struct describes a racing circuit somewhere that is on the F1 Calendar
type Race struct {
	ID      string `json:"Id"`
	Round   string `json:"Round"`
	Race    string `json:"Race"`
	Country string `json:"Country"`
	Circuit string `json:"Circuit"`
	Date    string `json:"Date"`
}

func main() {
	//marshal primitive types
	aBool, _ := json.Marshal(true)
	fmt.Println(string(aBool))

	anInt, _ := json.Marshal(345)
	fmt.Println(string(anInt))

	aFloat, _ := json.Marshal((345 + 302) / 2)
	fmt.Println(string(aFloat))

	aString, _ := json.Marshal("Lewis Hamilton")
	fmt.Println(string(aString))

	var races []Race
	races = getRaces("./races.json")
	fmt.Println("First 3 Races...................")
	fmt.Println(races[0:3])
	fmt.Println("The last race is held on", races[len(races)-1].Date)

	var results map[string]Result
	results = make(map[string]Result)
	drivers := map[int]Driver{
		1:  {Name: Name{First: "Lewis", Last: "Hamilton"}, Country: "GBR", Constructor: "Mercedes", Score: 345, Results: results},
		2:  {Name: Name{First: "Sebastian", Last: "Vettel"}, Country: "GER", Constructor: "FERRARI", Score: 302, Results: results},
		3:  {Name: Name{First: "Valtteri", Last: "Bottas"}, Country: "FIN", Constructor: "MERCEDES", Score: 280, Results: results},
		4:  {Name: Name{First: "Daniel", Last: "Ricciardo"}, Country: "AUS", Constructor: "RED BULL RACING TAG HEUER", Score: 200, Results: results},
		5:  {Name: Name{First: "Kimi", Last: "Räikkönen"}, Country: "FIN", Constructor: "FERRARI", Score: 193, Results: results},
		6:  {Name: Name{First: "Max", Last: "Verstappen"}, Country: "NED", Constructor: "RED BULL RACING TAG HEUER", Score: 158, Results: results},
		7:  {Name: Name{First: "Sergio", Last: "Perez"}, Country: "MEX", Constructor: "FORCE INDIA MERCEDES", Score: 94, Results: results},
		8:  {Name: Name{First: "Esteban", Last: "Ocon"}, Country: "FRA", Constructor: "FORCE INDIA MERCEDES", Score: 83, Results: results},
		9:  {Name: Name{First: "Carlos", Last: "Sainz"}, Country: "ESP", Constructor: "RENAULT", Score: 54, Results: results},
		10: {Name: Name{First: "Felipe", Last: "Massa"}, Country: "BRA", Constructor: "WILLIAMS MERCEDES", Score: 42, Results: results},
		11: {Name: Name{First: "Lance", Last: "Stroll"}, Country: "CAN", Constructor: "WILLIAMS MERCEDES", Score: 40, Results: results},
		12: {Name: Name{First: "Nico", Last: "Hulkenberg"}, Country: "GER", Constructor: "RENAULT", Score: 35, Results: results},
		13: {Name: Name{First: "Romain", Last: "Grosjean"}, Country: "FRA", Constructor: "HAAS FERRARI", Score: 28, Results: results},
		15: {Name: Name{First: "Fernando", Last: "Alonso"}, Country: "ESP", Constructor: "MCLAREN HONDA", Score: 15, Results: results},
		14: {Name: Name{First: "Kevin", Last: "Magnussen"}, Country: "DEN", Constructor: "HAAS FERRARI", Score: 19, Results: results},
		16: {Name: Name{First: "Stoffel", Last: "Vandoorne"}, Country: "BEL", Constructor: "MCLAREN HONDA", Score: 13, Results: results},
		17: {Name: Name{First: "Jolyon", Last: "Palmer"}, Country: "GBR", Constructor: "RENAULT", Score: 8, Results: results},
		18: {Name: Name{First: "Pascal", Last: "Wehrlein"}, Country: "GER", Constructor: "SAUBER FERRARI", Score: 5, Results: results},
		19: {Name: Name{First: "Daniil", Last: "Kvyat"}, Country: "RUS", Constructor: "TORO ROSSO", Score: 5, Results: results},
		20: {Name: Name{First: "Marcus", Last: "Ericsson"}, Country: "SWE", Constructor: "SAUBER FERRARI", Score: 0, Results: results},
		21: {Name: Name{First: "Pierre", Last: "Gasly"}, Country: "FRA", Constructor: "TORO ROSSO", Score: 0, Results: results},
		22: {Name: Name{First: "Antonio", Last: "Giovinazzi"}, Country: "ITA", Constructor: "SAUBER FERRARI", Score: 0, Results: results},
		23: {Name: Name{First: "Brendon", Last: "Hartley"}, Country: "NZL", Constructor: "TORO ROSSO", Score: 0, Results: results},
	}

	driver, ok := drivers[1]
	if ok {
		driver.Results["GBR"] = Result{time.Now(), "Silverstone", "1", 0, 1}
		race, ok := driver.Results["GBR"]
		if ok {
			race.Points = 25
			driver.Results["GBR"] = race
		}
	}
	fmt.Println("Driver........................")
	fmt.Println(driver)
	fmt.Println("Driver to JSON................")
	fmt.Println(toJSON(driver))
	err := ioutil.WriteFile("driver.json", toJSON(driver), 0644)
	if err != nil {
		panic("Couldn't write drivers to file driver.json")
	}
	//comment this out to start the exercise
	// fmt.Println("Encoding.......................")
	// enc := json.NewEncoder(os.Stdout) //will encode to anything that takes a stream of bytes
	// enc.SetIndent(" ", "\t")
	// for k := range drivers {
	// 	if k > 10 {
	// 		delete(drivers, k)
	// 	}
	// }
	// if err := enc.Encode(&drivers); err != nil {
	// 	log.Println(err)
	// }

}
func (p Race) String() string {
	return string(toJSON(p))
}

func toJSON(i interface{}) []byte {
	bytes, err := json.Marshal(i)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return bytes
}

func getRaces(filePath string) []Race {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var races []Race
	json.Unmarshal(raw, &races)
	return races
}
