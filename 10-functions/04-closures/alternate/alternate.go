package main

import (
	"fmt"
)

//Rule struct
type Rule struct {
	Key   string
	Value string
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//MonitorRules takes any number of rules and prints them
func MonitorRules(rules ...Rule) {
	for _, rule := range rules {
		fmt.Println(rule)
	}
}
func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())
	RuleA := Rule{"block-servers", "true"}
	RuleB := Rule{"accept-credit", "false"}
	MonitorRules(RuleA, RuleB)
}
