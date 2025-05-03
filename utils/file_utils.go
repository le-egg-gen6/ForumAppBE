package utils

import (
	"forum/constant"
	"mime/multipart"
	"strings"
)

func IsFileValid(fileHeader *multipart.FileHeader) bool {
	fileExtension := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, ".")+1:]
	_, ok := constant.AllowedFileTypes[fileExtension]
	return fileHeader.Size < constant.MaxFileSize && ok
}
