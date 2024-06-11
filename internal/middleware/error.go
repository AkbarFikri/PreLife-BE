package middleware

import (
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"net/http"
)

var (
	ErrorUnableToVerifyToken = NewError.New("unable to authorize user", http.StatusUnauthorized)
)
