package main

import (
	"fmt"
)

type film struct {
	title string
	gross float64
	year  int
}

func main() {
	m := film{title: "Avatar", gross: 2787965087, year: 2008}
	//type switch
	typeName := func(i interface{}) string {
		switch T := i.(type) {
		case bool:
			return fmt.Sprintf("%t is a bool", T)
		case int:
			return fmt.Sprintf("%d is an int", T)
		case string:
			return fmt.Sprintf("%s is a string", T)
		case film:
			return fmt.Sprintf("%s is a film", T.title)
		default:
			return fmt.Sprintf("unknown type:%g", T)
		}
	}
	fmt.Println(typeName(m))
	fmt.Println(typeName("Index Exchange"))
	fmt.Println(typeName(true))
	fmt.Println(typeName(22.5))

}
