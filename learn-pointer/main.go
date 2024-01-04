package main

import "fmt"

type Coordinate struct {
	lat, lng float32
}

type Address struct {
	District, City, Province string
	Coordinate               Coordinate
}

func (address *Address) changeCity(cityName string) {
	address.City = cityName
}

func changeCity(address *Address, cityName string) {
	address.City = cityName
}

func main() {
	address1 := Address{}
	address1.District = "Setiabudi"
	address1.City = "Jakarta Selatan"
	address1.Province = "DKI Jakarta"
	address1.Coordinate = Coordinate{1, 2}

	// // Pass by value
	// address2 := address1 // deep copy
	// address2.District = "Kuningan"

	// fmt.Println(address1)
	// fmt.Println(address2)

	// Pointer
	address2Pointer := &address1
	address2Pointer.District = "Kuningan"

	fmt.Println(address1)
	fmt.Println(address2Pointer)

	// Ini jadinya buat 2 pointer
	// address2Pointer = &Address{
	// 	"Menteng", "Jakarta Selatan", "DKI Jakarta", Coordinate{2, 1},
	// }
	// fmt.Println(address1)
	// fmt.Println(address2Pointer)

	// Ini force ganti value parent nya
	*address2Pointer = Address{
		"Menteng", "Jakarta Selatan", "DKI Jakarta", Coordinate{2, 1},
	}
	fmt.Println(address1)
	fmt.Println(address2Pointer)

	// New Pointer
	address1New := new(Address)
	address2New := address1New

	address2New.City = "Padalarang"
	fmt.Println(address1New.City)
	fmt.Println(address2New.City)

	changeCity(address2New, "Bandung")
	fmt.Println(address1New.City)
	fmt.Println(address2New.City)

	address2New.changeCity("Pyongyang")
	fmt.Println(address1New.City)
	fmt.Println(address2New.City)

}
