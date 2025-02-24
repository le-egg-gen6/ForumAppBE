package middlewares

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/shared"
	"net/http"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				shared.SendError(c, http.StatusInternalServerError, "Internal Server Error")
				c.Abort()
			}
		}()
		c.Next()
	}
}
