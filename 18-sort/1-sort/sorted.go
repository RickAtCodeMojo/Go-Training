package main

import (
	"fmt"
	"sort"
)

func main() {

	//use make to create a map
	drivers := make(map[string]int)

	//set and assign values using map[key] = value
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
	drivers["Marcus Ericsson"] = 0
	drivers["Pierre Gasly"] = 0
	drivers["Antonio Giovinazzi"] = 0
	drivers["Brendon Hartley"] = 0

	fmt.Println("")
	fmt.Println("All Drivers ............")
	//print keys and values
	print(drivers)

	fmt.Println("Sorted Drivers ............")
	for _, res := range sortedKeys(drivers) {
		fmt.Println(res, drivers[res])
	}

}
func byValue(i, j int) bool {
	return false
}
func print(m map[string]int) {
	for k, v := range m {
		fmt.Println(k, v)
	}

}

type sortedMap struct {
	m map[string]int
	s []string
}

func (sm *sortedMap) Len() int {
	return len(sm.m)
}

func (sm *sortedMap) Less(i, j int) bool {
	return sm.m[sm.s[i]] > sm.m[sm.s[j]]
}

func (sm *sortedMap) Swap(i, j int) {
	sm.s[i], sm.s[j] = sm.s[j], sm.s[i]
}

func sortedKeys(m map[string]int) []string {
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
