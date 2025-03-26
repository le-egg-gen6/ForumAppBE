package middlewares

import (
	"forum/constant"
	"forum/shared"
	"forum/util"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

func AuthenticationMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := ExtractTokenFromRequest(c)
		if tokenStr == "" {
			shared.SendError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		jwtToken, err := util.ValidateToken(tokenStr)
		if err != nil {
			shared.SendError(c, http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}

		claims, ok := jwtToken.Claims.(*jwt.RegisteredClaims)
		if !ok {
			shared.SendError(c, http.StatusBadRequest, "Invalid credentials")
			c.Abort()
			return
		}

		c.Set(constant.UserIDContextKey, claims.Subject)
		c.Next()
	}
}

func ExtractTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	parts := strings.Split(bearerToken, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}
