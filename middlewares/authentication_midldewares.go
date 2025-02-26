package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"myproject/forum/shared"
	"myproject/forum/util"
	"net/http"
)

const UserIDContextKey = "UserID"

func AuthenticationMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := util.ExtractTokenFromRequest(c)
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

		claims, ok := jwtToken.Claims.(jwt.MapClaims)
		if !ok {
			shared.SendError(c, http.StatusBadRequest, "Invalid credentials")
			c.Abort()
			return
		}

		c.Set(UserIDContextKey, claims["sub"])
		c.Next()
	}
}
