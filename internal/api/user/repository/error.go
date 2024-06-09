package userRepository

import (
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"net/http"
)

var (
	ErrorExecContext = NewError.NewError("Something when wrong when exec context", http.StatusInternalServerError)
)
