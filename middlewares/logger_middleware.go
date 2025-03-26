package middlewares

import (
	"forum/constant"
	"forum/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.GetLogInstance().Info("HTTP Request",
			zap.String(constant.RequestIDContextKey, GetRequestID(c)),
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("remote_addr", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int64("elapsed_time", time.Since(start).Milliseconds()),
		)
	}
}

func GetRequestID(c *gin.Context) string {
	if id, ok := c.Value(constant.RequestIDContextKey).(string); ok {
		return id
	}
	return "unknown"
}
