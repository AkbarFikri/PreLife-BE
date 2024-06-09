package NewError

import (
	"errors"
	"net/http"
)

type Error struct {
	Message  string
	HttpCode int
}

func NewError(msg string, httpCode int) Error {
	return Error{
		Message:  msg,
		HttpCode: httpCode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbiddenAccess = errors.New("forbidden access")
)

var (
	ErrorGeneral         = NewError("internal server error", http.StatusInternalServerError)
	ErrorBadRequest      = NewError("bad request", http.StatusBadRequest)
	ErrorNotFound        = NewError(ErrNotFound.Error(), http.StatusNotFound)
	ErrorUnauthorized    = NewError(ErrUnauthorized.Error(), http.StatusUnauthorized)
	ErrorForbiddenAccess = NewError(ErrForbiddenAccess.Error(), http.StatusForbidden)

	//Authentication
	ErrorEmailAlreadyUsed       = NewError("email already used", http.StatusBadRequest)
	ErrorInvalidEmailOrPassword = NewError("invalid email or password", http.StatusBadRequest)
	ErrorInvalidAccessToken     = NewError("access token invalid or expire", http.StatusBadRequest)
)
