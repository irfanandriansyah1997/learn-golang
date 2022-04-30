package booleantypes

import "fmt"

func BooleanTypeExample() {
	var myBoolean bool = true
	var anotherBoolean = false // Infered type

	var truthExample = 3 < 5
	var falseExample = 10 != 10

	// Short Circuiting
	var res1 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	var res2 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true

	fmt.Println(myBoolean, anotherBoolean, truthExample, falseExample, res1, res2)
}
