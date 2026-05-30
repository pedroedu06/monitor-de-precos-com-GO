package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pedroedu06/monitor-de-precos-com-GO/internal/pkg/jwt"
	resterr "github.com/pedroedu06/monitor-de-precos-com-GO/rest_err"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			restErr := resterr.NewUnauthorizedErr("JWT token not found")
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, err := jwt.ValidateToken(tokenString)
		if err != nil {
			restErr := resterr.NewUnauthorizedErr("invalid token")
			c.JSON(restErr.Code, restErr)
			c.Abort()
			return
		}

		c.Set("UserID", claims.UserID)
		c.Next()
	}
}
