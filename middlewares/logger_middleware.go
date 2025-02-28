package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myproject/forum/logger"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		if logger.LogInstance != nil {
			logger.LogInstance.Info("HTTP Request",
				zap.String("method", c.Request.Method),
				zap.String("url", c.Request.URL.Path),
				zap.Int("status", c.Writer.Status()),
				zap.String("remote_addr", c.ClientIP()),
				zap.String("user_agent", c.Request.UserAgent()),
				zap.Int64("elapsed_time", time.Since(start).Milliseconds()),
			)
		}
	}
}
