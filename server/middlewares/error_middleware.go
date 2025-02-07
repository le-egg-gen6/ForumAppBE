package middlewares

import (
	"github.com/gin-gonic/gin"
	"myproject/forum/server/shared"
	"net/http"
)

func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if c.Errors != nil && len(c.Errors) > 0 {
			err := c.Errors.Last().Err
			if appErr, ok := err.(*shared.AppError); ok {
				//log
				shared.SendError(c, appErr.Code, appErr.Message)
			} else {
				//log
				shared.SendError(c, http.StatusInternalServerError, "Internal Server Error")
			}
		}
	}
}
