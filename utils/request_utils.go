package utils

import (
	"forum/constant"
	"github.com/gin-gonic/gin"
)

func GetCurrentContextUserID(c *gin.Context) int64 {
	if id, ok := c.Value(constant.UserIDContextKey).(int64); ok {
		return id
	}
	return -1
}

func GetCurrentContextAuthorizationToken(c *gin.Context) string {
	if token, ok := c.Value(constant.AuthorizationTokenContextKey).(string); ok {
		return token
	}
	return ""
}

func GetRequestHeader(c *gin.Context, key string) string {
	return c.Request.Header.Get(key)
}

func GetRequestParam(c *gin.Context, key string) string {
	return c.Param(key)
}
