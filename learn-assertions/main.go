package main

import "fmt"

func random() any {
	return "Ok"
}

func main() {
	var result = random()

	// Cara 1 ❌
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	// Cara 2 ✅
	switch result.(type) {
	case string:
		fmt.Println("String", result)
	case int:
		fmt.Println("Int", result)
	default:
		fmt.Println("Unknown")
	}

}
