// Package Main
package main

import "fmt"

func main() {
	/////////////////////////////////////////////////////////////////
	// Variable
	/////////////////////////////////////////////////////////////////

	//  Cara Assign Variable
	var name string
	age := 10

	name = "John Doe"
	fmt.Printf("Hello my name is %v, I %v years old\n", name, age)

	name = "John"
	age = 12
	fmt.Printf("Hello my name is %v, I %v years old\n", name, age)

	// Cara Assign Variable Multiple
	var (
		firstName = "John"
		lastName  = "Doe"
	)

	fmt.Printf("Hello my name is %v %v\n", firstName, lastName)

	/////////////////////////////////////////////////////////////////
	// Constant
	/////////////////////////////////////////////////////////////////

	const (
		address    = "Karet Belakang"
		postalCode = 40001
	)

	fmt.Println(address, postalCode)

	/////////////////////////////////////////////////////////////////
	// Casting Variable
	/////////////////////////////////////////////////////////////////

	var number32 int32 = 32767
	var number64 = int64(number32)
	var number16 = int16(number32)

	fmt.Println(number32, number64, number16)
}
