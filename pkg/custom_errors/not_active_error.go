package custom_errors

type NotActiveError struct {
	message string
}

func NewNotActiveError(message string) *NotActiveError {
	return &NotActiveError{
		message: message,
	}
}

func (e *NotActiveError) Error() string {
	return e.message
}
