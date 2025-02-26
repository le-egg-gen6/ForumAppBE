package util

import (
	"context"
	"github.com/gin-gonic/gin"
	"myproject/forum/middlewares"
	"strings"
)

func GetRequestID(ctx context.Context) string {
	if id, ok := ctx.Value(middlewares.RequestIDKey).(string); ok {
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
