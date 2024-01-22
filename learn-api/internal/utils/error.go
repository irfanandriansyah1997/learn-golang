package utils

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{err}
}

func PanicIfNotFoundError(err error) {
	if err != nil {
		panic(NewNotFoundError(err.Error()))
	}
}
