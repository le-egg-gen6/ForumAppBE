package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("remote_addr", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int64("elapsed_time", time.Since(start).Milliseconds()),
		)
	}
}
