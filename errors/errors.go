package errors

// CustomError a custom error to show a message error
// @Description a custom error to show a message error
type CustomError struct {
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}
