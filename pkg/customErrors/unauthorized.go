package customErrors

type UnauthorizedError struct {
}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func (e *UnauthorizedError) Error() string {
	return "Unauthorized"
}
