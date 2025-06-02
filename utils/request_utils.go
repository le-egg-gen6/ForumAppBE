package utils

import (
	"forum/constant"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

func GetCurrentContextUserID(c *gin.Context) int {
	if id, ok := c.Value(constant.UserIDContextKey).(int); ok {
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

func GetCurrentContextUserValidatedStatus(c *gin.Context) bool {
	if validated, ok := c.Value(constant.UserValidatedStatusKey).(bool); ok {
		return validated
	}
	return false
}

func GetRequestHeader(c *gin.Context, key string) string {
	return c.Request.Header.Get(key)
}

func GetRequestParam(c *gin.Context, key string) string {
	return c.Param(key)
}

func GetRequestMultipartFileHeaders(c *gin.Context, fieldName string) []*multipart.FileHeader {
	if c.Request.MultipartForm == nil || c.Request.MultipartForm.File == nil {
		return nil
	}
	return c.Request.MultipartForm.File[fieldName]
}

func GetRequestMultipartFiles(c *gin.Context, fieldName string) ([]*multipart.File, []*multipart.FileHeader, error) {
	headers := GetRequestMultipartFileHeaders(c, fieldName)
	if len(headers) == 0 {
		return nil, nil, nil
	}

	files := make([]*multipart.File, 0, len(headers))
	for _, header := range headers {
		file, err := header.Open()
		if err != nil {
			// Close any files we've already opened
			for _, f := range files {
				(*f).Close()
			}
			return nil, nil, err
		}
		files = append(files, &file)
	}

	return files, headers, nil
}

func GetFirstMultipartFileHeader(c *gin.Context, fieldName string) *multipart.FileHeader {
	headers := GetRequestMultipartFileHeaders(c, fieldName)
	if len(headers) == 0 {
		return nil
	}
	return headers[0]
}

func GetFirstMultipartFile(c *gin.Context, fieldName string) (*multipart.File, *multipart.FileHeader, error) {
	header := GetFirstMultipartFileHeader(c, fieldName)
	if header == nil {
		return nil, nil, nil
	}

	file, err := header.Open()
	if err != nil {
		return nil, nil, err
	}

	return &file, header, nil
}
