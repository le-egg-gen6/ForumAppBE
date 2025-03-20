package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"myproject/forum/logger"
	"myproject/forum/util"
	"time"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		logger.GetInstance().Info("HTTP Request",
			zap.String(RequestIDContextKey, util.GetRequestID(c)),
			zap.String("method", c.Request.Method),
			zap.String("url", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.String("remote_addr", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
			zap.Int64("elapsed_time", time.Since(start).Milliseconds()),
		)
	}
}
