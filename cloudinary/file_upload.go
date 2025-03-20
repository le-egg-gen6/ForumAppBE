package cloudinary

import (
	"context"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/go-playground/validator/v10"
	"myproject/forum/dtos"
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

func NewFileUploader() *FileUploader {
	cfg, err := LoadCloudinaryConfig()
	if err != nil {
		panic("Cloudinary configuration file not found")
	}
	Instance = &FileUploader{
		Validator: validator.New(),
		Config:    cfg,
	}
	return Instance
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
