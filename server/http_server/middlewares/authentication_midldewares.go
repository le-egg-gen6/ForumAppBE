package middlewares

import (
	"forum/3rd_party_service/redis_service"
	"forum/constant"
	"forum/shared"
	"forum/utils"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func AuthenticationMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := ExtractTokenFromRequest(c)
		if tokenStr == "" {
			shared.SendUnauthorized(c)
			c.Abort()
			return
		}

		tokenUsed, _ := redis_service.Get[bool](tokenStr)
		if tokenUsed {
			shared.SendUnauthorized(c)
			c.Abort()
			return
		}

		jwtToken, err := utils.ValidateToken(tokenStr)
		if err != nil {
			shared.SendUnauthorized(c)
			c.Abort()
			return
		}

		claims, ok := jwtToken.Claims.(*utils.CustomClaims)
		if !ok {
			shared.SendUnauthorized(c)
			c.Abort()
			return
		}

		if claims.IssuedAt.Unix() > time.Now().Unix() {
			shared.SendUnauthorized(c)
			c.Abort()
			return
		}

		if claims.ExpiresAt.Unix() < time.Now().Unix() {
			shared.SendUnauthorized(c)
			c.Abort()
			return
		}

		c.Set(constant.UserIDContextKey, claims.UserID)
		c.Set(constant.AuthorizationTokenContextKey, tokenStr)
		c.Set(constant.UserValidatedStatusKey, claims.Validated)
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
