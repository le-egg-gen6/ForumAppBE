package util

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/middlewares"
	"strings"
)

func GetRequestID(c *gin.Context) string {
	if id, ok := c.Value(middlewares.RequestIDContextKey).(string); ok {
		return id
	}
	return "unknown"
}

func ExtractTokenFromRequest(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	parts := strings.Split(bearerToken, " ")
	if len(parts) == 2 && parts[0] == "Bearer" {
		return parts[1]
	}
	return ""
}

func GetCurrentContextUserID(c *gin.Context) uint64 {
	if id, ok := c.Value(middlewares.UserIDContextKey).(uint64); ok {
		return id
	}
	return -1
}
