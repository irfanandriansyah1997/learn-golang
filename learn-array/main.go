// Package main
package main

import "fmt"

func main() {
	/////////////////////////////////////////////////////////////////
	// Assign Array
	/////////////////////////////////////////////////////////////////

	var users = [...]string{
		"users 1",
		"users 2",
		"users 3",
	}

	fmt.Println(users) // users 1, users 2, users 3

	// set value array

	users[1] = "updated users 2"
	fmt.Println(users) // users 1, updated users 2, users 3

	/////////////////////////////////////////////////////////////////
	// Slice Array
	/////////////////////////////////////////////////////////////////

	sliceUsers := users[1:3]
	fmt.Println(sliceUsers, len(sliceUsers), cap(sliceUsers)) // updated users 2, users 3
	sliceUsers[0] = "re-updated users 2"
	sliceUsers = append(sliceUsers, "new user")
	fmt.Println(sliceUsers, len(sliceUsers), cap(sliceUsers)) // updated users 2, users 3, new user

	sliceUsers2 := users[2:]
	fmt.Println(sliceUsers2) // users 3

	sliceUsers3 := users[:2]
	fmt.Println(sliceUsers3) // users 1, updated users 2

}
