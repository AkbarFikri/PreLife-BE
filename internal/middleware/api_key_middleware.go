package middleware

import (
	NewError "github.com/AkbarFikri/PreLife-BE/internal/pkg/error"
	"github.com/AkbarFikri/PreLife-BE/internal/pkg/response"
	"github.com/gin-gonic/gin"
	"os"
)

func (m Middleware) ApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			m.log.Warnf("request without API_KEY at route %s with body %v", c.Request.RequestURI, c.Request.Body)
			response.New(response.WithError(NewError.ErrorForbiddenAccess)).SendAbort(c)
			return
		}

		key := c.GetHeader("API_KEY")
		if key != os.Getenv("API_KEY") {
			m.log.Warnf("request with invalid API_KEY at route %s with body %v", c.Request.RequestURI, c.Request.Body)
			response.New(response.WithError(NewError.ErrorForbiddenAccess)).SendAbort(c)
			return
		}

		c.Next()
	}
}
