package zero_values

import "fmt"

func ZeroValuesExample() {
	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)
}
