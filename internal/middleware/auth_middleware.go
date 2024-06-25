package middleware

import (
	"github.com/AkbarFikri/PreLife-BE/internal/dto"
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

func (m Middleware) AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			m.log.Warnf("request without authorization header detected : %v access to route %s", c.Request.RemoteAddr, c.Request.RequestURI)
			response.New(response.WithError(NewError.ErrorForbiddenAccess)).SendAbort(c)
			return
		}

		if !strings.Contains(c.GetHeader("Authorization"), "Bearer") {
			m.log.Warnf("request without authorization header detected : %v access to route %s", c.Request.RemoteAddr, c.Request.RequestURI)
			response.New(response.WithError(NewError.ErrorForbiddenAccess)).SendAbort(c)
			return
		}

		header := c.GetHeader("Authorization")
		token := strings.Split(header, "Bearer ")[1]

		idToken, err := m.authClient.VerifyIDToken(c, token)
		if err != nil {
			m.log.Warnf("unable to authorize user : %v access to route %s with token %v", c.Request.RemoteAddr, c.Request.RequestURI, token)
			response.New(response.WithError(ErrorUnableToVerifyToken)).SendAbort(c)
			return
		}

		claims := idToken.Claims
		user := dto.UserTokenData{
			ID:          claims["id"].(string),
			Email:       claims["email"].(string),
			RoleId:      claims["role_id"].(float64),
			ProfileId:   claims["profile_id"].(string),
			ProfileType: claims["profile_type"].(float64),
		}

		c.Set("user", user)
		c.Next()
	}
}
