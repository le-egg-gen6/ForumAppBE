package cloudinary_service

import (
	"context"
	"fmt"
	"forum/dtos"
	"forum/logger"
	"forum/models"
	"forum/repository"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"mime/multipart"
	"time"
)

type IFileUploader interface {
	UploadFile(file dtos.File) (string, error)
}

type FileUploader struct {
	Validator *validator.Validate
	Config    *CloudinaryConfig
}

func InitializeCloudinaryInstance(cfg *CloudinaryConfig) (*cloudinary.Cloudinary, error) {
	return cloudinary.NewFromParams(cfg.CloudName, cfg.APISecret, cfg.APISecret)
}

var Instance *FileUploader

func GetFileUploaderInstance() *FileUploader {
	return Instance
}

func InitializeFileUploader() {
	cfg, err := LoadCloudinaryConfig()
	if err != nil {
		panic("Cloudinary configuration not found")
	}
	Instance = &FileUploader{
		Validator: validator.New(),
		Config:    cfg,
	}
}

func (u *FileUploader) UploadFileToCloudinary(file *multipart.File) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := u.Validator.Struct(file)
	if err != nil {
		return "", err
	}

	cld, err := InitializeCloudinaryInstance(u.Config)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{Folder: u.Config.UploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}

const (
	maxRetry = 5
)

func (u *FileUploader) UploadFile(file *dtos.File) (*models.Image, error) {
	var fileUrl string
	for attempt := 0; attempt <= maxRetry; attempt++ {
		if attempt > 0 {
			logger.GetLogInstance().Warn(fmt.Sprintf("Retrying upload file: %s, attempt %d of %d", file.FileHeader.Filename, attempt, maxRetry))
		}

		fileUrlTemp, err := u.UploadFileToCloudinary(file.File)
		if err == nil {
			fileUrl = fileUrlTemp
			break
		}
		if attempt == maxRetry {
			logger.GetLogInstance().Error(fmt.Sprintf("Failed to upload file: %s", file.FileHeader.Filename))
			break
		}
	}
	if fileUrl != "" {
		image := models.Image{
			URL: fileUrl,
		}
		return repository.GetImageRepositoryInstance().Create(&image)
	}
	return nil, fmt.Errorf("failed to upload file: %s", file.FileHeader.Filename)
}
