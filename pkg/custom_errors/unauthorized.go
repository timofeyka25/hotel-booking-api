package custom_errors

type UnauthorizedError struct {
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}
