package customerror

type ValidationError struct {
	Message string
}

func (e *ValidationError) Error() string {
	return e.Message
}

func (e *ValidationError) CustomError() string {
	return "[Custom Log] Validation Error: " + e.Message
}

type NotFoundError struct {
	Message string
}

func (e *NotFoundError) Error() string {
	return e.Message
}

func (e *NotFoundError) CustomError() string {
	return "[Custom Log] Not Found Error: " + e.Message
}
