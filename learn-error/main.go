package main

import (
	"fmt"
	"learn-error/calculator"
	"learn-error/customerror"
)

func main() {
	result, err := calculator.Pembagian(10, 0)

	if err != nil {
		if validationErr, ok := err.(*customerror.ValidationError); ok {
			fmt.Println(validationErr.CustomError())
		} else if notFoundErr, ok := err.(*customerror.NotFoundError); ok {
			fmt.Println(notFoundErr.CustomError())
		} else {
			fmt.Println(err.Error())
		}

		switch result := err.(type) {
		case *customerror.NotFoundError:
		case *customerror.ValidationError:
			fmt.Println(result.CustomError())
		default:
			fmt.Println(result.Error())
		}

	} else {
		fmt.Println(result)
	}
}
