package customErrors

type AlreadyReservedError struct {
	message string
}

func NewAlreadyReservedError(message string) *AlreadyReservedError {
	return &AlreadyReservedError{
		message: message,
	}
}

func (e *AlreadyReservedError) Error() string {
	return e.message
}
