package helper

import (
	"firebase.google.com/go/v4/auth"
	"github.com/AkbarFikri/PreLife-BE/internal/domain"
	"github.com/gin-gonic/gin"
)

func UserDataFromToken(ctx *gin.Context) (domain.User, error) {
	userToken, _ := ctx.Get("idToken")
	token := userToken.(*auth.Token)
	return domain.User{
		ID:    token.UID,
		Email: token.Claims["email"].(string),
	}, nil
}
