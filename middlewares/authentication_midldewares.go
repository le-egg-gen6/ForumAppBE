package middlewares

import (
	"forum/constant"
	"forum/shared"
	"forum/utils"
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

		jwtToken, err := utils.ValidateToken(tokenStr)
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
		c.Set(constant.AuthorizationTokenContextKey, tokenStr)
		c.Next()
	}
}

func ExtractTokenFromRequest(c *gin.Context) string {
	bearerToken := utils.GetRequestHeader(c, constant.AuthorizationHeader)
	parts := strings.Split(bearerToken, " ")
	if len(parts) == 2 && parts[0] == constant.AuthorizationHeaderPrefix {
		return parts[1]
	}
	return ""
}
