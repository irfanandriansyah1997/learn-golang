package helper

import "fmt"

func init() {
	fmt.Println("initialize helper package")
}

func SayHello(name string) string {
	return "Hello " + name
}
