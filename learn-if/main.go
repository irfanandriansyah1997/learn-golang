package main

import "fmt"

func isKeyAttributeExists(args map[string]int, keyName string) *int {
	if val, ok := args[keyName]; ok {
		return &val
	}

	return nil
}

func main() {
	statePopulations := map[string]int{
		"Jakarta": 1000,
		"Bandung": 2000,
		"Bogor":   3000,
	}

	fmt.Println(*isKeyAttributeExists(statePopulations, "Jakarta"))

	if 5 < 10 {
		fmt.Println("5 is lower than 10")
	}

}
