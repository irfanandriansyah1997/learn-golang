package numericoperations

import (
	"fmt"
	"math"
)

func NumericOperationsExample() {
	var a, b = 4, 5

	var output1 = (a + b) * (a + b) / 2 // Arithmetic operations

	a++     // Increment variable a by 1
	a += 10 // Increment variable b by 10

	var output2 = a ^ b // Bitwise XOR

	var radius = 7.0
	var output3 = math.Pi * radius * radius

	fmt.Printf("operation 1 : %v, operation 2 : %v, operation 3 : %v, math.Pi: %f \n",
		output1, output2, output3, math.Pi)
}
