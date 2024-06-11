package authService

import (
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"net/http"
)

var (
	//Authentication
	ErrorEmailAlreadyUsed       = NewError.New("email already used", http.StatusBadRequest)
	ErrorInvalidEmailOrPassword = NewError.New("invalid email or password", http.StatusBadRequest)
	ErrorInvalidAccessToken     = NewError.New("access token invalid or expire", http.StatusBadRequest)
	ErrorInvalidDateFormat      = NewError.New("invalid date format", http.StatusBadRequest)
)
