package cloudinary

import (
	"context"
	"forum/dtos"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
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

func (u *FileUploader) UploadFile(file dtos.File) (string, error) {
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

	uploadParam, err := cld.Upload.Upload(ctx, file.File, uploader.UploadParams{Folder: u.Config.UploadFolder})
	if err != nil {
		return "", err
	}
	return uploadParam.SecureURL, nil
}
