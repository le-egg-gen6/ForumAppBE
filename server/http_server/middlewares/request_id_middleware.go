package middlewares

import (
	"forum/constant"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.New().String()
		c.Set(constant.RequestIDContextKey, requestID)
		c.Writer.Header().Set("X-Request-ID", requestID)
		c.Next()
	}
}
