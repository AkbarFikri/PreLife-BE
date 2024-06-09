package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func (m *Middleware) AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") == "" {
			c.AbortWithStatusJSON(401, gin.H{
				"code": 401,
			})
			m.log.Warnf("request without authorization header detected : %v access to route %s", c.Request.RemoteAddr, c.Request.RequestURI)
			return
		}

		if !strings.Contains(c.GetHeader("Authorization"), "Bearer") {
			c.AbortWithStatusJSON(401, gin.H{
				"code": 401,
			})
			m.log.Warnf("request without authorization header detected : %v access to route %s", c.Request.RemoteAddr, c.Request.RequestURI)
			return
		}

		header := c.GetHeader("Authorization")
		token := strings.Split(header, "Bearer ")[1]

		idToken, err := m.authClient.VerifyIDToken(c, token)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"code": err.Error(),
			})
			m.log.Warnf("unable to authorize user : %v access to route %s", c.ClientIP(), c.Request.RequestURI)
		}

		c.Set("idToken", idToken)
		c.Next()
	}
}
