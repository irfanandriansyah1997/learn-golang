package utils

import "fmt"

///////////////////////////////////////////////////////////
// Custom Error
///////////////////////////////////////////////////////////

type NotFoundError struct {
	errorMessage string
}

func (n *NotFoundError) Error() string {
	return fmt.Sprintf("[Not Found]: %s", n.errorMessage)
}

func NewNotFoundError(err string) error {
	return &NotFoundError{err}
}

///////////////////////////////////////////////////////////
// Panic Handler
///////////////////////////////////////////////////////////

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicIfNotFoundError(err error) {
	if err != nil {
		panic(NewNotFoundError(err.Error()))
	}
}
