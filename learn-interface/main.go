package main

import "fmt"

type HasName interface {
	GetName() string
}

type User struct {
	Name, Address string
	Age           int
}

func (user User) GetName() string {
	return user.Name
}

type Animal struct {
	Name   string
	Origin string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(instance HasName) {
	fmt.Printf("Hello %s \n", instance.GetName())
}

func main() {
	user := User{"John Doe", "Bandung", 30}
	animal := Animal{"Cat", "earth"}

	SayHello(user)
	SayHello(animal)
}
