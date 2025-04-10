package handler

import (
	"forum/3rd_party_service/cloudinary"
	"forum/dtos"
	"forum/shared"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func InitializeFileHandler(router *gin.RouterGroup) {
	fileGroup := router.Group("/file")
	{
		fileGroup.POST("/save", SaveFile)
	}
}

const (
	maxFileSize = 10 << 20 //10MB
)

var allowedFileTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func SaveFile(c *gin.Context) {
	if err := c.Request.ParseMultipartForm(maxFileSize); err != nil {
		shared.SendError(c, http.StatusBadRequest, "Failed to parse form")
		return
	}

	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		shared.SendError(c, http.StatusBadRequest, "Invalid file")
		return
	}

	if fileHeader.Size > maxFileSize {
		shared.SendError(c, http.StatusBadRequest, "File size is too large")
		return
	}

	fileExtension := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):]
	if _, ok := allowedFileTypes[fileExtension]; !ok {
		shared.SendError(c, http.StatusBadRequest, "Invalid file type")
		return
	}
	fileDtos := dtos.File{
		FileHeader: fileHeader,
		File:       &file,
	}

	fileUrl, err := cloudinary.GetFileUploaderInstance().UploadFile(&fileDtos)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	shared.SendSuccess(c, dtos.Url{Url: fileUrl})
}
