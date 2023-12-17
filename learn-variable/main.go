package main

import fmt2 "fmt"

func main() {
	// Variable (1)
	var name string
	// Variable (2)
	age := 10

	name = "John Doe"
	fmt2.Printf("Hello my name is %v, I %v years old\n", name, age)

	name = "John"
	age = 12
	fmt2.Printf("Hello my name is %v, I %v years old\n", name, age)
}
