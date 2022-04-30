package main

import (
	"fmt"

	booleanTypes "github.com/irfanandriansyah1997/basic_types/boolean_types"
	characters "github.com/irfanandriansyah1997/basic_types/characters"
	numericOperations "github.com/irfanandriansyah1997/basic_types/numeric_operations"
	numeric "github.com/irfanandriansyah1997/basic_types/numeric_types"
)

func main() {
	fmt.Println("Hello World")

	numeric.NumericTypesExample()
	characters.CharactersExample()
	numericOperations.NumericOperationsExample()
	booleanTypes.BooleanTypeExample()
}
