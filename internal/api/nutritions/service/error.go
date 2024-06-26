package nutritionsService

import (
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"net/http"
)

var (
	ErrorInvalidPictureProvided = NewError.New("invalid picture provided, picture must be food", http.StatusBadRequest)
)
