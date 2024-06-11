package NewError

import (
	"errors"
	"net/http"
)

type Error struct {
	Message  string
	HttpCode int
}

func New(msg string, httpCode int) Error {
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
	ErrorGeneral         = New("internal server error", http.StatusInternalServerError)
	ErrorBadRequest      = New("bad request", http.StatusBadRequest)
	ErrorNotFound        = New(ErrNotFound.Error(), http.StatusNotFound)
	ErrorUnauthorized    = New(ErrUnauthorized.Error(), http.StatusUnauthorized)
	ErrorForbiddenAccess = New(ErrForbiddenAccess.Error(), http.StatusForbidden)
)
