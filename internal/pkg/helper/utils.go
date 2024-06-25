package helper

import (
	"crypto/rand"
	"encoding/base64"
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	"github.com/gin-gonic/gin"
)

func GenerateUID(length int) string {
	randomBytes := make([]byte, length)

	_, err := rand.Read(randomBytes)
	if err != nil {
		return ""
	}

	uid := base64.RawURLEncoding.EncodeToString(randomBytes)
	return uid[:length]
}

func GetUserLoginData(c *gin.Context) dto.UserTokenData {
	getUser, _ := c.Get("user")
	user := getUser.(dto.UserTokenData)

	return user
}
