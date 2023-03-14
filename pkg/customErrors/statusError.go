package customErrors

type StatusError struct {
	message string
}

func NewStatusError(message string) *StatusError {
	return &StatusError{
		message: message,
	}
}

func (e *StatusError) Error() string {
	return e.message
}
