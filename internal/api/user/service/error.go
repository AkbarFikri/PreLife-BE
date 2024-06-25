package userService

import (
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"net/http"
)

var (
	ErrorInvalidUserId     = NewError.New("invalid user id provided", http.StatusBadRequest)
	ErrorInvalidProfileId  = NewError.New("invalid profile id provided", http.StatusBadRequest)
	ErrorInvalidTimeFormat = NewError.New("invalid time format", http.StatusBadRequest)
)
