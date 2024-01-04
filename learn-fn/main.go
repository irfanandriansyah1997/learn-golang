package main

import "fmt"

// Normal Function
func getFullName() (firstName, lastName string, age int) {
	firstName = "John"
	lastName = "Doe"
	age = 10
	return firstName, lastName, age
}

// Variadic Function (spread operator)
func sumAllNumber(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// Callback function
func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	}

	return name
}

type _FilterFn func(string) string

func sayHello(name string, filter _FilterFn) {
	fmt.Printf("Hello %v \n", filter(name))
}

// reqursive function

func factorial(value int) int {
	if value > 1 {
		return factorial(value-1) * value
	}

	return value
}

func main() {
	firstName, _, _ := getFullName()
	spamFilterAnonymousFn := func(name string) string {
		if name == "Anjing" {
			return "...."
		}

		return name
	}

	fmt.Println(firstName)
	fmt.Println(sumAllNumber(10, 11, 20))

	numbers := []int{10, 11, 20}
	fmt.Println(sumAllNumber(numbers...))

	sayHello("Anjing", spamFilter)
	sayHello("Andi", spamFilter)
	sayHello("Anjing", spamFilterAnonymousFn)
	sayHello("Andi", spamFilterAnonymousFn)

	fmt.Println(factorial(5))

}
