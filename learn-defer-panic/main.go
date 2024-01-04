package main

import "fmt"

func endApp() {
	fmt.Println("End App")
	message := recover()

	if message != nil {
		fmt.Printf("Panic: %v \n", message)
	}
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("Oops Error")
	}
}

func main() {
	runApp(true)
}
