package dtos

import "mime/multipart"

type File struct {
	FileHeader *multipart.FileHeader
	File       *multipart.File
}
