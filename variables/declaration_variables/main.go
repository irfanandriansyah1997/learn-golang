package declarationvariables

import "fmt"

func Execute() {
	var message string = "Hello world"
	var score int = 100
	var rating float64 = 4.6

	fmt.Println(message, score, rating)

	// Multiple Declaration
	var (
		employeedId         int    = 5
		firstName, lastName string = "Satoshi", "Nakamoto"
	)
	fmt.Println(employeedId, firstName, lastName)

	// Short Declaration
	name := "Irfan Andriansyah"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer)

}
