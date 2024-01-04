package main

import "fmt"

type User struct {
	Address, Name string
	Age           int
	Friends       []User
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}

type Bird struct {
	Animal
	SpeedKPH float32
	CanFly   bool
}

func (animal Bird) sayHello() {
	fmt.Printf("Hello %s \n", animal.Name)
}

func main() {
	user := User{
		Name:    "John A",
		Age:     16,
		Address: "Jakarta",
		Friends: []User{
			{
				Name:    "John B",
				Age:     17,
				Address: "Bandung",
				Friends: []User{},
			},
			{
				Name:    "John C",
				Age:     11,
				Address: "Bogor",
				Friends: []User{},
			},
		},
	}

	fmt.Println((user.Friends))

	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	b.sayHello()

	fmt.Println(b)
}
