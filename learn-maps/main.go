package main

import "fmt"

func main() {
	statePopulations := map[string]int{
		"jakarta": 1172917911,
		"bandung": 2271970000,
		"bogor":   10000,
	}
	fmt.Println(statePopulations)

	statePopulations["sample"] = 100001
	fmt.Println(statePopulations)

	delete(statePopulations, "sample")
	fmt.Println(statePopulations)

	pop, ok := statePopulations["bog"]
	fmt.Println(pop, ok)
	fmt.Println(len(statePopulations))

}
