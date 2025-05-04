package handler

import (
	"forum/3rd_party_service/cloudinary_service"
	"forum/constant"
	"forum/dtos"
	"forum/models"
	"forum/repository"
	"forum/server/http_server/middlewares"
	"forum/shared"
	"forum/utils"
	"github.com/gin-gonic/gin"
)

func InitializePostHandler(router *gin.RouterGroup) {
	postGroup := router.Group("/post")
	{
		postGroup.POST("/create", middlewares.AuthenticationMiddlewares(), CreateNewPost)
	}
}

func CreateNewPost(c *gin.Context) {
	userID := utils.GetCurrentContextUserID(c)
	if userID < 0 {
		shared.SendUnauthorized(c)
		return
	}

	user, err := repository.GetUserRepositoryInstance().FindByID(uint64(userID))
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	if user == nil {
		shared.SendUnauthorized(c)
		return
	}

	var createPostDTO dtos.SimplePostDTO
	if err := c.ShouldBind(&createPostDTO); err != nil {
		shared.SendBadRequest(c, "Bad Request")
		return
	}

	files, fileHeaders, err := utils.GetRequestMultipartFiles(c, constant.FileFormKey)
	if err != nil {
		shared.SendBadRequest(c, "Bad Request")
		return
	}

	checkFile := true
	for _, fileHeader := range fileHeaders {
		checkFile = checkFile && utils.IsFileValid(fileHeader)
	}
	if !checkFile {
		shared.SendBadRequest(c, "Invalid file")
		return
	}

	images := make([]*models.Image, 0, len(files))
	for i := 0; i < len(files); i++ {
		fileDto := dtos.File{
			FileHeader: fileHeaders[i],
			File:       files[i],
		}
		image, err := cloudinary_service.GetFileUploaderInstance().UploadFile(&fileDto)
		if err == nil {
			images = append(images, image)
		}
	}

	post := &models.Post{
		AuthorID: &user.ID,
		Content:  createPostDTO.Content,
		Images:   images,
	}
	post, err = repository.GetPostRepositoryInstance().Create(post)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	user.Posts = append(user.Posts, post)
	err = repository.GetUserRepositoryInstance().Update(user)
	if err != nil {
		shared.SendInternalServerError(c)
		return
	}
	shared.SendSuccess(c, post)
}
