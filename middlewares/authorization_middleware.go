package middlewares

import "github.com/gin-gonic/gin"

func AuthorizationMiddleware(requiredRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
