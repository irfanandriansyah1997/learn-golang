package main

import (
	declarationVariable "github.com/irfanandriansyah1997/variables/declaration_variables"
	typeInference "github.com/irfanandriansyah1997/variables/type_inference"
	zeroValues "github.com/irfanandriansyah1997/variables/zero_values"
)

func main() {
	declarationVariable.Execute()
	typeInference.TypeInferenceExample()
	zeroValues.ZeroValuesExample()
}
