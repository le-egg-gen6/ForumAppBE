package middlewares

import (
	"forum/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AccountValidationMiddlewares() gin.HandlerFunc {
	return func(c *gin.Context) {
		validated := utils.GetCurrentContextUserValidatedStatus(c)
		if !validated {
			c.AbortWithStatus(http.StatusNotAcceptable)
			return
		}
		c.Next()
	}
}
