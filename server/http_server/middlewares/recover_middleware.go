package middlewares

import (
	"forum/shared"
	"github.com/gin-gonic/gin"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				shared.SendInternalServerError(c)
				c.Abort()
			}
		}()
		c.Next()
	}
}
